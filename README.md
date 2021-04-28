# NeoVim Twitch Integration

NTI (Neovim Twitch Integration) is a Golang Twitch Bot what your viewers can change your Neovim theme, and move your cursor inside Neovim.

# Setting Up

Unzip `colors.zip`

Create a `.env` file with your Twitch's token, like:

```
TOKEN=[token]
USER=[user]
CHANNEL=[channel]
PORT=[number]
```

> Remember: PORT is a UNIX port, like `6126` or `7352`. Put a port what is not being use to any other program.

# Running

## Golang

Run `go run main.go` to start the server.

## Docker

Run `docker run -it --net=host neovim-integration-twitch` to start the server. If you prefer docker composer, use `docker-compose up` in the repository directory.

---

And Use `nvim --listen 127.0.0.1:[PORT]` and you Vim is running with the integration.

# Commands

- `!themes` list all themes in `/colors`
- `!color` set a theme from all installed themes
- `!move` Move your cursor inside Vim
