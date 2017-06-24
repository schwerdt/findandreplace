package main

import(
  "fmt"
  "os"
  "io/ioutil"
  "bytes"
)

func main() {
  filename := os.Args[1]
  find := os.Args[2]
  replace := os.Args[3]
  fmt.Println("Replace '", find, "' with '", replace, "' in file: ", filename)

  filetext, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("No file found")
  }

  findBytes := []byte(find)
  replaceBytes := []byte(replace)

  // Search for a given string in a file and replace it with another one
  // Read into bytes array.  Can I just replace it with an slice of bytes representing the new string?
  // insert into middle of slice efficiently?
  replaced := bytes.Replace(filetext, findBytes, replaceBytes, -1)
  err = ioutil.WriteFile(filename, replaced, 0644)
  if err != nil {
    fmt.Println("File write was not successful")
  }
}
