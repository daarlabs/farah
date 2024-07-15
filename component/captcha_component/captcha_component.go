package captcha_component

import (
	"embed"
	"fmt"
	"math/rand"
	"time"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/farah/palette"
	"github.com/daarlabs/farah/ui/form_ui/checkbox_ui"
	"github.com/daarlabs/farah/ui/form_ui/error_message_ui"
	"github.com/daarlabs/farah/ui/form_ui/hidden_field_ui"
	"github.com/daarlabs/hirokit/alpine"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/hiro"
	"github.com/daarlabs/hirokit/hx"
	"github.com/daarlabs/hirokit/tempest"
)

type Captcha struct {
	hiro.Component
	Answer         int       `json:"answer"`
	Attempt        int       `json:"attempt"`
	LastAttempt    time.Time `json:"lastAttempt"`
	ChallengeToken string    `json:"-"`
	SuccessToken   string    `json:"-"`
	Valid          bool      `json:"-"`
}

type captcha struct {
	ChallengeToken string `json:"challengeToken"`
	SuccessToken   string `json:"successToken"`
}

const (
	stateUnchecked = iota
	stateChecked
	stateSuccess
	stateFail
	stateLimitExceeded
)

const (
	maxAttempts = 5
)

const (
	captchaChallenge = "captcha-challenge"
	captchaSuccess   = "captcha-success"
)

var (
	limitTimeWindow = time.Hour
)

//go:embed img/*.webp
var images embed.FS

func (c *Captcha) Name() string {
	return "captcha"
}

func (c *Captcha) Mount() {
	if c.Request().Is().Action() {
		c.Parse().MustQuery(captchaChallenge, &c.ChallengeToken)
	}
	if !c.Request().Is().Action() {
		c.Answer = 0
		c.ChallengeToken = uniuri.New()
		c.SuccessToken = ""
	}
	if !c.Request().Is().Action() || c.Request().Is().Action("HandleChoose") {
		if time.Now().After(c.LastAttempt.Add(limitTimeWindow)) {
			c.Attempt = 0
		}
	}
}

func (c *Captcha) Node() Node {
	return c.createCaptcha(stateUnchecked, []int{})
}

func (c *Captcha) HandleCheck() error {
	randomRange := c.generateRandomRange()
	for i, number := range randomRange {
		if number == 1 {
			c.Answer = i + 1
			break
		}
	}
	return c.Response().Render(c.createCaptcha(stateChecked, randomRange))
}

func (c *Captcha) HandleChoose() error {
	if c.Attempt >= maxAttempts {
		c.Answer = 0
		return c.Response().Render(c.createCaptcha(stateLimitExceeded, []int{}))
	}
	var answer int
	c.Parse().MustQuery("answer", &answer)
	c.LastAttempt = time.Now()
	c.Attempt += 1
	if answer != c.Answer {
		randomRange := c.generateRandomRange()
		for i, number := range randomRange {
			if number == 1 {
				c.Answer = i + 1
				break
			}
		}
		return c.Response().Render(c.createCaptcha(stateFail, randomRange))
	}
	c.Answer = 0
	c.Attempt = 0
	c.LastAttempt = time.Time{}
	c.SuccessToken = uniuri.New()
	if err := save(c, c.ChallengeToken, c.SuccessToken); err != nil {
		return c.Response().Error(err)
	}
	return c.Response().Render(c.createCaptcha(stateSuccess, []int{}))
}

func (c *Captcha) GetImg() error {
	var img string
	c.Parse().MustQuery("img", &img)
	imgBytes, err := images.ReadFile("img/" + img + ".webp")
	if err != nil {
		return c.Response().Error(err)
	}
	return c.Response().File(img+".webp", imgBytes)
}

