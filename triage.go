package main

import (
	"os"
	"fmt"
   "flag"
   "path/filepath"
)

var file string

/*func findOpenConnections() {
   var tika string = "http://127.0.0.1:9998/"
   resp := makeConnection(GET, tika, nil)
   fmt.Fprintln(os.Stdout, resp)
}*/

func init() {
   flag.StringVar(&file, "file", "false", "File to find the distance between.")
}

func sayHello (output bool) {
   if output {
      //to get a hello response from tika...
	   //e.g CURL: curl -X GET http://localhost:9998/tika

      var tika_path_hello string = "http://127.0.0.1:9998/tika"
      resp := makeConnection(GET, tika_path_hello, nil)
	   fmt.Fprintln(os.Stdout, "DETECT:", resp)
   }
}

func getMetadata (fp *os.File) {

   var tika_path_detect string = "http://127.0.0.1:9998/detect/stream"
   //var tika_path_meta string = "http://127.0.0.1:9998/meta"

   resp := makeConnection(PUT, tika_path_detect, fp)
	fmt.Fprintln(os.Stdout, "DETECT:", resp)
}

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {
   
   fp, err := os.Open(path)
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR:", err)
      os.Exit(1)  //should only exit if root is null, consider no-exit
   }

   switch mode := fi.Mode(); {
   case mode.IsRegular():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a file.")
      getMetadata(fp)
   case mode.IsDir():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a directory.")      
   default: 
      fmt.Fprintln(os.Stderr, "INFO: Something completely different.")
   }
   return nil
}

func main() {

   flag.Parse()

   if flag.NFlag() <= 0 {    // can access args w/ len(os.Args[1:]) too
      fmt.Fprintln(os.Stderr, "Usage:  triage [-file ...]")
      fmt.Fprintln(os.Stderr, "Output: [TBD]")
      flag.Usage()
      os.Exit(0)
   }

   //findOpenConnections()

   var test bool = true
   if test {
      sayHello(true)
      filepath.Walk(file, readFile)
   }
}
