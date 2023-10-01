build_game:
	GOOS=js GOARCH=wasm go build -o static-server/assets/game.wasm pinball/main.go

run:
	cd static-server && go run main.go
