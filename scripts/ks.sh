#!/bin/sh

### ks

function __ks {
    kubectl-switch list --suggest
}

function ks {
    if [ "$1" = "" ]; then 
        kubectl-switch list
    else
        \. <(kubectl-switch use $1)
    fi
}

complete -o default -F __ks ks

### ksns

function __ksns {
    kubectl get namespace --no-headers -o custom-columns=:.metadata.name 2> /dev/null
}

function ksns {
    kubectl config set-context --current --namespace $1
}

complete -o default -F __ksns ksns