func (c *Captcha) createCaptcha(state int, randomRange []int) Node {
	return Div(
		Div(
			Id(hx.Id("captcha")),
			tempest.Class().Grid().Gap(1),
			Div(
				tempest.Class().Transition().Grid().PlaceItems("center").
					W("full").MinH(20).P(4).Rounded().
					BgWhite().Bg(tempest.Slate, 800, tempest.Dark()).
					Border(1).BorderColor(tempest.Slate, 300).BorderColor(tempest.Slate, 600, tempest.Dark()),
				hidden_field_ui.HiddenField(captchaChallenge, c.ChallengeToken),
				If(
					state == stateUnchecked,
					checkbox_ui.Checkbox(
						checkbox_ui.Props{Id: "captcha", Label: c.getCheckboxLabel()},
						hx.Get(c.Generate().Action("HandleCheck", hiro.Map{captchaChallenge: c.ChallengeToken})),
						hx.Target(hx.HashId("captcha")),
						hx.Trigger("change delay:400ms"),
						hx.Swap(hx.SwapOuterHtml),
					),
				),
				If(
					state == stateChecked || state == stateFail,
					If(
						state == stateChecked,
						Div(
							tempest.Class().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()).Pb(4).TextCenter().TextXs(),
							Text(c.Translate("component.captcha.images.title")),
						),
					),
					If(
						state == stateFail, Div(
							Div(
								tempest.Class().FontSemibold().TextSlate(900).TextWhite(tempest.Dark()).Pb(4).TextCenter().TextXs(),
								Text(c.Translate("component.captcha.images.title")),
							),
							Div(
								tempest.Class().FontSemibold().TextRed(600).TextRed(500, tempest.Dark()).
									Pb(4).TextCenter().TextXs().LhRelax(),
								Text(c.Translate("component.captcha.fail")),
							),
						),
					),
					Div(
						alpine.Data(hiro.Map{"active": 0}),
						tempest.Class().Grid().GridCols(4).Gap(2).PlaceItemsCenter(),
						Range(
							randomRange, func(number int, i int) Node {
								name := fmt.Sprintf("galaxy_%d", number)
								return Button(
									tempest.Class().Name(c.Request().Action()).Peer().Transition().Size(12).Overflow("hidden").Rounded().
										Border(1).BorderSlate(900).BorderWhite(tempest.Dark()).
										BorderColor(palette.Primary, 400, tempest.Hover()).
										BorderColor(palette.Primary, 200, tempest.Hover(), tempest.Dark()),
									Type("button"),
									alpine.MouseEnter(fmt.Sprintf("active = %d", number)),
									alpine.MouseLeave("active = 0"),
									alpine.Bind(
										"class",
										fmt.Sprintf(
											"active > 0 && active != %d && '%s'", number, tempest.Class().Opacity(0.3).String(),
										),
									),
									hx.Get(
										c.Generate().Action(
											"HandleChoose", hiro.Map{"answer": i + 1, captchaChallenge: c.ChallengeToken},
										),
									),
									hx.Target(hx.HashId("captcha")),
									hx.Trigger("click"),
									hx.Swap(hx.SwapOuterHtml),
									Img(
										tempest.Class().Name(c.Request().Name()).Size(12).BgSlate(300).BgSlate(600, tempest.Dark()),
										Loading("lazy"),
										Src(c.Generate().Action("GetImg", hiro.Map{"img": name})),
										Alt(name),
									),
								)
							},
						),
					),
				),
				If(
					state == stateSuccess,
					hidden_field_ui.HiddenField(captchaSuccess, c.SuccessToken),
					Div(
						tempest.Class().FontSemibold().TextEmerald(500).TextEmerald(400, tempest.Dark()).
							P(4).TextCenter().TextXs().LhRelax(),
						Text(c.Translate("component.captcha.success")),
					),
				),
				If(
					state == stateLimitExceeded,
					Div(
						tempest.Class().FontSemibold().TextRed(600).TextRed(500, tempest.Dark()).
							P(4).TextCenter().TextSize("10px").LhRelax(),
						Text(c.Translate("component.captcha.limit")),
					),
				),
			),
			If(!c.Valid, error_message_ui.ErrorMessage(c.Translate("error.invalid.captcha"))),
		),
	)
}

func (c *Captcha) getCheckboxLabel() string {
	if c.Config().Localization.Enabled {
		return c.Translate("component.captcha.checkbox.label")
	}
	return "Confirm, you are a human being"
}

func (c *Captcha) generateRandomRange() []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	numbers := []int{1, 2, 3, 4}
	r.Shuffle(
		len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		},
	)
	return numbers
}
