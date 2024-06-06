package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	stat, err := fromFile.Stat()
	if err != nil {
		return err
	}

	fileSize := stat.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit > fileSize-offset {
		limit = fileSize - offset
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()

	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	percentage := 0
	buffer := make([]byte, 1024)

	for {
		n, err := fromFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("failed to red: %v", err)
		}
		if n == 0 {
			break
		}

		_, err = toFile.Write(buffer[:n])
		if err != nil {
			return err
		}

		newOffset, err := toFile.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}

		percentage = int(float64(newOffset-offset) / float64(limit) * 100)
		fmt.Printf("\rCopying... %d%%", percentage)
	}
	fmt.Println("\nCopy completed!")
	return nil
}
