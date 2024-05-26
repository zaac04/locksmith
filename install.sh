#!/bin/bash
url="https://github.com/zaac04/locksmith/releases/download/v1.0/locksmith"
wget -O locksmith $url
binary_destination_folder="/usr/local/bin"
sudo cp locksmith $binary_destination_folder
sudo chmod +x "$binary_destination_folder/locksmith"
sudo locksmith completion bash >> /etc/bash_completion.d/locksmith.sh
sudo echo "source /etc/bash_completion.d/locksmith.sh" >> ~/.bashrc
. ~/.bashrc

