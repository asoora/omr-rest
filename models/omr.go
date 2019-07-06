package models

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func GetAnswers(imgPath string) {
	imgOrigin := gocv.IMRead(imgPath, gocv.IMReadColor)
	// defer imgOrigin.Close()

	blurred := gocv.NewMat()
	defer blurred.Close()

	ksize := image.Point{11, 11}
	gocv.GaussianBlur(imgOrigin, &blurred, ksize, 10, 0, gocv.BorderDefault)

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(blurred, &gray, gocv.ColorBGRToGray)

	normalized := gocv.NewMat()
	defer normalized.Close()
	gocv.Normalize(gray, &normalized, 0, 255, gocv.NormMinMax)

	threshold := gocv.NewMat()
	defer threshold.Close()
	gocv.Threshold(normalized, &threshold, 127, 255, gocv.ThresholdBinary)

	contours := gocv.FindContours(threshold, gocv.RetrievalTree, gocv.ChainApproxSimple)
	fmt.Println(contours)

	window := gocv.NewWindow("omr")
	// defer window.Close()
	window.IMShow(imgOrigin)
}
