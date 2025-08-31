# Volume 
A Docker volume is persistent storage which Docker uses to stored data outside the container it can be (S3. EC2, EFS, etc.)
basically Docker is ephemeral in nature, normally when you create or run Docker container all data stores in that container's **filesystem**
If container goes down your data will die

## Solution 
- Bind
- Volume

<img width="549" height="456" alt="Image" src="https://github.com/user-attachments/assets/028cce79-8aab-4cfb-b38e-13b581db1354" />

Basically using Bind mount you can bind container in docker with file (``/app``) inside your HOST/OS. you can access your container using this file's on the HOST (the name can be change)

so whenever the container goes down you can access through your directory on the HOST. whenever a new container comes up (C2) or you create new container you can Bind that container in the same folder in your Host


In this concept of Bind that means you Bind spesific Directory on the Host to specific Directory on the HOST

## Volume

A volume solves this by storing data in a separate place managed by Docker, so it survives container deletion or recreation. 

Volume is kind of the same solution but volume has better Lifecycle, when we compare with Bind mount your trying to bind 1 Directory on the Host to another specific Docker container, so you're only restricted to specific Host. 

whereas If you are using volume you can create volume in any place like: you can create on the same Host, or on external ec2 instance, S3, NFS. 

you're not restricted with the options, This is the Advantage using volume where you can use external source.

## Lifecycle of docker volume
To generate a new volume follow this command:
```bash
docker volume create <<volume_name>>
```
Once you created ther's a logical Folder or File system that is created on this specific Host which you can't see.

To list all of the volume run this command 
```bash
docker volume ls 
```
Output
```bash
docker volume ls
DRIVER    VOLUME NAME
local     minikube
local     yoga
```
With this command you will see the the volume and you can dedicate this volume to 0 container to another container 

To get all the details about the volume that you've create you can use this command:
```bash
docker volume inspect <<volume_name>>
```

To delete the volume follow this command:
```bash
docker volume rm <<volume_name>>
```
you can also delete multiple volume just added another volume: ``docker volume rm <<name_the_volume>> <<name_the_volume>>

## mount a volume to specific container
First create a docker images
create docker images of your application
```bash
docker build -t name-image .
```

Then activate the Image:
```bash
docker run -d <<image_name>>
```

Create a volume that you want to dedicate.
```bash
docker volume create example
```
once volume is created you can mount it into container using -v or --mount when running a Docker run command
```bash
docker run -d --mount source=example,target=/docker name_image
```
### Explanation of the command above 
- ``docker run -d`` where Im running docker in detach mode
- ``--mount`` it will give me a verbose option 
- ``source`` source is the name of volume
- ``target`` to stored the volume on the specific folder 
- ``name_image`` this is the name image you want to store in volume

## Chek the images has mounted
chek docker which are running
```bash
docker ps 
```
Then inspeck the container
```bash
docker inspect <<container_ID>>
```
You will see that docker has mounted
Output
```bash
      "Mounts": [
            {
                "Type": "volume",
                "Name": "example",
                "Source": "/var/lib/docker/volumes/example/_data",
                "Destination": "/docker",
                "Driver": "local",
                "Mode": "z",
                "RW": true,
                "Propagation": ""
```

If you want to provide additioanl detail In the command ``docker run -d --mount source=example,target=/app name_image`` you can keep adding this using the , (comma)

If you want to delete the volume you firstly deactive the container then follow this command ``docker volume rm <<name_volume>>``.
