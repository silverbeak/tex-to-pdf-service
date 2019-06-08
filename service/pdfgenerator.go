package pdfgenerator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/satori/go.uuid"
)

func getArgs(texFilename string, outputDirectory string) ([]string, string) {
	identifier, err := uuid.NewV4()
	if err != nil {
		return nil, ""
	}
	workingDirectory := fmt.Sprintf("%v/%v", outputDirectory, identifier)
	output := fmt.Sprintf("%v/work", workingDirectory)
	args := []string{"-synctex=1", "-interaction=nonstopmode", "-file-line-error", "-d", "-ps", "-xelatex", texFilename, fmt.Sprintf("-jobname=%v", output), "-pdflatex=pdflatex"}
	return args, workingDirectory
}

func writeTexToFile(tex []byte, tempdir string) (string, error) {
	identifier, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	texFilename := fmt.Sprintf("%v/%v.tex", tempdir, identifier.String())
	err = ioutil.WriteFile(texFilename, tex, 0644)
	if err != nil {
		return "", err
	}
	return texFilename, nil
}

func readPdfFile(workdir string) ([]byte, error) {
	return ioutil.ReadFile(workdir + "/work.pdf")
}

func makeWorkDir(workdir string) error {
	return os.MkdirAll(workdir, os.ModePerm)
}

func removeWorkDir(workdir string) error {
	return os.RemoveAll(workdir)
}

// Generate Generates a pdf file, using latexmk, reads the result and returns it
func Generate(payload []byte, tempdir string, removeWhenDone bool) ([]byte, error) {
	texFilename, err := writeTexToFile(payload, tempdir)
	if err != nil {
		return []byte{}, err
	}
	command := "latexmk"
	args, workdir := getArgs(texFilename, tempdir)

	err = makeWorkDir(workdir)
	if err != nil {
		return []byte{}, err
	}

	if removeWhenDone {
		defer removeWorkDir(workdir)
	}

	cmd := exec.Command(command, args...)

	err = cmd.Run()
	if err != nil {
		return []byte{}, err
	}

	pdfBytes, err := readPdfFile(workdir)
	if err != nil {
		return []byte{}, err
	}

	log.Println("PDF successfully generated")
	return pdfBytes, nil
}
