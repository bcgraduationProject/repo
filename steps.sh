#!/bin/bash

###hyperledger documentation v2.5###

#run this command
#run with sudo
#chmod +x steps.sh

#to run the other two files correctly (steps2.sh, steps3.sh)
sudo chmod +x steps2.sh
sudo chmod +x steps3.sh



sudo apt-get install git
sudo apt-get install curl
sudo apt-get -y install docker-compose
sudo systemctl start docker
#Optional: If you want the Docker daemon to start when the system starts, use the #following:
sudo systemctl enable docker
#change the usernaame of the below file
sudo usermod -a -G docker user
#go files
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin


echo "done ya raye2"
