package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/disintegration/imaging"
)

// GetOrCreateThumbnail descarga y redimensiona una miniatura, ya sea de imagen o video
func GetOrCreateThumbnail(url, name string, isVideo bool) (string, error) {
	thumbPath := filepath.Join("thumbs", name+".jpg")
	if _, err := os.Stat(thumbPath); err == nil {
		return thumbPath, nil // Ya existe
	}

	tempPath := filepath.Join("thumbs", "temp_"+name+".jpg")
	var genErr error

	if isVideo {
		genErr = generateVideoThumbnail(url, tempPath)
	} else {
		genErr = generateImageThumbnail(url, tempPath)
	}

	if genErr != nil {
		return "", genErr
	}

	// Redimensionar
	err := resizeImage(tempPath, thumbPath)
	if err != nil {
		return "", err
	}

	// Eliminar temp
	os.Remove(tempPath)

	return thumbPath, nil
}

func generateImageThumbnail(url, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, img, nil)
}

func generateVideoThumbnail(url, outputPath string) error {
	cmd := exec.Command("ffmpeg",
		"-ss", "00:00:01",
		"-i", url,
		"-frames:v", "1",
		"-q:v", "2",
		outputPath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w", err)
	}

	return nil
}

func resizeImage(inputPath, outputPath string) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	thumb := imaging.Resize(img, 300, 0, imaging.Lanczos)

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, thumb, nil)
}
