__kh-sel() {
  setopt localoptions noglobsubst noposixbuiltins pipefail no_aliases 2> /dev/null
  selected=( $(__kube-helper-list-record-commands |
    FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-40%} $FZF_DEFAULT_OPTS -n2..,.. --tiebreak=index --bind=ctrl-z:ignore $FZF_DEFAULT_OPTS $FZF_CTRL_T_OPTS" $(__fzfcmd) -m "$@") )

  if [[ "$selected" == "" ]]; then
    return
  fi

  setopt localoptions pipefail no_aliases 2> /dev/null
  item=( $(__kube-helper-load-record "$$-$selected[1]" | 
    FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-40%} --reverse --bind=ctrl-z:ignore $FZF_DEFAULT_OPTS $FZF_CTRL_T_OPTS" $(__fzfcmd) -m "$@") )

  if [[ "$item" == "" ]]; then
    return
  fi

  setopt localoptions pipefail no_aliases 2> /dev/null
  ret=( $(echo $item | tr ' ' '\n' | 
    FZF_DEFAULT_OPTS="--height ${FZF_TMUX_HEIGHT:-40%} --reverse --bind=ctrl-z:ignore $FZF_DEFAULT_OPTS $FZF_CTRL_T_OPTS" $(__fzfcmd) -m "$@") )
  echo $ret
}

kube-helper-analyse-get-result() {
  LBUFFER="${LBUFFER}$(__kh-sel)"
  local ret=$?
  zle reset-prompt
  return $ret
}

kube-helper-add-recorder() {
  [[ "$BUFFER" = *"kube-record" ]] || BUFFER="$BUFFER | kube-record"
  local ret=$?
  zle reset-prompt
  return $ret
}

zle     -N             kube-helper-analyse-get-result
bindkey -M emacs '\ek' kube-helper-analyse-get-result
bindkey -M vicmd '\ek' kube-helper-analyse-get-result
bindkey -M viins '\ek' kube-helper-analyse-get-result

zle     -N             kube-helper-add-recorder
bindkey -M emacs '\er' kube-helper-add-recorder
bindkey -M vicmd '\er' kube-helper-add-recorder
bindkey -M viins '\er' kube-helper-add-recorder
