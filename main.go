package main

import (
  "os"
  "os/exec"
  "fmt"
  "io"
)

func main() {
  args := os.Args[1:]
  gopath := os.Getenv("GOPATH")
  path := gopath + "/src/github.com/ryannedolan/go-dark/go-dark"
  cmd := exec.Command(path, args...)
  if stdout, err := cmd.StdoutPipe(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  } else {
    go io.Copy(os.Stdout, stdout)
  }
  if stderr, err := cmd.StderrPipe(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  } else {
    go io.Copy(os.Stderr, stderr)
  }
 
  if err := cmd.Run(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
  fmt.Println("done.")
}
