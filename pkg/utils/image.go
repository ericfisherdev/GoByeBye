package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// LoadImage loads an image from file
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()
	
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}
	
	_ = format // Format can be used for logging
	return img, nil
}

// SaveImage saves an image to file
func SaveImage(img image.Image, path string, quality int) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()
	
	ext := strings.ToLower(filepath.Ext(path))
	
	switch ext {
	case ".png":
		return png.Encode(file, img)
	case ".jpg", ".jpeg":
		opts := &jpeg.Options{Quality: quality}
		return jpeg.Encode(file, img, opts)
	default:
		return fmt.Errorf("unsupported image format: %s", ext)
	}
}

// ResizeImage resizes an image to the specified dimensions
func ResizeImage(img image.Image, width, height int) image.Image {
	// TODO: Implement image resizing
	// This is a placeholder - actual implementation would use
	// a proper image processing library
	return img
}

// GetImageDimensions returns the width and height of an image
func GetImageDimensions(img image.Image) (int, int) {
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy()
}

// IsImageFile checks if a file path represents a supported image format
func IsImageFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	supportedFormats := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff", ".webp"}
	
	for _, format := range supportedFormats {
		if ext == format {
			return true
		}
	}
	return false
}