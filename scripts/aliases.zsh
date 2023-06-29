__kube-helper-list-record-commands() {
    fc -rl 1 | perl -ne 'print if !$seen{(/^\s*[0-9]+\**\s+(.*)/, $1)}++' | awk '{raw=$0;file=sprintf("%s/.kube-helper/records/%s-%s","'$HOME'","'$$'",$1);if((getline<file)>=0){print raw}}'
}

__kube-helper-record() {
    id=`echo $$`-`print -P %h`
    mkdir -p ~/.kube-helper/records
    recordfile=~/.kube-helper/records/$id
    cat | tee -a $recordfile
}

__kube-helper-load-record() {
    recordfile=~/.kube-helper/records/$1
    cat $recordfile 2> /dev/null
}

__kube-helper-clear-records() {
    rm -rf `find ~/.kube-helper/records | grep $$-`
}

trap __kube-helper-clear-records EXIT

alias -g kube-record="__kube-helper-record"

__kube-helper-add-recorder_for_get() {
    NEW="$BUFFER"
    if [[ "$NEW" = "kubectl "*"get "* ]] || [[ "$NEW" = "k "*"get "* ]]; then
        [[ "$NEW" = *"kube-record" ]] || NEW="$NEW | kube-record"
    fi
    BUFFER="$NEW"
    zle .$WIDGET "$@"
}

zle -N accept-line __kube-helper-add-recorder_for_get
