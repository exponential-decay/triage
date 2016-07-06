package main

import (
  	"os"
	"fmt"
)

func getTikaId (fp *os.File) {
   resp := makeConnection(PUT, tika_path_detect, fp)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

func getTikaMetadataPUT (fp *os.File) {
   fp.Seek(0,0)
   resp := makeConnection(PUT, tika_path_meta, fp)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

func getTikaMetadataPOST (fname string, fp *os.File) {
   fp.Seek(0,0)
   resp := makeMultipartConnection(POST, tika_path_meta_form, fp, fname)
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}
