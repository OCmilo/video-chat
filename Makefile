server:
	go build -o go-soundboard *.go

ui:
	yarn --cwd ./front/ run build

dev:
	go run *.go || yarn --cwd ./front/ run dev