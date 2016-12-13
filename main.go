package main

import (
  "os/exec"
  "net/http"
  "html/template"
  "os"
  "io"
  "io/ioutil"
)

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

  cmd := exec.Command("/bin/bash", "/run.bash")
  err1 := cmd.Run()
  if err1 != nil {
    panic(err1)
  }

  time, _ := ioutil.ReadFile("/time")
  t, _ := template.ParseFiles("/result.html")
  t.Execute(w, string(time))
}
