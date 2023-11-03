package main

import (
	"RnpServer/Db"
)

func main() {
	db := new(Db.Common)

	err := db.Start("user=postgres password=1611 dbname=RollAndPlay sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
