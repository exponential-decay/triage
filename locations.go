package main

import (
   "os"
	"fmt"
   "crypto/sha1"
   "encoding/hex"
)

var primedir string = ""

func hashFilepath(path string) string {
   hash := sha1.Sum([]byte(path))
   return hex.EncodeToString(hash[:len(hash)])
}

func makePrimaryDirectory(hash string) {
   primedir = hash
   //http://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value/31151508#31151508
   err := os.Mkdir(hash, 0644)   //http://permissions-calculator.org/
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR: Creating directory,", err)
   }
   fmt.Println("Prime directory:", primedir)
}

func deletePrimaryDirectory() {
   err := os.Remove(primedir)
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR: Removing directory,", err)
   }
}

func createLocs(path string) {
   hash := hashFilepath(path)
   makePrimaryDirectory(hash)
   deletePrimaryDirectory()
}


