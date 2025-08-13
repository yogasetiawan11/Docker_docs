# Volume 
A Docker volume is persistent storage which Docker uses to stored data outside the container, 

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

To generate a new volume follow this command:
```bash
docker volume create <<name_volume>>
```

once volume is created you can mount it into container using -v or --mount when running a Docker run command
```bash
docker run -it -v <volume name>:/data <imagename> /bin/bash
```

```bash
docker volume inspect <<name_volume>>
```
