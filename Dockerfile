FROM golang:1.16.3-alpine3.13

WORKDIR /neovim-integration-twitch
COPY . /neovim-integration-twitch

CMD go run main.go
