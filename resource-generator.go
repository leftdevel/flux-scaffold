package main

import (
    "fmt"
    "os"
)

type ResourceGenerator struct {
    Res ResourceInterface
}

func (rg *ResourceGenerator) Execute() {
    rg.createDirIfNotExists()
    rg.createFileIfNotExists()
}

func (rg *ResourceGenerator) createDirIfNotExists() {
    fullDirPath := GetFullDirPath(rg.Res.GetDir())
    exists, err := PathExists(fullDirPath)

    if err != nil {
        fmt.Println("Unexpected error while checking ", rg.Res.GetDir(), err)
        os.Exit(1)
    }

    if exists == false {
        CreateDir(fullDirPath)
    }
}

func (rg *ResourceGenerator) createFileIfNotExists() bool {
    fullFilePath := GetFullFilePath(rg.Res.GetDir(), rg.Res.GetFileName())
    exists, err := PathExists(fullFilePath)

    if exists || err != nil {
        fmt.Println("File alrady exists, skipping", fullFilePath)
        return false
    }

    file := CreateFile(fullFilePath);
    defer CloseFile(file, fullFilePath);

    WriteFile(file, rg.Res.GetFileContent(), fullFilePath)
    fmt.Println("File", fullFilePath, "created")

    return true
}
