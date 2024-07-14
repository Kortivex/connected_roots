package uploads

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
)

// SaveUploadedImage function to handle file uploads.
func SaveUploadedImage(file *multipart.FileHeader, path string, maxWidth, maxHeight uint) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	img, format, err := image.Decode(src)
	if err != nil {
		return err
	}

	width := uint(img.Bounds().Dx())
	height := uint(img.Bounds().Dy())
	if width > maxWidth {
		height = uint(float64(height) * float64(maxWidth) / float64(width))
		width = maxWidth
	}
	if height > maxHeight {
		width = uint(float64(width) * float64(maxHeight) / float64(height))
		height = maxHeight
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	switch format {
	case "jpeg":
		if err = jpeg.Encode(out, resizedImg, nil); err != nil {
			return err
		}
	case "png":
		if err = png.Encode(out, resizedImg); err != nil {
			return err
		}
	case "gif":
		if err = gif.Encode(out, resizedImg, nil); err != nil {
			return err
		}
	default:
		return fmt.Errorf("image format not supported: %s", format)
	}

	if _, err = io.Copy(out, src); err != nil {
		return err
	}

	return nil
}

// DeleteUploadedImage function to handle file uploads deletion.
func DeleteUploadedImage(path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("error deleting '%s': %w", path, err)
	}
	return nil
}
