this folder contains a shellscript that can be used to compile .sol files to a .go ABI file
dependencies that are needed are:
- solc
- abigen

both of these dependencies can be build using the go-ethereum library from github.com
they also to be globaly available by including them in $PATH 

this can be done by running the provision.sh script that is used to provision new vagrant virtual machine's

these are the commands

# go get -u github.com/ethereum/go-ethereum/...
# cd $GOPATH/src/github.com/ethereum/go-ethereum
# make && make devtools

more information can be found in the DEV_VM folder

