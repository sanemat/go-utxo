package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func buildCommand(fileName string, commandName string) (string, error) {
	var commandNameWithExt string
	if runtime.GOOS == "windows" {
		commandNameWithExt = commandName + ".exe"
	} else {
		commandNameWithExt = commandName
	}
	builtFile := filepath.Join("out", commandNameWithExt)
	comm := []string{"go", "build", "-o", builtFile, fileName}
	if _, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		return builtFile, err
	}
	return builtFile, nil
}

func buildUtxo() {
	if out, err := buildCommand(filepath.Join("cmd", "utxo", "main.go"), "utxo"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func clean() {
	files := []string{
		filepath.Join("out", "utxo.exe"),
		filepath.Join("out", "utxo"),
	}
	for _, name := range files {
		if exists(name) {
			if err := os.Remove(name); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("rm " + name)
			}
		}
	}
}

func fmtAll() {
	comm := []string{"go", "fmt", "./..."}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func beforeScript() {
	var comm []string
	comm = []string{"go", "get", "golang.org/x/tools/cmd/goimports"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
	comm = []string{"go", "get", "github.com/golang/lint/golint"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func goVet() {
	comm := []string{"go", "vet", "./..."}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func goLint() {
	comm := []string{"golint", "./..."}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else if len(out) != 0 {
		log.Fatal(string(out))
	}
}

func goImports() {
	comm := []string{"goimports", "-l", "./"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else if len(out) != 0 {
		log.Fatal(string(out))
	}
}

func main() {
	action := ""
	if len(os.Args) == 2 {
		action = os.Args[1]
	}

	switch action {
	case "build":
		buildUtxo()
	case "clean":
		clean()
	case "fmt":
		fmtAll()
	case "beforescript":
		beforeScript()
	case "govet":
		goVet()
	case "golint":
		goLint()
	case "goimports":
		goImports()
	case "script":
		goVet()
		goLint()
		goImports()
	default:
		panic("Unknown action: " + action)
	}
}
