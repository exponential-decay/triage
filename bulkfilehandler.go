package main 

import (
  	"os"
	"fmt"
)

func bulkfilehandler(fp *os.File, fi os.FileInfo) string {
	fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a file.")
	ids := getSiegfried(fi.Name(), fp, "")

	getTikaId(fp)

	//_ = getTikaMetadataPOST(fi.Name(), fp, ACCEPT_MIME_JSON)

	//placeholder ids from SF
	for _, id := range ids {
		fmt.Println(id)
	}

	//fmt.Println(fl_available_md_keys)
	//fmt.Println(fl_keys_values)

	_ = getTikaRecursive(fi.Name(), fp, ACCEPT_MIME_JSON)

	//fmt.Println(fl_recursive_keys_values)

	//for _, v := range fl_recursive_md_keys {
	//   fmt.Println(v)
	//}
	//fmt.Println(fl_recursive_md_keys)

	//fmt.Println(test)	
	return "xxx"
}