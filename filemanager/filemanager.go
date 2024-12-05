package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath string 
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return nil, errors.New("Error occured while opening the file.")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		fmt.Println("Reading the lines in the file failed.")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("Error occured while reading the lines in the file.")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil{
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		errors.New("Failed to encode data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
