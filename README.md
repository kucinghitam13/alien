# alien
Alien is a simple service that I use in my LINE's bot to scrape and parse lecture schedule from web page of [baak.gunadarma.ac.id
](https://baak.gunadarma.ac.id/jadwal/cariJadKul)

## How to run
### On host machine
1. Install golang from https://golang.org/doc/install
2. Run `go get -v -u github.com/golang/dep/cmd/dep` to install dep as golang's dependency manager
3. Run `dep ensure -v` to install required dependencies
4. Run by either directly running from source:
```
go run cmd/alien/app.go
```
or by build the binary file first from Makefile:
```
make gobuildalien
make gorunalien
```
### On docker
Or by simply using docker:
1. Build
```
docker build -t alien-img .
```

2.Run
```
docker run --name alien --network host -d alien-img
```

## Endpoint
You can search "jadwal kuliah" from
```
localhost:1981/api/jadwal-kuliah?query=your_class
```
