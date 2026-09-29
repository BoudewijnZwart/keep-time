package main

import (
	"fmt"
	"keep_time/internal/sqlite_db"
)

func main() {
	fmt.Printf("%s", storage.Schema )
	fmt.Print(storage.NewClientRepo())
}
