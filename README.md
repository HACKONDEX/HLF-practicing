# HLF-practicing
Investigating HLF


# Step by step

#### Installation of deps

- install docker - [https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04)

- install docker compose - [https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04)

- innstall jq, git, golang, java

- hlf docks - [https://hyperledger-fabric.readthedocs.io/en/release-2.5](https://hyperledger-fabric.readthedocs.io/en/release-2.5)

#### Start with HLFs

- clone hlf project - `git clone https://github.com/hyperledger/fabric-samples`

- `cd fabric-samples`

- `cd test-network && ./network.sh prereq` - install fabric binaries and docker images

- `add fabric-samples/bin to PATH`

- `./network.sh up` - set up the network

- `./network.sh up -s couchdb` - set up the network with couchdb

- `./network.sh down` - remove all running containers, shut down the network, delete all chain-codes

