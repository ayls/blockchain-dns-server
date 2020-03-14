#!/bin/bash
geth --nousb --datadir=$HOME/.rinkeby init rinkeby.json
NEW_UUID=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | head -c 32)
geth --nousb --networkid=4 --rpc --datadir=$HOME/.rinkeby --syncmode=light --ethstats='$NEW_UUID:Respect my authoritah!@stats.rinkeby.io' --bootnodes=enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303