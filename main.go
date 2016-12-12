package main

import (
  "os/exec"
  "net/http"
  "os"
  "io"
)

func fibbRecc(numIndex int) (numValue int){
  if numIndex == 0 || numIndex == 1 {
    return numIndex
  }

  return fibbRecc(numIndex - 2) + fibbRecc(numIndex -1)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "/index.html")
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "/out0050.png")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/render/", handler)
  http.HandleFunc("/result/", resultHandler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  out, err := os.Create("/model.blend")
  if err != nil {
    panic("error creating")
  }
  defer out.Close()

  r.ParseMultipartForm(32 << 20)
  file, _, err := r.FormFile("file")
  if err != nil {
    panic("error parsing form")
  }
  defer file.Close()
  io.Copy(out, file)
  _, err1 := exec.Command("/bin/bash", "/run.bash").Output()
  if err1 != nil {
    panic("error rendering")
  }
  http.ServeFile(w, r, "/result.html")
}
