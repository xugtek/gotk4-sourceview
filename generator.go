package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"libdb.so/gotk4-sourceview/gensourceview"
)

func main() {
	genmain.Run(genmain.Overlay(
		gendata.Main,
		gensourceview.Data,
	))
}
