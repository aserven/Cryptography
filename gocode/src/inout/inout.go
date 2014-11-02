package inout

import (
    "bufio"
    "io"
    "os"
    "bytes"
)

func ReadFile(fileToOpen string) []byte {
    // open input file
    fi, err := os.Open(fileToOpen)
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // open output file
    /*
    fo, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)
    */

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    var fileToCode []byte
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }
        /*
        // write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }*/
        fileToCode = append(fileToCode, buf[:n]...)

    }

    /*if err = w.Flush(); err != nil {
        panic(err)
    }*/

    return fileToCode
}

func WriteFile(fileToWrite []byte, fileName string) {

    // make a read buffer
    r := bytes.NewReader(fileToWrite)

    // open output file
    
    fo, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)
    
    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    var fileToCode []byte
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }
        
        // write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
        fileToCode = append(fileToCode, buf[:n]...)

    }

    if err = w.Flush(); err != nil {
        panic(err)
    }
}