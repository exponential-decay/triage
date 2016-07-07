package main 

import (
   "io"
  	"os"
	"fmt"
   "crypto/sha1"
   "encoding/hex"
)

func hashFile(fp *os.File) {
  hasher := sha1.New()
   if _, err := io.Copy(hasher, fp); err != nil {
      fmt.Fprintln(os.Stderr, "Error hashing object,", err)
   }
   fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}

func bulkfilehandler(fp *os.File, fi os.FileInfo) string {

   hashFile(fp)
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
