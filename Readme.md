# Asset Management Smart Contract

This project implements a Hyperledger Fabric smart contract for managing assets. The smart contract allows for the creation, querying, and updating of asset information associated with dealers in a blockchain environment. 


# Hyperledger Samples Reference
```
https://github.com/hyperledger/fabric-samples.git
```

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Smart Contract Functions](#smart-contract-functions)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

This smart contract is designed for a financial institution to manage and track assets efficiently. It allows dealers to create and manage their asset information on a secure blockchain platform, ensuring transparency and immutability.


# Docker Commands for Hyperledger Fabric Project

#1 Start the Hyperledger Fabric Network: If you have a docker-compose.yaml file set up for your Hyperledger Fabric network, you can start the network with:

```
docker-compose up -d
```
This command starts the containers in detached mode.


#2 Stop the Hyperledger Fabric Network: To stop the running containers, use:
```
docker-compose down
````


This command stops and removes the containers defined in the docker-compose.yaml file.

#3View Running Containers: To check the status of your running Docker containers, you can use:
```
docker ps
```
#4View Logs: To view the logs of a specific container (e.g., the peer container), use:

```
docker logs <container_name>
```
Replace <container_name> with the actual name of the container. You can get the container names from the output of docker ps.

#5Access a Container’s Shell: To open a shell inside a running container, use:
```
docker exec -it <container_name> /bin/bash
```
This command allows you to interact directly with the container.

#6Clean Up Docker Resources: If you want to remove all stopped containers and unused images, you can run:

```
docker system prune
```
Be cautious with this command, as it will remove any stopped containers and dangling images.

#7Remove Specific Docker Volumes: If your project uses Docker volumes and you want to remove them (for example, to reset the database), you can list all volumes with:

```
docker volume ls
```
Then, remove a specific volume with:

```
docker volume rm <volume_name>
```

#8Rebuild the Docker Images: If you need to rebuild the Docker images (for example, after changing the chaincode), you can use:

```
docker-compose build
```

## Features

- **Create Asset**: Add a new asset with details such as Dealer ID, MSISDN, MPIN, balance, status, transaction amount, transaction type, and remarks.
- **Query Asset**: Retrieve asset information using the Dealer ID.
- **Update Asset**: Update the balance and status of an existing asset.

## Getting Started

To get started with this project, you'll need to have the following installed:

- Go (1.18 or later)
- Hyperledger Fabric binaries
- Docker and Docker Compose


# Smart Contract Functions

```
1 CreateAsset: CreateAsset(dealerId string, msisdn string, mpin string, balance int, status string, transAmount int, transType string, remarks string) error

2 QueryAsset: QueryAsset(dealerId string) (*Asset, error)

3 UpdateAsset: UpdateAsset(dealerId string, balance int, status string) error
```
# Testing

To test the smart contract, use the Hyperledger Fabric test framework or implement unit tests in Go. Ensure that all tests pass before deploying to a production environment.


### Modifications

- **Repository URL**: Replace `git@github.com:jitendra-jitu/SMARTFALCON-FabricAssignment.git` with the actual URL of your GitHub repository.
- **Additional Commands**: If you have more commands or specific scripts, feel free to add them to the `Main Commands` section.
- **License**: Adjust the license section according to the license you choose for your project.

### Conclusion

This `README.md` now provides comprehensive instructions for setting up and running your Hyperledger Fabric project. 
