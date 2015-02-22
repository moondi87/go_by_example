package main

import "os"
import "fmt"

func main() {
	f := createFile("")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "date")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
