package main

import (
   "log"
   "time"
)

// Every 10 minutes run the cleaner to clean out old logs
func setupDaemon() {
   log.Println("Starting up daemon")
   cleanup()
   ticker := time.NewTicker(10 * time.Minute)
   select {
      case <- ticker.C:
         cleanup()
   }
}

func cleanup() {
   log.Println("Running cleaner")
}
