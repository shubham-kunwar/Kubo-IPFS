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
git clone https://github.com/shubham-kunwar/kubo-ipfs
cd kubo-ipfs
```

Build the Docker Image:

```console
docker build -t ipfs-image .
```

Run the docker container:

```console
docker run -d -p 4001:4001 --name ehr-ipfs-container ipfs-image
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
```console
docker ps         (get your container-id)
```

```console
docker exec -it container-id sh
```

check ipfs version:
```console
ipfs--version
```




