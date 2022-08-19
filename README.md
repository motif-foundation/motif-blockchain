#Motif Blockchain

##Preqs
``
	sudo apt update
	sudo apt install curl nano screen git build-essential
	curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
	sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
	sudo nano ~/.profile
	add this to the end!!!!!!
	export PATH=$PATH:/usr/local/go/bin
``
## Create Genesis  
``
	make
	cd /root/motif/motif-blockchain/build
	./motif --datadir /root/motif/motif-blockchain/build/validatordata validator genesis 
	give password 
	note the keys and wallet addresses
``