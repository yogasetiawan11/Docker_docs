
# What is Docker
Docker is an open-source platform for containerization that allows you to automate the deployment, scaling, and management of applications using containerization. Containers package your application code together with all its dependencies, libraries, and configuration files, so it can run consistently across different computing environments.

## Infrastructure of Docker
<img width="1662" height="718" alt="Image" src="https://github.com/user-attachments/assets/81419d6c-375b-43cf-be09-7f7371899be3" />

In this picture when you use model Hypervisor on the right when you're using this model you have to maintain pyshical servers so there's a lot of maintenance there has to be a team who mantain all of this pyshical servers. whenever there are new patches, whenever there are security fixes or whatever is the case. you have to maintain your data center.

Whereas if your using model on the left picture it can be create ec2 intance, or creating VM on top of pyshical servers, using this approach your mantenace is less. unless you own that pyshical server. so many organization has sifted to this approach. because everybody use cloud provider or people looking for to reduce the maintenace or organization are not willing to mantain their own data center so this is the main reason. 

### The difference between Container and Virtual Machine
the diagram above looks simmilar but The major difference is that container are very lightweight in nature, becaue they do not have a complete operating system or docker use resources from the Virtual Machine or pyshical server that you're running on, whereas a VM (the right image) they have a complete operating system, have complete isolation, and security.

So a container is a package which is combination of your application + application library/dependencies + system dependencies (python etc).

# Docker Architecture
<img width="1022" height="550" alt="Image" src="https://github.com/user-attachments/assets/abc7d333-d16d-4810-9ebf-3005fe70d93f" />

This picture clearly indicates that Docker Daemon is a brain of docker, if docker is killed, stops working for some reasons, Docker can be dead

## Docker lifcycle
we can use above image as reference to understand Docker lifecycle
there are 3 main command:
- docker build  -> build docker image from docker file
- docker run  -> run containers from docker images
- docker push  -> push docker image to public/private registries to shared the docker images.

<img width="1355" height="562" alt="Image" src="https://github.com/user-attachments/assets/6c04ff7c-55e7-4769-919e-5fa5d1a90d49" />

## Understanding Terminology in docker (inspired by docker docs)
- ``Docker CLI``	The Docker CLI is the command-line interface for interacting with the Docker Engine. It provides commands like docker run, docker build, docker ps, and others to manage Docker containers, images, and services.

- ``Docker Compose``	Docker Compose is a tool for defining and running multi-container Docker applications using a YAML file (compose.yaml). With a single command, you can start all services defined in the configuration.

- ``Docker Desktop``	Docker Desktop is an easy-to-install application for Windows, macOS, and Linux that provides a local Docker development environment. It includes Docker Engine, Docker CLI, Docker Compose, and a Kubernetes cluster.

- ``Docker Engine``	Docker Engine is the client-server technology that creates and runs Docker containers. It includes the Docker daemon (dockerd), REST API, and the Docker CLI client.

- ``Docker Hub``	Docker Hub is Docker’s public registry service where users can store, share, and manage container images. It hosts Docker Official Images, Verified Publisher content, and community-contributed images.

- ``base image``	A base image is an image you designate in a FROM directive in a Dockerfile. It defines the starting point for your build. Dockerfile instructions create additional layers on top of the base image. A Dockerfile with the FROM scratch directive uses an empty base image.

- ``build``	Build is the process of building Docker images using a Dockerfile. The build uses a Dockerfile and a "context". The context is the set of files in the directory in which the image is built.

- ``container``	A container is a runnable instance of an image. You can start, stop, move, or delete a container using the Docker CLI or API. Containers are isolated from one another and the host system but share the OS kernel. They provide a lightweight and consistent way to run applications.

- ``image``	 An image is a read-only template used to create containers. It typically includes a base operating system and application code packaged together using a Dockerfile. Images are versioned using tags and can be pushed to or pulled from a container registry like Docker Hub.

- ``registry``	A registry is a storage and content delivery system for Docker images. The default public registry is Docker Hub, but you can also set up private registries using Docker Distribution.

- ``volume``	A volume is a special directory within a container that bypasses the Union File System. Volumes are designed to persist data independently of the container lifecycle. Docker supports host, anonymous, and named volumes.

# Install Docker
[you can see the documentation on live](https://docs.docker.com/get-started/get-docker/)

or 

```bash
sudo apt install docker.io
```
## Verify The Docker has installed
Run this Command
``docker run hello-world ``

The Output would be
```bash
Hello from Docker!
This message shows that your installation appears to be working correctly.
```

## Login With your Docker Hub
Before you push to Docker Hub you must Login to your Docker Hb If it has done you can proceed push your Docker Images

Use this command 

```bash
docker login --username (enter your dockerhub username)
**example**
docker login --username yogasn
```


## Build your first Docker Image
You need to change the username and repo accordingly in the command below
```bash
docker build -t yogas4/my_first_docker:latest .
```

explanation of the Command above 
- ``docker build`` This represent you will build Docker image from Docker File
- ``-t`` Tag to as a markup if you don't use a tag your image will be lost 
- ``yogas4`` This is a username in your Docker registry
- ``/path`` This contain your repo where you will store your Docker file and application
- ``:latest`` This is a version of your Tag where you can say V1, V2 etc
- `` . `` The dot . represent that Dockerfile in the same Directory

Output of the Above command 

```bash
Step 5/6 : ENV NAME world
 ---> Running in 902360b0097b
 ---> Removed intermediate container 902360b0097b
 ---> 50ca3b446592
Step 6/6 : CMD ["python3", "app.py"]
 ---> Running in 220c8d43e71c
 ---> Removed intermediate container 220c8d43e71c
 ---> 1cc893691ed0
Successfully built 1cc893691ed0
Successfully tagged yogas4/my_first_docker:latest
```

## Verify If Docker has created
```bash
docker images
```

Output
```bash
REPOSITORY               TAG                     IMAGE ID       CREATED              SIZE
yogas4/my_first_docker   latest                  1cc893691ed0   About a minute ago   557MB
```

## Run Your First Container
```bash
docker run -it yogas4/my_first_docker:latest
```

Output 

```bash
hello
```

## Stop container that are running
```bash
docker stop CONTAINERID
```

## Push The Image to Docker Hub
```bash
docker push yogas4/my_first_docker
```

Output

```bash
Using default tag: latest
The push refers to repository [docker.io/yogas4/my_first_docker]
5013d20fc7c2: Pushed
7f69ec9bb59a: Pushed
107284f29743: Pushed
107cbdaeec04: Mounted from library/ubuntu
latest: digest: sha256:c6154796d19b6041f35d6dec319f28e053f23ac44251911e62cc804970758a04 size: 1155
```



Finally you have finished this step :v
