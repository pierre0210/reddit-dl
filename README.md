# reddit-dl

command line tool for downloading reddit video and convert to other format.

## Build

```bash
go build -o reddit-dl cmd/reddit-dl/reddit-dl.go
```

## Usage

```bash
$ ./reddit-dl -h
Usage of ./reddit-dl:
  -g   Convert to GIF.                                                                                                                  
  -m   Merge video and audio if exists.
  -u string Reddit post url.
```
