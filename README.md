Demo IPFS Client for Golang
==================================

## Requirements
- Go 1.5 or later.
- [IPFS prebuilt package]
- [go-ipfs-api]

## Usage

1. Open IPFS daemon
```bash
ipfs daemon
```

2. Build runnable go app
```bash
go install
```

3. Run go app
```bash
<GOPath>/bin/ipfsclientgo.exe
```

4. Check result on [IPFS gateway]

5. Check result on terminal using IPFS command
```bash
ipfs cat QmQjkqb78AEDXnqNsn7g8BVVmJbsZU6k6VQ2b2pmnipq53
//Error: this dag node is a directory

ipfs ls -v QmQjkqb78AEDXnqNsn7g8BVVmJbsZU6k6VQ2b2pmnipq53
ipfs object links QmQjkqb78AEDXnqNsn7g8BVVmJbsZU6k6VQ2b2pmnipq53
```


[go-ipfs-api]: https://github.com/ipfs/go-ipfs-api
[IPFS prebuilt package]: https://dist.ipfs.io/#go-ipfs
[IPFS gateway]: http://localhost:8080/ipfs/QmQjkqb78AEDXnqNsn7g8BVVmJbsZU6k6VQ2b2pmnipq53
