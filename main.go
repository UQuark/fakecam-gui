package main

import (
	"os"

	"github.com/UQuark0/fcbe"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

const (
	gladeFile          = "gtk/main_window.glade"
	invalidTypecastErr = "Invalid typecast"
)

var (
	wnMain        *gtk.Window
	btChooseFile  *gtk.FileChooserButton
	cbLoopForever *gtk.CheckButton
	sbLoopCount   *gtk.SpinButton
	btRun         *gtk.ToggleButton
	sbDevice      *gtk.SpinButton
	lbModule      *gtk.Label

	fbceProcess *os.Process
)

func initGTK() {
	gtk.Init(&os.Args)
}

func initBuilder() *gtk.Builder {
	var (
		err  error
		b    *gtk.Builder
		data []byte
	)

	b, err = gtk.BuilderNew()
	checkErr(err)
	data, err = Asset(gladeFile)
	checkErr(err)
	err = b.AddFromString(string(data))
	checkErr(err)
	return b
}

func initComponents(b *gtk.Builder) {
	var (
		err error
		o   glib.IObject
		ok  bool
	)

	o, err = b.GetObject("wnMain")
	checkErr(err)
	wnMain, ok = o.(*gtk.Window)
	checkTypecast(ok)

	o, err = b.GetObject("btChooseFile")
	checkErr(err)
	btChooseFile, ok = o.(*gtk.FileChooserButton)
	checkTypecast(ok)

	o, err = b.GetObject("cbLoopForever")
	checkErr(err)
	cbLoopForever, ok = o.(*gtk.CheckButton)
	checkTypecast(ok)

	o, err = b.GetObject("sbLoopCount")
	checkErr(err)
	sbLoopCount, ok = o.(*gtk.SpinButton)
	checkTypecast(ok)

	o, err = b.GetObject("btRun")
	checkErr(err)
	btRun, ok = o.(*gtk.ToggleButton)
	checkTypecast(ok)

	o, err = b.GetObject("sbDevice")
	checkErr(err)
	sbDevice, ok = o.(*gtk.SpinButton)
	checkTypecast(ok)

	o, err = b.GetObject("lbModule")
	checkErr(err)
	lbModule, ok = o.(*gtk.Label)
	checkTypecast(ok)
}

func prepareComponents() {
	wnMain.Connect("destroy", gtk.MainQuit)
	btChooseFile.Connect("file-set", btChooseFileOnFileSet)
	cbLoopForever.Connect("toggled", cbLoopForeverOnToggle)
	btRun.Connect("toggled", btRunOnToggle)

	v4l2loaded, err := fcbe.IsV4L2Loaded()

	if err != nil || !v4l2loaded {
		wnMain.SetSensitive(false)
	}

	if err != nil {
		lbModule.SetText(err.Error())
	} else if !v4l2loaded {
		lbModule.SetText("'v4l2loopback' kernel module is not loaded")
	}
}

func run() {
	wnMain.ShowAll()
	gtk.Main()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkTypecast(b bool) {
	if !b {
		panic(invalidTypecastErr)
	}
}

func main() {
	initGTK()
	initComponents(initBuilder())
	prepareComponents()
	run()
}

func btChooseFileOnFileSet() {
	if btChooseFile.FileChooser.GetFilename() != "" {
		btRun.SetSensitive(true)
	}
}

func cbLoopForeverOnToggle() {
	sbLoopCount.SetSensitive(!cbLoopForever.GetActive())
}

func waitProcessFinish() {
	fbceProcess.Wait()
	btRun.SetActive(false)
}

func btRunOnToggle() {
	var (
		err error
	)

	if btRun.GetActive() {
		loop := sbLoopCount.GetValueAsInt() - 1
		if cbLoopForever.GetActive() {
			loop = -1
		}

		fbceProcess, err = fcbe.PlayFile(btChooseFile.FileChooser.GetFilename(), loop, sbDevice.GetValueAsInt())
		if err != nil {
			panic(err)
		}
		go waitProcessFinish()
	} else {
		fbceProcess.Kill()
	}
}
