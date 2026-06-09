
# Added by ghayr
nvim() {
  local current
  current="$(cat "$HOME/.config/ghayr/.current" 2>/dev/null)"
  if [ -n "$current" ]; then
    NVIM_APPNAME=ghayr/configs/$current command nvim "$@"
  else
    command nvim "$@"
  fi
}
