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

__kubectl() {
  kubectl "$@" | __kube-helper-record
}

alias k=__kubectl
