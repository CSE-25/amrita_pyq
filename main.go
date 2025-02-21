package main

import (
	"amrita_pyq/cmd/root"
	"amrita_pyq/cmd/semChoose"
	"amrita_pyq/cmd/semTable"
	"amrita_pyq/cmd/year"
)

func main() {
	rootInter := &root.UseInterface{}

	// Initialize semTable and semChoose with the Navigator instance
	semTable.Init(rootInter)
	semChoose.Init(rootInter)
	year.Init(rootInter)

	root.Execute()
}
