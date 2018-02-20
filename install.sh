#!/bin/bash
set -e
echo Asking for your user password because we need root privileges.
echo "The compiled executable will be put in root's path."
echo so that when you type sudo pianokeylogger, it works
echo ...
echo First creating a directory /opt/util/sound/ to put the wav files in

sudo mkdir -p /opt/util/sound
sudo cp ./pianosoundwavs/*.wav /opt/util/sound
go build
sudo cp pianokeylogger /usr/bin/
rm pianokeylogger

