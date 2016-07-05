package main

import (
  	"os"
	"fmt"
   "encoding/json"
)

func getSiegfried (fname string, fp *os.File) {
   resp := makeMultipartConnection(POST, siegfried_id, fp, fname)


   var dat map[string]interface{}

   if err := json.Unmarshal([]byte(resp), &dat); err != nil {
      panic(err)
   }

	fmt.Fprintln(os.Stdout, "RESPONSE:", dat)
}
