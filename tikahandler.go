package main

import (
  	"os"
	"fmt"
	"encoding/json"
)

//fl: denotes file level metadata keys
var fl_available_md_keys []string
var fl_keys_values map[string]interface{}

func getTikaId (fp *os.File) {
   resp := makeConnection(PUT, tika_path_detect, fp, "")
	fmt.Fprintln(os.Stdout, "RESPONSE:", resp)
}

func getTikaMetadataPUT (fp *os.File, accepttype string) string {
   fp.Seek(0,0)
   resp := makeConnection(PUT, tika_path_meta, fp, accepttype)
	return resp
}

func getTikaMetadataPOST (fname string, fp *os.File, accepttype string) string {
   fp.Seek(0,0)
   resp := makeMultipartConnection(POST, tika_path_meta_form, fp, fname, accepttype)
	return resp
}

func readTikaJson (output string, key string) {

	var tikamap map[string]interface{}
	if err := json.Unmarshal([]byte(output), &tikamap); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Handling TIKA JSON.")
	}

	getTikaKeys(tikamap)
	fl_keys_values = tikamap

	fmt.Println("\n", output, "\n")
} 

func getTikaKeys (tikamap map[string]interface{}) {	
	keys := make([]string, len(tikamap))
	i := 0
	for k := range tikamap {
		keys[i] = k
		i++
	}

	//replaces /meta/{field} TIKA URL to guarantee key existance
	fl_available_md_keys = keys
}

