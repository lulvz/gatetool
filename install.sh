#!/bin/sh

# run this as normal user and input password when prompted to

#check if golang is installed
if ! hash go 2>/dev/null; then
    echo "\ngolang isn't installed yet, instaling...\n"
    wget https://dl.google.com/go/go1.13.7.linux-armv6l.tar.gz
    echo "\nextracting tar.gz file into /usr/local\n"
	sudo tar -C /usr/local -xzf go1.13.7.linux-armv6l.tar.gz

    export PATH="$PATH:/usr/local/go/bin"
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile

fi
echo "\ngolang is installed, continuing...\n"

# check if rpitx is installed
if ! hash rpitx 2>/dev/null; then
    echo "\nrpitx isn't installed yet, installing...\n"
    sudo apt update

    if ! hash git 2>/dev/null; then
        echo "\ninstalling git...\n"
        sudo apt install git
    fi
    echo "\ngit is already installed, continuing\n"
    git clone https://github.com/F5OEO/rpitx /tmp/rpitx/
    echo "\ngiving execute permissions to the install.sh file on /tmp/rpitx/\n"
    sudo chmod +x /tmp/rpitx/install.sh

    # stores current directory in variable
    startingDir="$PWD"
    cd /tmp/rpitx/
    
    echo "\nstarting install of rpitx...\n"
    sudo ./install.sh
    # goes back to the directory in which the install.sh was executed
    cd "$startingDir"
fi
echo "\nrpitx is installed, continuing...\n"

echo "\ncreating gateToolScripts directory for the scripts in $HOME/.gateToolFiles/\n"
mkdir $HOME/.gateToolFiles
cp -r ./gateToolScripts $HOME/.gateToolFiles/

# this puts the webserver and main config files in the .gateToolFiles directory in the home folder
echo "\ncreating config files gateToolConfig.json and gateToolWebConfig.json in $Home/.gateToolFiles/\n"
cp dist/gateToolConfig.json $HOME/.gateToolFiles
cp dist/gateToolWebConfig.json $HOME/.gateToolFiles/

#### webserver installation ####

echo "\ncompiling webserver...\n"
startingDir="$PWD"
cd controlserver
go build -o build/controlserver
cd "$startingDir"

# config file and web files being copied here
echo "\ninstalling webserver in /usr/local/gateToolWeb/\n"
sudo mkdir /usr/local/gateTool
sudo mkdir /usr/local/gateTool/gateToolWeb
sudo cp -r ./dist/web/* /usr/local/gateToolWeb
sudo cp ./controlserver/build/controlserver /usr/local/gateTool/gateToolWeb

# enabling the services here
echo "\ncreating and enabling gateToolWeb.service\n"
sudo cp dist/gateToolWeb.service /etc/systemd/system/gateToolWeb.service

sudo systemctl enable gateToolWeb.service
sudo systemctl start gateToolWeb.service
