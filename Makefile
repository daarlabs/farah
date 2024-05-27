dev: dev-go dev-assets

dev-go:
	reflex --all=false -r '(\.go$$|go\.mod|\.css$$|\.ts$$|\.js$$)' -s go run .

dev-assets:
	cd assets && pnpm watch

deps:
	GOPRIVATE=github.com/software-trading-platform/repository go get github.com/software-trading-platform/repository