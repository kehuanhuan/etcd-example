package main

import (
    "flag"
    "fmt"
    "../discovery"
)

func main() {
    var role       = flag.String("role", "", "master | worker")
    var workerName = flag.String("workerName", "localhost", "this is worker name")
    flag.Parse()
    endpoints := []string{"http://172.16.238.101:2379"}
    if *role == "master" {
        master := discovery.NewMaster(endpoints)
        master.WatchWorkers()
    } else if *role == "worker" {
        worker := discovery.NewWorker(*workerName, "127.0.0.1", endpoints)
        worker.HeartBeat()
    } else {
        fmt.Println("example -h for usage")
    }
}
