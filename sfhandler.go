package main

import (
  	"os"
   "bufio"
   "strings"
)

func getSiegfried (fname string, fp *os.File) []string {

   fp.Seek(0,0)
   resp := makeMultipartConnection(POST, siegfried_id, fp, fname)

   sreader := strings.NewReader(resp)
   scanner := bufio.NewScanner(sreader)

   var identifier_count int = 0 
   for scanner.Scan() {
      tmp := scanner.Text()
      if strings.Contains(tmp, "- name") {
         identifier_count+=1
      }
      if strings.Contains(tmp, "filename :") {
         break
      }
    }

   var ids []string

   for scanner.Scan() {
      tmp := scanner.Text()

      if strings.Contains(tmp, "id") {
         tmp := strings.Split(tmp, ":")
         if len(tmp) > 1 && tmp[1] != " " {
            trimmed := strings.Trim(tmp[1], " '")
            ids = append(ids, trimmed)
         }
      }
   }
   return ids
}
