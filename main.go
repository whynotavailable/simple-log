package main

import (
   "flag"
   "log"
   "net/http"
   "os"
)

func main() {
   var mode string

   flag.StringVar(&mode, "mode", "all", "Which mode to use (all, api, daemon)")
   flag.Parse()

   if mode != "all" && mode != "api" && mode != "daemon" {
      flag.Usage()
      os.Exit(1)
   }

   switch mode {
   case "all":
      go setupAPI()
      go setupDaemon()
   case "api":
      go setupAPI()
   case "daemon":
      go setupAPI()
   }

   select {}
}

func setupAPI() {
   http.HandleFunc("/", handler)
   log.Println("Starting up api")
   http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
   writer.Write([]byte("yolo"))
}
