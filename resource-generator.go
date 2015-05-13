package main

import (
    "fmt"
)

type ResourceGenerator struct {
    Res ResourceInterface
}

func (rg *ResourceGenerator) Execute(race chan<- bool) bool {
    if _, err := rg.createDirIfNotExists(); err != nil {
        fmt.Println("Unexpected error while checking ", rg.Res.GetDir(), err)
        race <- true
        return false
    }

    rg.createFileIfNotExists()
    race <- true
    return true
}

func (rg *ResourceGenerator) createDirIfNotExists() (bool, error) {
    fullDirPath := GetFullDirPath(rg.Res.GetDir())
    exists, err := PathExists(fullDirPath)

    if err != nil {
        return false, err
    }

    if exists == false {
        CreateDir(fullDirPath)

        return true, nil
    } else {
        return false, nil
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
