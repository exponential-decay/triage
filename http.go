package main

import (
   "os"
   "fmt"
   "log"
   "strings"
	"net/http"
	"io/ioutil"
)

const PUT = http.MethodPut
const GET = http.MethodGet

func makeConnection (VERB string, request string, fp *os.File) string {

   var stream *http.Request
   var err error

   if fp != nil {
	   stream, err = http.NewRequest(VERB, request, fp) 
   } else {
	   stream, err = http.NewRequest(VERB, request, nil) 
   }

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request.")
	}

	client := &http.Client{}
	meta, err := client.Do(stream)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Error doing request")
	}

	resp, err := ioutil.ReadAll(meta.Body)
	meta.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
   sresp := string(resp)
   resptrimmed := strings.TrimSpace(sresp)   //byte [] becomes string
   return string(resptrimmed)

}
