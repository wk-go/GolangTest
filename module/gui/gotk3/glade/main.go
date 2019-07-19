package main

import (
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
)

const appId = "com.github.gotk3.gotk3-examples.glade"

func main() {

	// Create a new application.
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		render(application)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func render(application *gtk.Application) {
	log.Println("application activate")

	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("main.glade")
	errorCheck(err)

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"on_main_window_destroy": onMainWindowDestroy,
	}
	builder.ConnectSignals(signals)

	// Get the object with the id of "main_window".
	obj, err := builder.GetObject("main_window")
	errorCheck(err)

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	win, err := isWindow(obj)
	errorCheck(err)

	obj, err = builder.GetObject("button1")
	btn1, err := isButton(obj)
	errorCheck(err)
	btn1.Connect("clicked", func() {
		label, _ := btn1.GetLabel()
		log.Println("click the button:", label)
	})

	obj, err = builder.GetObject("button2")
	btn2, err := isButton(obj)
	errorCheck(err)
	btn2.Connect("clicked", func() {
		dialog, _ := gtk.FileChooserDialogNewWith1Button("Choose a file", win, gtk.FILE_CHOOSER_ACTION_OPEN, "Select", 1)
		res := dialog.Run()
		if res == gtk.RESPONSE_ACCEPT {
			log.Println(res)
			filename := dialog.GetFilename()
			log.Println("filename:", filename)
		}
	})

	// Show the Window and all of its components.
	win.Show()
	application.AddWindow(win)
}

func isButton(obj glib.IObject) (*gtk.Button, error) {
	// Make type assertion (as per gtk.go).
	if o, ok := obj.(*gtk.Button); ok {
		return o, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
