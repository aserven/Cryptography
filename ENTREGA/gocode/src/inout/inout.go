package inout

import (
    "bufio"
    "io"
    "os"
    "bytes"
)

func ReadFile(fileToOpen string) []byte {

    fi, err := os.Open(fileToOpen)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    r := bufio.NewReader(fi)
    buf := make([]byte, 1024)
    var fileToCode []byte
    for {
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }
        fileToCode = append(fileToCode, buf[:n]...)
    }
    return fileToCode
}

func WriteFile(fileToWrite []byte, fileName string) {

    r := bytes.NewReader(fileToWrite)    
    fo, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    w := bufio.NewWriter(fo)
    buf := make([]byte, 1024)
    var fileToCode []byte
    for {
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
        fileToCode = append(fileToCode, buf[:n]...)
    }
    if err = w.Flush(); err != nil {
        panic(err)
    }
}