<h1 align="center">
  <br>
  <a href="https://docs.ipfs.tech/how-to/command-line-quick-start/"><img src="https://en.wikipedia.org/wiki/File:ADGPI_Indian_Army.svg" alt="EHR logo" title="EHR logo" width="200"></a>
  <br>
  Kubo: IPFS Implementation in GO For EHR
  <br>
</h1>

<p align="center" style="font-size: 1.2rem;">  Forked from [Kubo-IPFS](https://github.com/ipfs/kubo).</p>

<hr />

## What is Kubo?

Kubo was the first IPFS implementation and is the most widely used one today. Implementing the *Interplanetary Filesystem* - the Web3 standard for content-addressing, interoperable with HTTP. Thus powered by IPLD's data models and the libp2p for network communication. Kubo is written in Go.

## What is IPFS?

IPFS is a global, versioned, peer-to-peer filesystem. It combines good ideas from previous systems such as Git, BitTorrent, Kademlia, SFS, and the Web. It is like a single BitTorrent swarm, exchanging git objects. IPFS provides an interface as simple as the HTTP web, but with permanence built-in. You can also mount the world at /ipfs.

For more info see: https://docs.ipfs.tech/concepts/what-is-ipfs/

### Minimal System Requirements

IPFS can run on most Linux, macOS, and Windows systems. We recommend running it on a machine with at least 4 GB of RAM and 2 CPU cores (kubo is highly parallel). On systems with less memory, it may not be completely stable, and you run on your own risk.

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




