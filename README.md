# Twitter Pic Downloader

Download images from a twitter profile

## Installation

Twitter credential needed :
```bash
export ACCESS_TOKEN=dlkflksdjfldsjfsdjflksd
export ACCESS_TOKEN_SECRET=dlkflksdjfldsjfsdjflksd
export CONSUMER_KEY=dlkflksdjfldsjfsdjflksd
export CONSUMER_SECRET=dlkflksdjfldsjfsdjflksd
```

clone the repository 

```bash
git clone git@github.com:hermescanuto/twitter-pic-downloader.git
```

Download all dependencies 
```bash
go get ./...
```

## Usage

To run 
```bash
./main
```

The aplication will create a file call files/twitter.json.

Example of profiles
```bash
[
	{"screenname": "Marvel"},
	{"screenname": "DCComics"}
]



