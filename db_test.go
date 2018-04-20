package main

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db := Database{}
	// rows = db.Query("SELECT * FROM Route")
	fmt.Println(db.Query("SELECT * FROM Route WHERE number LIKE \"%1%\""))
	fmt.Println(db.Query("SELECT * FROM Route WHERE number LIKE \"%1%\""))
}
