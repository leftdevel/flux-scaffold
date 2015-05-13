package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }

    return false, err
}

func GetFullDirPath(dir string) string {
    return GetRootDir() + dir
}

func GetFullFilePath(dir, fileName string) string {
    return GetRootDir() + dir + string(filepath.Separator) + fileName
}

func CreateDir(fullDirPath string) {
    os.Mkdir(fullDirPath, 0777)
}

func GetRootDir() string {
    return "." + string(filepath.Separator)
}

func CreateFile(fullFilePath string) *os.File {
    f, err := os.Create(fullFilePath)
    if err != nil {
        panic(err)
    }
    return f
}

func WriteFile(f *os.File, content string, fullFilePath string) {
    fmt.Fprintln(f, content)
}

func CloseFile(f *os.File, fullFilePath string) {
    f.Close()
}
