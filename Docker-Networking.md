# what is Docker Networking
Docker container in ability for container to communicate and connect with other containers, Container run isolated on the Host and need a way to communicate with the other.

Containers have Networking enabled by default you can followed this command to check it.
```bash
docker network ls
```
```bash
NETWORK ID     NAME      DRIVER    SCOPE
************   bridge    bridge    local
************   host      host      local
************   none      null      local
```
# Why Networking?
There are 2 reason why you need attach Networking to Containers. 
1. You should connect 1 Container to another Container because each of the Containers might have different function like Frontend and Backend Container and obviously these have to talk each other.

2. You might want good isolation of your container example, You have container which contain Payment session then you want isolated this container to your Public accessible container like Login container, this is where you have to isolated you container Because this Payment container there's sensitive Information.

# List Populer Models Networking in Docker 
## Bridge Networking
whenever you install Docker By default you will have Bridge Networking 

<img width="300" height="367" alt="Image" src="https://github.com/user-attachments/assets/6c91273e-388f-44d4-9420-7e5ef243cc6b" />


By default using Bridge network a Container will talk to your Host using ``veth0`` This veth0 are connected using ``Docker0`` network which is a virtual ethernet.

If you want make your container secure and isolate them from the default Bridge Network you can also create your own Bridge Network.
```bash
docker network create -d bridge my_bridge 
```
Run this command To see your network has installed.
```bash
docker network ls

NETWORK ID          NAME                DRIVER
xxxxxxxxxxxx        bridge              bridge
xxxxxxxxxxxx        my_bridge           bridge
xxxxxxxxxxxx        none                null
xxxxxxxxxxxx        host                host
```
You can attached This new Network to the containers. just follow this command
```bash
docker run -d  --name db --net=my_bridge nginx/latest
```

This way, you can run multiple containers on a single host platform where one container is attached to the default network and the other is attached to the my_bridge network.

These containers are completely isolated with their private networks and cannot talk to each other.

<img width="720" height="448" alt="Image" src="https://github.com/user-attachments/assets/e6485f7d-ad01-440a-b6f2-de9407591d33" />

you can also enable to connect your container to ``my_bridge`` or anything you have and communicate each other

```bash
docker connect network my_bridge web
```

## Host Networking
When you use this option the container has access to the Host's Network stack, and shares the Host's network name space. This means that this container will use the same networking configuration and IP address as the Host. 

Example run Container using Host network:
```bash
docker run -d --name test-host --network="host" nginx:latest
```

Keep in mind when you use this option (Host network) your container is less isolated from the Host's system. and has access to all network resources,Because It uses Host's IP address. It can perform any potential security risk, so use this with caution.

## Overlay Network
This enables communication between container accross different Docker Host machine. Overlay Network quite technical It's often use with Orchertration Tolls such as Kuberenetes or Docker swarm. If you only use Docker to containerized your app you can use Bridge and Host based on your need (often used is Bridge)
