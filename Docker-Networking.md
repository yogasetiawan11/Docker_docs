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


## Host Networking
When you use this option the container has access to the Host's Network stack, and shares the Host's network name space. This means that this container will use the same networking configuration and IP address as the Host. 

<img width="233" height="217" alt="Image" src="https://github.com/user-attachments/assets/a7eaeaf9-3246-425c-b12a-fd6d00124a10" />

Keep in mind when you use this option your container is less isolated from the Host's system. and has access to all network resources, It can perform any potential security risk, so use this with caution.

To solve the risk you can go with ``Costume Bridge Network `` 
By default using Bridge network a Container will talk to your Host using ``veth0`` This veth0 are connected using ``Docker0`` network which is a virtual ethernet.

Using Costume Bridge Network 

