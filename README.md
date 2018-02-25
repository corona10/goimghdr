[![Build Status](https://travis-ci.org/corona10/goimghdr.svg?branch=master)](https://travis-ci.org/corona10/goimghdr)
[![GoDoc](https://godoc.org/github.com/corona10/goimghdr?status.svg)](https://godoc.org/github.com/corona10/goimghdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/corona10/goimghdr)](https://goreportcard.com/report/github.com/corona10/goimghdr)

# goimghdr
> Inspired by Python's [imghdr](https://docs.python.org/3/library/imghdr.html)

## Installation
```
go get github.com/corona10/goimghdr
```

## List of return value

| Value  | Image format                      |
|--------|-----------------------------------|
| "rgb"  | SGI ImgLib Files                  |
| "gif"  | GIF 87a and 89a Files             |
| "pbm"  | Portable Bitmap Files             |
| "pgm"  | Portable Graymap Files            |
| "ppm"  | Portable Pixmap Files             |
| "tiff" | TIFF Files                        |
| "rast" | Sun Raster Files                  |
| "xbm"  | X Bitmap Files                    |
| "jpeg" | JPEG data in JFIF or Exif formats |
| "bmp"  | BMP files                         |
| "png"  | Portable Network Graphics         |
| "webp" | WebP files                        |
| "exr"  | OpenEXR Files                     |

## Usage
```go
package main

import (
        "fmt"

        "github.com/corona10/goimghdr"
)

func main() {
        ret, _ := goimghdr.What("filepath/sample.jpeg")
        fmt.Println(ret)
}
```

## Special thanks to
* [Haeun Kim](https://github.com/haeungun/)

## Reference
* [CPython imghdr module](https://docs.python.org/3/library/imghdr.html)
* [DSCN0029.jpg](imgData/DSCN0029.jpg) is from [ianare/exif-samples](https://github.com/ianare/exif-samples)
