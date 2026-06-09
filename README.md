# ghayr

> غير — Arabic for "change"

Switch between Neovim configurations using `NVIM_APPNAME`.

## How it works

Configs live in `~/.config/ghayr/configs/`, each as a subdirectory. The active config name is stored in `~/.config/ghayr/.current`. A shell function wraps `nvim` to read it on every launch.

```
~/.config/ghayr/
├── configs/
│   ├── lazyvim/
│   ├── kickstart/
│   └── astronvim/
└── .current
```

## Requirements

- [fzf](https://github.com/junegunn/fzf)
- Neovim 0.9+ (`NVIM_APPNAME` support)

## Installation

### 1. Install the binary

```sh
go install github.com/isaacwassouf/ghayr@latest
```

### 2. Initialize ghayr

```sh
ghayr init
```

This creates `~/.config/ghayr/configs/` and `~/.config/ghayr/.current`, and appends the `nvim` shell function to your `.zshrc` or `.bashrc`.

Reload your shell:

```sh
source ~/.zshrc  # or ~/.bashrc
```

## Usage

```sh
ghayr                        # open fzf picker to select a config
ghayr use <name>             # switch directly without fzf
ghayr list                   # print all available configs
ghayr add -r <repo> -n <name> # clone a config from a git repo
```

### Add a config from a git repo

```sh
ghayr add --repo https://github.com/LazyVim/starter --name lazyvim
```

This clones the repo into `~/.config/ghayr/configs/lazyvim`.

### Switch config

```sh
ghayr use lazyvim
```

Or just run `ghayr` to open an interactive fzf picker.

After switching, run `nvim` as normal — it picks up the active config automatically via `NVIM_APPNAME=ghayr/configs/<name>`.
