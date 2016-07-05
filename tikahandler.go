package main

import (
  	"os"
	"fmt"
)

func getTikaId (fp *os.File) {
   resp := makeConnection(PUT, tika_path_detect, fp)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

func getTikaMetadata (fp *os.File) {
   resp := makeConnection(PUT, tika_path_meta, fp)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

