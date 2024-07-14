package controllers

import (
	"bytes"
	"image"
	"image-processing/utils"
	"image-processing/views/components/images"
	"image-processing/views/components/text"
	"image-processing/views/pages"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendComponent(w, r, pages.IndexPage())
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	file, _, err := r.FormFile("file")
	if err != nil {
		utils.SendComponent(w, r, text.ErrorText("Error uploading file"))
		return
	}
	defer file.Close()

	// Read the file's content into a buffer
	var fileBuf bytes.Buffer
	tee := io.TeeReader(file, &fileBuf)

	// Decode the image
	decodedImage, format, err := image.Decode(tee)
	if err != nil {
		utils.SendComponent(w, r, text.ErrorText("Error decoding image"))
		return
	}

	beforeFileBytes := fileBuf.Bytes()

	var buf bytes.Buffer
	switch format {
	case "jpeg":
		err = jpeg.Encode(&buf, decodedImage, &jpeg.Options{Quality: 1}) // Adjust quality as needed
	case "png":
		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(&buf, decodedImage)
	default:
		utils.SendComponent(w, r, text.ErrorText("Unsupported image format"))
		return
	}

	if err != nil {
		utils.SendComponent(w, r, text.ErrorText("Error encoding image"))
		return
	}

	before, after := utils.ToBase64(beforeFileBytes, format), utils.ToBase64(buf.Bytes(), format)

	utils.SendComponent(w, r, images.ImageResult(images.ImageResultProps{
		Before: before,
		After:  after,
	}))

}
