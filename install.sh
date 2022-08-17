#!/bin/sh

set -e

install() {
    go install ./cmd/kubectl-switch

    mkdir -p $HOME/.kube-helper
    cp scripts/ks.sh $HOME/.kube-helper
    cp scripts/aliases.zsh $HOME/.kube-helper
    cp scripts/key-bindings.zsh $HOME/.kube-helper
    cp scripts/source $HOME/.kube-helper

    echo "write the following command into your .zshrc and re-source:"
    echo ""
    echo "    source $HOME/.kube-helper/source"
    echo ""
    echo "and copy hack/ys.zsh-theme to ~/.oh-my-zsh/themes/ys.zsh-theme"
}

install