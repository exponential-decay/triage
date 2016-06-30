package main

import (
   "os"
   "fmt"
   "strings"
	"net/http"
	"io/ioutil"
)

const PUT = http.MethodPut
const GET = http.MethodGet

const CONN_OKAY int8 = 0
const CONN_BAD int8 = 1

func testConnection (request string) int8 {

   conn := CONN_OKAY
	stream, err := http.NewRequest(GET, request, nil) 
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request,", err)
      os.Exit(1)
	}

	client := &http.Client{}
	_, err = client.Do(stream)
	if err != nil {
      conn = CONN_BAD
	}

   return conn
}

func makeConnection (VERB string, request string, fp *os.File) string {

   var stream *http.Request
   var err error

   if fp != nil {
	   stream, err = http.NewRequest(VERB, request, fp) 
   } else {
	   stream, err = http.NewRequest(VERB, request, nil) 
   }

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request,", err)
      os.Exit(1)
	}

	client := &http.Client{}
	response, err := client.Do(stream)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Doing request,", err)
      os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR: Reading response body,", err)
      os.Exit(1)
	}
   
   data_string := string(data)
   trimmed_response := strings.TrimSpace(data_string)   //byte [] becomes string
   return trimmed_response

}
