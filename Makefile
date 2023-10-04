build_game:
	GOOS=js GOARCH=wasm go build -o static-server/assets/game.wasm pinball/*.go

run:
	cd static-server && go run main.go
