#!/bin/sh

install_solc(){

    sudo apt install python3-pip

    pip3 install solc-select

    export PATH=$PATH:/home/vagrant/.local/bin

    sudo touch "$HOME/.zshrc"
    {
        echo "export PATH=\$PATH:/home/vagrant/.local/bin"
    } >> "$HOME/.zshrc"
}

# install protoc
install_protoc(){
    sudo apt install protobuf-compiler -y
}

install_solc
install_protoc