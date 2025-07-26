# PicGeon 

A shitty "Gallery" made in Go that serves pictures from an index like Apache or NGINX and generates thumbnails.

## Requirements

- [Go](https://golang.org/) 1.18 or newer
- [ffmpeg](https://ffmpeg.org/) installed and available in your `$PATH` (for video thumbnails)

## Installation

Clone the repository and run:

```bash
git clone https://github.com/ImnotEdMateo/picgeon.git
cd picgeon 
go run main.go
```

> [!NOTE]
> Make sure that `BASE_URL` and `PICGEON_PORT` are defined.

## TO-DO 

- [ ] Add a CSS
- [ ] Make a decent layout for the thumbnails
- [ ] Make customizable the size of the thumbnails
