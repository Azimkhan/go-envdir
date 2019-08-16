package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("2 arguments must be supplied")
	}
	envDir := flag.Arg(0)
	programName := flag.Arg(1)

	info, e := os.Stat(envDir)

	if e != nil {
		log.Fatal("File error -", e)
	}

	if !info.IsDir() {
		log.Fatal("File is not a directory")
	}

	err := SetEnvFromDir(envDir)
	cmd := exec.Command(programName)
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Fatal("Program execution error -", err)
	}
}

func SetEnvFromDir(envDir string) error {
	infos, e := ioutil.ReadDir(envDir)
	if e != nil {
		return e
	}
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		// читаем только файлы до 1 кб
		if info.Size() > 1024 {
			continue
		}
		fullPath := filepath.Join(envDir, info.Name())
		bytes, err := ioutil.ReadFile(fullPath)
		if err != nil {
			log.Printf("Warning: unable to read file %s:\n%v", fullPath, err)
		}

		varName := info.Name()
		err = os.Setenv(varName, string(bytes))
		if err != nil {
			log.Printf("Failed to set variable %s:\n%v", varName, err)
		}
	}
	return nil
}
