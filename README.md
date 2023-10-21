<h1 align="center">
  <br>
  <a href="https://docs.ipfs.tech/how-to/command-line-quick-start/"><img src="https://en.wikipedia.org/wiki/File:ADGPI_Indian_Army.svg" alt="EHR logo" title="EHR logo" width="200"></a>
  <br>
  Kubo: IPFS Implementation in GO For EHR
  <br>
</h1>

<p align="center" style="font-size: 1.2rem;">  Forked from [Kubo-IPFS](https://github.com/ipfs/kubo).</p>

<hr />


#### Using With Docker Command 

Clone the Repository:


```console
git clone https://github.com/ipfs/kubo
cd kubo-ipfs
```

Build the Docker Image:

```console
docker build -t ipfs_host .
```

Run the docker container:

```console
docker run -d --name ehr-ipfs-container -v /export -v /data/ipfs -p 4001:4001 -p 4001:4001/udp -p 127.0.0.1:8080:8080 -p 127.0.0.1:5001:5001 ipfs_host
```

Access IPFS Commands:
```console
docker exec -it ehr-ipfs-container ipfs --version
```

Watch the ipfs log:
```console
docker logs -f ehr-ipfs-container
```

or Open docker container:

 (get your container-id)
```console
docker ps        
```

execute the container:
```console
docker exec -it container-id sh
```

check ipfs version:
```console
ipfs --version
```




