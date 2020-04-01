package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

var count int

func main() {
	//dir :=
	dir := flag.String("dir", "/home/muse/Pictures/target",
		"points to the directory with target images, should be ordered")
	dest := flag.String("dest", "./giffer.gif", "creates .gif file to write to")
	help := flag.Bool("help", false, "prints help message")
	flag.Parse()

	if *help {
		fmt.Println("The program is a gif generator. The following flags should help")
		flag.PrintDefaults()
		return
	}

	// load image palettes
	dirF, err := os.Open(*dir)
	handErr(err)
	count++ // 1

	info, _ := dirF.Stat()
	if !info.IsDir() {
		count++ // 2
		fmt.Println(count)
		log.Fatal("err: provided arguement is not a directory")
		return
	}

	filePaths, err := dirF.Readdirnames(0)
	count++ // 2
	handErr(err)
	dirF.Close()

	var (
		img    image.Image
		ourGif = &gif.GIF{
			LoopCount: 45,
		}
	)

	for _, filePath := range filePaths {
		imgF, err := os.Open(*dir + "/" + filePath)
		count++ // 3
		handErr(err)

		img, _,err = image.Decode(imgF)
		count++ // 4
		handErr(err)

		paletted := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.FloydSteinberg.Draw(paletted, img.Bounds(),img, image.Point{})

		ourGif.Image = append(ourGif.Image, paletted)
		ourGif.Delay = append(ourGif.Delay, 0)
	}

	// generic.gif
	genGif, err := os.Create(*dest)
	count++ // 5
	handErr(err)

	gif.EncodeAll(genGif, ourGif)
	genGif.Close()
}

func handErr(err error) {
	if err != nil {
		fmt.Println(count)
		log.Fatal(err)
	}
}
