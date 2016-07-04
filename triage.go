package main

import (
  	"os"
	"fmt"
   "flag"
   "path/filepath"
   "encoding/json"
)

var file string
var vers bool

func init() {
   flag.StringVar(&file, "file", "false", "File to find the distance between.")
   flag.BoolVar(&vers, "version", false, "[Optional] Output version of the tool.")
}

func findOpenConnections() {
   var tika string = "http://127.0.0.1:9998/"
   resp := testConnection(tika)
   if resp == CONN_BAD {
      fmt.Fprintln(os.Stdout, "INFO: Tika connection not available to connect to.")
   }
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

func getSiegfried (fname string, fp *os.File) {
   resp := makeMultipartConnection(POST, siegfried_id, fp, fname)


   var dat map[string]interface{}

   if err := json.Unmarshal([]byte(resp), &dat); err != nil {
      panic(err)
   }

	fmt.Fprintln(os.Stdout, "RESPONSE:", dat)
}

func getTikaId (fp *os.File) {
   resp := makeConnection(PUT, tika_path_detect, fp)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

func getTikaMetadata (fp *os.File) {
   resp := makeConnection(PUT, tika_path_meta, fp)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
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
      getSiegfried(fi.Name(), fp)
      //getTikaId(fp)
      //getTikaMetadata(fp)
      
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
      fmt.Fprintln(os.Stderr, "               [Optional -version]")
      fmt.Fprintln(os.Stderr, "Output: [TBD]")
      flag.Usage()
      os.Exit(0)
   }

   if vers {
      fmt.Fprintln(os.Stdout, getVersion())
      os.Exit(1)
   }

   //findOpenConnections()

   var test bool = true
   if test {
      //sayHello(true)
      filepath.Walk(file, readFile)
   }
}
