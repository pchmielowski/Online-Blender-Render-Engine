package main

import (
  "os"
  "os/exec"
  "net/http"
  "html/template"
  "io"
  "io/ioutil"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/render/", renderHandler)
  http.HandleFunc("/result/", resultHandler)
  http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "/index.html")
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
  name := "/out0050.png"
  http.ServeFile(w, r, name)
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
  saveModel(r)
  render()
  showResult(w)
}

func saveModel(r *http.Request) {
  modelFile, err := os.Create("/model.blend")
  if err != nil {
    panic("error creating")
  }
  defer modelFile.Close()
  r.ParseMultipartForm(32 << 20)
  file, _, err := r.FormFile("file")
  if err != nil {
    panic("error parsing form")
  }
  defer file.Close()
  io.Copy(modelFile, file)
}

func showResult(w http.ResponseWriter) {
  time, timeErr := ioutil.ReadFile("/time")
  if timeErr != nil {
    panic(timeErr)
  }
  templ, templErr := template.ParseFiles("/result.html")
  if templErr != nil {
    panic(templErr)
  }
  templ.Execute(w, string(time))
}

func render() {
  command := exec.Command("/bin/bash", "/run.bash")
  err := command.Run()
  if err != nil {
    panic(err)
  }
}