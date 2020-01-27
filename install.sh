#!/bin/sh

# run this as normal user and input password when prompted to

if ! hash rpitx 2>/dev/null; then
    echo "rpitx isn't installed yet, installing..."
    sudo apt update

    if ! hash git 2>/dev/null; then
        echo "installing git..."
        sudo apt install git
    fi
    echo "git is already installed, continuing"
    git clone https://github.com/F5OEO/rpitx /tmp/rpitx/
    echo "giving execute permissions to the install.sh file on /tmp/rpitx/"
    sudo chmod +x /tmp/rpitx/install.sh

    # stores current directory in variable
    startingDir="$PWD"
    cd /tmp/rpitx/
    
    echo "starting install of rpitx..."
    sudo ./install.sh
    # goes back to the directory in which the install.sh was executed
    cd "$startingDir"
fi

echo "rpitx is already installed, continuing..."

echo "creating .gateToolScripts directory for the scripts in $HOME"
mkdir $HOME/.gateToolScripts
cp ./scripts_waves/* $HOME/.gateToolScripts/
