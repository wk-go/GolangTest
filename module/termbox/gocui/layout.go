// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"github.com/jroimartin/gocui"
	"fmt"
	"bufio"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, 0, int(0.2*float32(maxX)), maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Side Bar"
	}
	if v, err := g.SetView("main", int(0.2*float32(maxX)), 0, maxX, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Main"
	}
	if v, err := g.SetView("cmdline", int(0.2*float32(maxX)), maxY-5, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("cmdline"); err != nil {
			return err
		}
	}
	return nil
}

func send(g *gocui.Gui, v *gocui.View) error {
	msgView,err := g.View("cmdline")
	if err != nil {
		return err
	}
	b := make([]byte,2048)
	bufio.NewReader(msgView).Read(b)
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return err
		}
		//v.Clear()
		fmt.Fprintln(v, string(b))
		return nil
	})
	msgView.Clear()
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)


	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("cmdline", gocui.KeyEnter, gocui.ModNone, send); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
