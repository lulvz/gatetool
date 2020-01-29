#!/bin/sh

# this just automates uninstall process

# it doesnt uninstall rpitx or golang tho just for the sake of time
# if you want to uninstall them it must be done manually

sudo rm -rf /usr/local/gateTool
sudo rm -rf $HOME/.gateToolFiles
sudo systemctl stop gateToolWeb.service
sudo systemctl disable gateToolWeb.service
sudo rm /etc/systemd/system/gateToolWeb.service
