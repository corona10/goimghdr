// Copyright 2018 The goimghdr Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimghdr

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	jpeg = "jpeg"
	png  = "png"
	gif  = "gif"
	tiff = "tiff"
	rgb  = "rgb"
	pbm  = "pbm"
	pgm  = "pgm"
	ppm  = "ppm"
	rast = "rast"
	xbm  = "xbm"
	bmp  = "bmp"
	exr  = "exr"
	webp = "webp"
)

// WhatFromReader returns a string value of image header from Reader.
func WhatFromReader(r io.Reader) (string, error) {
	return whatImpl(r)
}

// What returns a string value of image header.
func What(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return whatImpl(f)
}

func whatImpl(r io.Reader) (string, error) {
	var what string
	hdr := make([]byte, 32)
	nBytes, err := r.Read(hdr)
	if err != nil {
		return "", err
	}
	if nBytes < 32 {
		msg := fmt.Sprintf("The target file size should be at least 32 bytes: %d bytes", nBytes)
		return "", errors.New(msg)
	}
	h := string(hdr)
	switch {
	case isJpeg(h):
		what = jpeg
	case isExif(h):
		what = jpeg
	case isGif(h):
		what = gif
	case isTiff(h):
		what = tiff
	case isRgb(h):
		what = rgb
	case isPbm(h):
		what = pbm
	case isPgm(h):
		what = pgm
	case isPng(h):
		what = png
	case isPpm(h):
		what = ppm
	case isRast(h):
		what = rast
	case isXbm(h):
		what = xbm
	case isBmp(h):
		what = bmp
	case isExr(h):
		what = exr
	case isWebp(h):
		what = webp
	}
	return what, nil
}

func stringIn(s string, targets ...string) bool {
	ret := false
	for _, t := range targets {
		if s == t {
			return true
		}
	}
	return ret
}

func isJpeg(h string) bool {
	return h[6:10] == "JFIF"
}

func isExif(h string) bool {
	return h[6:10] == "Exif"
}

func isPng(h string) bool {
	return h[:8] == "\211PNG\r\n\032\n"
}

func isGif(h string) bool {
	header := h[:6]
	return stringIn(header, "GIF87a", "GIF89a")
}

func isTiff(h string) bool {
	return stringIn(h[:2], "MM", "II")
}

func isRgb(h string) bool {
	return h[:2] == "\001\332"
}

func isPbm(h string) bool {
	if len(h) >= 3 && h[0] == 'P' &&
		(h[1] == '1' || h[1] == '4') &&
		(h[2] == ' ' || h[2] == '\t' || h[2] == '\n' || h[2] == '\r') {
		return true
	}
	return false
}

func isPgm(h string) bool {
	if len(h) >= 3 && h[0] == 'P' &&
		(h[1] == '2' || h[1] == '5') &&
		(h[2] == ' ' || h[2] == '\t' || h[2] == '\n' || h[2] == '\r') {
		return true
	}
	return false
}

func isPpm(h string) bool {
	if len(h) >= 3 && h[0] == 'P' &&
		(h[1] == '3' || h[1] == '6') &&
		(h[2] == ' ' || h[2] == '\t' || h[2] == '\n' || h[2] == '\r') {
		return true
	}
	return false
}

func isRast(h string) bool {
	return h[:4] == "\x59\xA6\x6A\x95"
}

func isXbm(h string) bool {
	s := "#define "
	return h[:len(s)] == s
}

func isBmp(h string) bool {
	return h[:2] == "BM"
}

func isWebp(h string) bool {
	return strings.HasPrefix(h, "RIFF") && h[8:12] == "WEBP"
}

func isExr(h string) bool {
	return strings.HasPrefix(h, "\x76\x2f\x31\x01")
}
