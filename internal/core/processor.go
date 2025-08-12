package core

import (
	"fmt"
	"image"
)

// Processor handles the background removal process
type Processor interface {
	ProcessImage(img image.Image) (image.Image, error)
	ProcessBatch(images []image.Image) ([]image.Image, error)
}

// ImageProcessor implements the Processor interface
type ImageProcessor struct {
	modelPath string
	useGPU    bool
}

// NewImageProcessor creates a new image processor instance
func NewImageProcessor(modelPath string, useGPU bool) *ImageProcessor {
	return &ImageProcessor{
		modelPath: modelPath,
		useGPU:    useGPU,
	}
}

// ProcessImage removes background from a single image
func (p *ImageProcessor) ProcessImage(img image.Image) (image.Image, error) {
	// TODO: Implement U²-Net model inference
	// - Preprocess image
	// - Run through model
	// - Apply alpha matting
	// - Return processed image
	return nil, fmt.Errorf("not implemented")
}

// ProcessBatch removes background from multiple images
func (p *ImageProcessor) ProcessBatch(images []image.Image) ([]image.Image, error) {
	results := make([]image.Image, 0, len(images))
	
	for _, img := range images {
		processed, err := p.ProcessImage(img)
		if err != nil {
			return nil, fmt.Errorf("batch processing failed: %w", err)
		}
		results = append(results, processed)
	}
	
	return results, nil
}