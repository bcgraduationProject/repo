#!/bin/bash

#run without sudo
#run this command
chmod +x steps2.sh

mkdir -p $HOME/go/src/github.com/1234
cd $HOME/go/src/github.com/1234

curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh

./install-fabric.sh docker samples binary

echo "DONE!-------------"

