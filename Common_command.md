# Docker Command
Most Commonly used Docker Commands are

1. Image Management
These commands are used to build, manage, and distribute Docker images.

### docker build

Syntax: docker build [OPTIONS] PATH | URL | -

Explanation: Builds an image from a Dockerfile. The PATH is the location of the Dockerfile and the build context. The command reads the instructions in the Dockerfile and executes them to create a new image.

Example: ```docker build -t my-app:1.0``` . (Builds an image named my-app with tag 1.0 from the current directory).

### docker images

Syntax: docker images [OPTIONS] [REPOSITORY]

Explanation: Lists all the Docker images on your local machine. You can use options to filter the output.

Example: ```docker images```

### docker pull

Syntax: docker pull [OPTIONS] NAME[:TAG|@DIGEST]

Explanation: Pulls an image or a repository from a registry (e.g., Docker Hub).

Example: ```docker pull ubuntu:latest```

### docker push

Syntax: docker push [OPTIONS] NAME[:TAG]

Explanation: Pushes an image or a repository to a registry. You must be authenticated to push to a private registry.

Example: ```docker push my-username/my-app:1.0```

### docker rmi

Syntax: docker rmi [OPTIONS] IMAGE [IMAGE...]

Explanation: Removes one or more images. You must first stop any containers that are using the image.

Example: ```docker rmi ubuntu:latest```

### docker tag

Syntax: docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

Explanation: Creates a tag TARGET_IMAGE that refers to SOURCE_IMAGE. This is often used to give an image a more meaningful name or to prepare it for pushing to a registry.

Example:``` docker tag my-app:1.0 my-username/my-app:production```

2. Container Management
These commands are used to create, start, stop, and manage the lifecycle of containers.

### docker run

Syntax: docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Explanation: Creates and starts a new container from a given image. This is one of the most frequently used commands. It can run a command inside the container and also manage ports, volumes, and other configurations.

Example: ```docker run -d -p 8080:80 --name my-webserver nginx ```(Runs a detached nginx container, maps port 8080 on the host to port 80 in the container, and names it my-webserver).

### docker ps

Syntax: docker ps [OPTIONS]

Explanation: Lists all currently running containers. Use the -a option to show all containers (including stopped ones).

```Example: docker ps -a```

### docker start

Syntax: docker start [OPTIONS] CONTAINER [CONTAINER...]

Explanation: Starts one or more stopped containers.

Example: ```docker start my-webserver```

### docker stop

Syntax: docker stop [OPTIONS] CONTAINER [CONTAINER...]

Explanation: Stops one or more running containers. It sends a SIGTERM signal to the main process in the container, followed by SIGKILL after a timeout.

Example: ```docker stop my-webserver```

### docker restart

Syntax: docker restart [OPTIONS] CONTAINER [CONTAINER...]

Explanation: Restarts one or more containers.

Example: ```docker restart my-webserver```

### docker rm

Syntax: docker rm [OPTIONS] CONTAINER [CONTAINER...]

Explanation: Removes one or more containers. You must first stop a container before you can remove it, or use the -f flag to force removal.

Example: ```docker rm my-webserver```

### docker exec

Syntax: docker exec [OPTIONS] CONTAINER COMMAND [ARG...]

Explanation: Runs a command inside a running container. This is useful for debugging or interacting with a container's environment.

Example:``` docker exec -it my-webserver bash``` (Opens an interactive Bash shell inside the my-webserver container).

### docker logs

Syntax: docker logs [OPTIONS] CONTAINER

Explanation: Fetches the logs of a container. The -f option can be used to follow the log output in real-time.

Example: ```docker logs -f my-webserver```

3. Network Management
These commands are used to manage Docker networks.

### docker network ls

Syntax: docker network ls [OPTIONS]

Explanation: Lists all Docker networks.

Example: ```docker network ls```

### docker network create

Syntax: docker network create [OPTIONS] NETWORK

Explanation: Creates a new network. This is useful for creating isolated networks for your application's services.

Example: ```docker network create my-app-net```

### docker network connect

Syntax: docker network connect [OPTIONS] NETWORK CONTAINER

Explanation: Connects a container to a network.

Example: ```docker network connect my-app-net my-db-container```

### docker network rm

Syntax: docker network rm [OPTIONS] NETWORK [NETWORK...]

Explanation: Removes one or more networks.

Example: ```docker network rm my-app-net```

4. Volume Management
These commands are used to manage Docker volumes for persistent data.

### docker volume ls

Syntax: docker volume ls [OPTIONS]

Explanation: Lists all Docker volumes.

Example: ```docker volume ls```

### docker volume create

Syntax: docker volume create [OPTIONS] [VOLUME]

Explanation: Creates a new volume.

Example: ```docker volume create my-app-data```

### docker volume rm

Syntax: docker volume rm [OPTIONS] VOLUME [VOLUME...]

Explanation: Removes one or more volumes.

Example:``` docker volume rm my-app-data```

5. System Management
These commands are used for overall Docker system and environment management.

### docker info

Syntax: docker info [OPTIONS]

Explanation: Displays system-wide information. This is useful for checking the Docker version, storage driver, and other system details.

Example: ```docker info```

### docker system prune

Syntax: docker system prune [OPTIONS]

Explanation: Removes unused Docker objects (containers, images, networks, and volumes). This is a great way to free up disk space. Use with caution!

Example: ```docker system prune -a ```(Removes all unused images, not just dangling ones).
or ```docker image prune``` or ```docker image prune -a ``` 

### docker version

Syntax: docker version [OPTIONS]

Explanation: Shows the Docker version information.

Example: ```docker version``` 
