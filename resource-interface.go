package main

type ResourceInterface interface {
    GetDir() string
    GetFileName() string
    GetFileContent() string
}