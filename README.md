# NeoVim Twitch Integration

NTI (Neovim Twitch Integration) is a Golang Twitch Bot what your viewers can change your Neovim theme, and move your cursor inside Neovim.

# Setting Up

Unzip `colors.zip`

Install [Neovim Remote](https://github.com/mhinz/neovim-remote)

Create a `.env` file with your Twitch's token, like:

```
TOKEN=[token]
USER=[user]
CHANNEL=[channel]
```

# Commands

- `!themes` list all themes in `/colors`
- `!color` set a theme from all installed themes
- `!move` Move your cursor inside Vim
