package main

import (
   "fmt"
   "log"
   "os"
   "strconv"
)

type result struct {
   numDirs  int64
   numFiles int64
   totalSize int64
}

func dirWorker(dirNames chan string, done chan result) {
   var totalDirs, totalFiles, totalSize int64
   var r result
   for {
      dirName := <-dirNames
      dir, err := os.Open(dirName)
      if err != nil {
         log.Fatalf("failed opening directory: %s", err)
      }
      for {
         list, err := dir.Readdir(256) // 0 to read all files and folders
         if err != nil {
            break
         }
         for _, file := range list {
            totalSize+=file.Size()
            if file.IsDir() {
               totalDirs++
               dirNames <- dirName + "/" + file.Name()
            } else {
               totalFiles++
            }
         }
      }
      dir.Close()
      if len(dirNames) == 0 {
         r.numDirs = totalDirs
         r.numFiles = totalFiles
         r.totalSize= totalSize
         done <- r
         return
      }
   }
}

func main() {
   var numDirW int
   var numName, numDir, totalSize int64
   dirNames := make(chan string, 1048576)
   doneResults := make(chan result)
   dirNames <- os.Args[1]
   numDirW,_ = strconv.Atoi(os.Args[2])
   fmt.Printf("workers = %d\n",numDirW)
   for i := 0; i < numDirW; i++ {
      go dirWorker(dirNames, doneResults)
   }
   numName = 0
   numDir = 0
   for i := 0; i < numDirW; i++ {
      r := <-doneResults
      numName += r.numFiles
      numDir += r.numDirs
      totalSize += r.totalSize
   }
   fmt.Printf("Total Files=%d Total Dirs=%d Total Size=%d\n", numName, numDir,totalSize)
}
