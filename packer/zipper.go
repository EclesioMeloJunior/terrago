package packer

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// RemoveZip remove the zip file
func RemoveZip(zipPath string) error {
	return os.Remove(zipPath)
}

// ZipFiles zip all the search files in an.output directory
func ZipFiles(path string, output string, files []string) (string, error) {

	zipStorage := fmt.Sprintf("%s/.cfbuilds/", path)

	if _, err := os.Stat(zipStorage); os.IsNotExist(err) {
		os.Mkdir(zipStorage, os.ModeDir)
	}

	outPutZipDirectory := fmt.Sprintf("%s/%s.zip", zipStorage, output)
	newZip, err := os.Create(outPutZipDirectory)
	if err != nil {
		return "", err
	}

	defer newZip.Close()

	zipWriter := zip.NewWriter(newZip)
	defer zipWriter.Close()

	baseDir := filepath.Base(path)

	for _, file := range files {

		if err = addFileToZip(zipWriter, file, baseDir); err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return outPutZipDirectory, nil
}

func addFileToZip(writer *zip.Writer, file string, baseDir string) error {
	fileToZip, err := os.Open(file)
	if err != nil {
		return err
	}

	defer fileToZip.Close()

	fileInfo, err := fileToZip.Stat()

	if err != nil {
		return err
	}

	fileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	fileDirToZip, err := extractFileBaseDir(baseDir, file)
	if err != nil {
		return err
	}

	fileHeader.Name = strings.TrimPrefix(fileDirToZip, baseDir)
	fileHeader.Method = zip.Deflate

	zipWriter, err := writer.CreateHeader(fileHeader)

	if err != nil {
		return err
	}

	_, err = io.Copy(zipWriter, fileToZip)

	return err
}

func extractFileBaseDir(referencePath string, fileFullPath string) (string, error) {
	fileRegex := regexp.MustCompile(fmt.Sprintf("%s/(.*)", referencePath))
	match := fileRegex.FindStringSubmatch(fileFullPath)

	if len(match) < 0 {
		return "", errors.New("Cannt mach any string")
	}

	return match[0], nil
}
