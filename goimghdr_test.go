// Copyright 2018 The goimghdr Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goimghdr

import (
	"testing"
)

type imgHdrTest struct {
	Path     string
	Expected string
	Err      error
}

var goimghdrTests = []imgHdrTest{
	{
		Path:     "./imgData/gopher.jpeg",
		Expected: "jpeg",
	},
	{
		Path:     "./imgData/gopher.bmp",
		Expected: "bmp",
	},
	{
		Path:     "./imgData/gopher.gif",
		Expected: "gif",
	},
	{
		Path:     "./imgData/gopher.pbm",
		Expected: "pbm",
	},
	{
		Path:     "./imgData/gopher.png",
		Expected: "png",
	},
	{
		Path:     "./imgData/gopher.ras",
		Expected: "rast",
	},
	{
		Path:     "./imgData/gopher.tiff",
		Expected: "tiff",
	},
	{
		Path:     "./imgData/gopher.exr",
		Expected: "exr",
	},
	{
		Path:     "./imgData/gopher.pgm",
		Expected: "pgm",
	},
	{
		Path:     "./imgData/gopher.ppm",
		Expected: "ppm",
	},
	{
		Path:     "./imgData/gopher.sgi",
		Expected: "rgb",
	},
	{
		Path:     "./imgData/gopher.xbm",
		Expected: "xbm",
	},
	{
		Path:     "./imgData/gopher.webp",
		Expected: "webp",
	},
	{
		Path:     "./imgData/DSCN0029.jpg",
		Expected: "jpeg",
	},
}

var invalidTests = []imgHdrTest{
	{
		Path:     "./imgData/NONE",
		Expected: "",
	},
	{
		Path:     "./imgData/dummy.ee",
		Expected: "",
	},
}

func TestGoImghdr(t *testing.T) {
	for _, ght := range goimghdrTests {
		path := ght.Path
		expected := ght.Expected
		ret, err := What(path)
		if err != nil {
			t.Fatal(err)
		}

		if ret != expected {
			t.Fatalf("What(%s): got %s, want %s", path, ret, expected)
		}
	}

	t.Logf("total %d imghdr tests are passed", len(goimghdrTests))
}

func TestInvalid(t *testing.T) {
	for _, it := range invalidTests {
		path := it.Path
		expected := it.Expected
		ret, err := What(path)
		if err == nil {
			t.Fatalf("Error should be occurred!")
		}
		if ret != it.Expected {
			t.Fatalf("What(%s): got %s, want %s", path, ret, expected)
		}
	}

	t.Logf("total %d invalid tests are passed", len(invalidTests))
}
