# HLF-practicing
Investigating HLF


# Step by step

#### Installation of deps

- install docker - [https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04)

- install docker compose - [https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04)

- innstall jq, git, golang, java

- hlf docks - [https://hyperledger-fabric.readthedocs.io/en/release-2.5](https://hyperledger-fabric.readthedocs.io/en/release-2.5)

- cid package docs - [https://github.com/hyperledger/fabric-chaincode-go/blob/main/pkg/cid/README.md](https://github.com/hyperledger/fabric-chaincode-go/blob/main/pkg/cid/README.md)

#### Start with HLFs

- clone hlf project - `git clone https://github.com/hyperledger/fabric-samples`

- `cd fabric-samples`

- `cd test-network && ./network.sh prereq` - install fabric binaries and docker images

- `add fabric-samples/bin to PATH`

- `./network.sh up` - set up the network

- `./network.sh up -s couchdb` - set up the network with couchdb

- `./network.sh down` - remove all running containers, shut down the network, delete all chain-codes

- `./network.sh createChannel -s couchdb -c channelName` - set up network with channel

- `./network.sh up createChannel -ca -c mychannel -s couchdb`

- `./network.sh deployCC -ccn sacc -ccp ../sacc -ccl go -ccv 1` - deploy a chain-code

- `peer lifecycle chaincode queryinstalled` - get deployted on peer smart contracts(chain-codes)

- `peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -n sacc -c '{"Args":["set","asset113", "yellow", "5", "Tom", "1300"]}'`

- `docker logs -f` - see logs of certain container

- `docker exec -it ${containerID} env` - see env variables

- `peer chaincode query -C mychannel -n basic -c '{"function": "CreateAsset", "Args": ["assert113", "yellow", "5", "Tom", "1300"]}'`

- `docker rm, docker ps, docker stop`


```
export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/go/bin

export PATH=$PATH:/home/hakob/Desktop/HLF2024/fabric-samples/bin
export FABRIC_CFG_PATH=/home/hakob/Desktop/HLF2024/fabric-samples/config

# org 1
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

- `curl -v -d "{}" localhost:3333/securedValue -H "user: Admin@org1.example.com" -H "key: aboba" -H "value: zksnarks"` - test hometask
