#!/bin/bash

set +x

IPArr=( '192.168.100.61' '192.168.100.62' '192.168.100.63' )

function postBuild {
    AppName=${PWD##*/}
    BuildTime=$(date +"%F_%T")
    Tag=$(git describe --tags 2> /dev/null)
    Hash=$(git log --pretty=format:'%h' -n 1)
    if [[ $Tag == *"-g"* ]] || [[ $Tag == "" ]]; then
        Gitver="dev"
    else
        Gitver=$Tag
    fi
    go $1 -ldflags "-X main.name=$AppName -X main.version=$Gitver -X main.builded=$BuildTime -X main.hash=$Hash" .
}

function buildNode {
    echo "Build "${PWD##*/}
    postBuild build
}

function runNode {
    echo "Run "${PWD##*/}
    postBuild run
}

function deployNode {
    echo "Deploy code to nodes"
    for i in "${IPArr[@]}"
    do
        echo "Deploy to $i"
        ssh sysop@$i 'cd ~/go/src/elevator && git pull && docker-compose build && docker-compose up' &
    done
}

function killNode {
    echo -n "Enter IP of node: 192.168.100."
    read param
    IP="192.168.100.$param"
    echo "Kill node $IP"
    ssh sysop@$"$IP" 'docker stop elevator_elevator_1'
}

function help {
    echo -e "Managment script for control node\n"
    echo "Help:"
    echo -e "\t-b - build node"
    echo -e "\t-d - deploy and start node"
    echo -e "\t-k - kill node"
    echo -e "\t-r - run node"
    echo -e "\t-h - help"
}

if [ $# -eq 0 ]; then
    help
    exit 0
fi

while [ -n "$1" ]
do
    case "$1" in
        -d) deployNode ;;
        -k) killNode ;;
        -r) runNode ;;
        -b) buildNode ;;
        -h) help ;;
        *) echo "$1 is not an option" ;;
    esac
shift
done
