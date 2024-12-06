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
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the lines in the file failed.")
		fmt.Println(err)
		return nil, errors.New("Error occured while reading the lines in the file.")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil{
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		errors.New("Failed to encode data to JSON")
	}

	fmt.Println("Successfully wrote the results to local files.")
	return nil
}

func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
