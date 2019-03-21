/* method only use to Zooms in image now.*/
func ImageCropBin(url string) string {
	ret,e := http.Get(url)
	if e != nil {
		fmt.Println(e)
	}

	input := ret.Body
	defer input.Close()

	var srcImage, name, e1 = image.Decode(input)
	if e1 != nil {
		beego.Info(e1)
	}
	beego.Info(name)
	srcRect := srcImage.Bounds()
	beego.Info(srcRect)


	targetRect := image.Rect(0,0, DefaultMaxWidth, DefaultMaxHeight)
	targetImage := image.NewRGBA(targetRect)

	var ratio float64
	if float32(targetRect.Dx()) / float32(targetRect.Dy()) > float32(srcRect.Dx()) / float32(srcRect.Dy()) {
		ratio = float64(srcRect.Dx()) / float64(targetRect.Dx())
	} else {
		ratio = float64(srcRect.Dy()) / float64(targetRect.Dy())
	}

	if ratio < 1 {
		return ""
	}

	var r int
	r = int(ratio)

	position := 0
	for ;r > 0; {
		r = r >> 1
		if r > 0 {
			position++
		}
	}

	scale := int(math.Pow(2, float64(position)))

	if scale == 0 {
		return ""
	}

	srcWidth := srcRect.Dx()
	srcHeight := srcRect.Dy()

	midWidth := srcRect.Dx() / scale
	midHeight := srcRect.Dy() / scale

	offsetX := (midWidth - DefaultMaxWidth) / 2
	offsetY := (midHeight - DefaultMaxHeight) / 2

	for height := 0; height < targetRect.Dy() ; height ++ {
		for width := 0; width < targetRect.Dx(); width ++ {
			x := (width + offsetX) * scale
			y := (height + offsetY) * scale

			if x > srcWidth {
				x = srcWidth
			}

			if y > srcHeight {
				y = srcHeight
			}

			c := calculateRectRGB(srcImage,x,y,scale)
			targetImage.Set(width,height,c)
		}
	}


	file, e := os.Create("ssssssssss.png")
	if e != nil {
		beego.Debug("1234:",e)
	}

	e = png.Encode(file, targetImage)
	if e != nil {
		beego.Debug("123:",e)
	}

	file, _ = os.Open("ssssssssss.png")
	md := md5.New()
	_, e = io.Copy(md, file)

	// save file.
	sumStr := hex.EncodeToString(md.Sum(nil))
	_ = os.Rename("ssssssssss.png", param.ImagePath + sumStr)

	return sumStr
}

func calculateRectRGB(image image.Image,x int,y int,scale int) color.Color {
	var sumR,sumG,sumB uint64
	sumR,sumG,sumB = 0,0,0
	for height := y; height < y + scale ;height ++  {
		for width := x; width < x + scale ;width ++  {
			// at last row or column,may be out of bound
			r,g,b,_ := image.At(width,height).RGBA()
			sumR += uint64(r)
			sumG += uint64(g)
			sumB += uint64(b)
		}
	}

	factor := uint64(scale * scale * 257)
	c := color.RGBA{R: uint8(sumR  / factor), G: uint8(sumG / factor), B: uint8(sumB / factor), A: 255}
	return c
}
