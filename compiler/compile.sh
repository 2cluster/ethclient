#!/bin/bash

helpFunction()
{
   echo ""
   echo "Usage: Supply at least 1 arg so i know what contract to compile"
   echo -e "       Second argument can be used for a custom output location\n"
   exit 1 # Exit script after printing help
}

timerFunction()
{
   echo "";echo -n $1
   for i in `seq 1 40`;
   do
    echo -n "."
    sleep 0.02
   done
   echo -e "\n";sleep 0.2
}

BIN='./contracts/artifacts/bin'
ABI='./contracts/artifacts/abi'



if [ ! -z $1 ] 
then 
   ARG1=${1}
else
   helpFunction
fi


if [ ! -z $2 ] 
then 
    echo -e "\nOutput location '$2' will be used\n"
    ARG2=${2}
else
    echo -e  "\nOutput location is not provided"
    echo -e "The default location will be used './app/contract'\n"
    ARG2='./contracts/abi'
fi


timerFunction "Generating BIN"
solc --bin -o ${BIN} github.com/OpenZeppelin/=$GOPATH/src/github.com/ ./contracts/${ARG1} --allow-paths $GOPATH/src/github.com --overwrite &> ./error1
if [ $? -eq 0 ]; then
    echo OK
else
    echo FAIL
    cat ./error1
    exit 1
fi

timerFunction "Generating ABI"
solc --abi -o ${ABI} github.com/OpenZeppelin/=$GOPATH/src/github.com/ ./contracts/${ARG1} --allow-paths $GOPATH/src/github.com --overwrite &> ./error2
if [ $? -eq 0 ]; then
    echo OK
else
    echo FAIL
    cat ./error2
    exit 1
fi

NAME="$(cut -d'.' -f1 <<<"$ARG1")"

timerFunction "Generating GOABI"
abigen --bin ${BIN}/${NAME}.bin --abi ${ABI}/${NAME}.abi --pkg=${NAME} --out=${ARG2}/${NAME}.go &> ./error3
if [ $? -eq 0 ]; then
    echo OK
else
    echo FAIL
    cat ./error3
    exit 1
fi
