package main

import (
	"fmt"
	"github.com/codegangsta/cli" //library which manages the console
	"gopkg.in/fsnotify.v1"       //library neccessary for instant compilation
	"io/ioutil"
	"log"
	"os"
	//package for the repl dependen theings
	"github.com/tonnerre/golang-go.crypto/ssh/terminal"
)

func main() {
	log.SetFlags(0)
	NewApp().Run(os.Args)
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "run"
	app.Usage = "run programming language tools"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "info",
			Usage: "print some basic information",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "repl",
			Usage: "start a interactive repl",
			Action: func(c *cli.Context) {
				fmt.Println("diese funktion ist noch nicht implementiert und wird das Programm aufhängen")
				oldState, err := terminal.MakeRaw(0)
				if err != nil {
					panic(err)
				}
				b := make([]byte, 0)
				for {
					os.Stdin.Read(b)
					os.Stdout.Write(b)
				}
				defer terminal.Restore(0, oldState)
			},
		},
		{
			Name:  "run",
			Usage: "execute a given file",
			Action: func(c *cli.Context) {
				path := c.Args().Get(0)
				file := c.Args().Get(1)
				fmt.Println("add file", file, "to", path)
			},
		},
		{
			Name:  "build",
			Usage: "compile the given soruce file",
			Action: func(c *cli.Context) {
				format := c.Args().Get(0)
				file := c.Args().Get(1)
				fmt.Println("compile ", file, "to", format)
			},
		},
		{
			Name:  "instant",
			Usage: "this function looks at the source and updates the view on every save",
			Action: func(c *cli.Context) {
				watcher, err := fsnotify.NewWatcher()
				if err != nil {
					log.Fatal(err)
				}
				defer watcher.Close()

				done := make(chan bool)
				go func() {
					for {
						select {
						case event := <-watcher.Events:
							log.Println("event:", event)
							if event.Op&fsnotify.Write == fsnotify.Write {
								log.Println("modified file:", event.Name)
							}
						case err := <-watcher.Errors:
							log.Println("error:", err)
						}
					}
				}()
				log.Println("instant watcher started")
				err = watcher.Add("./")
				if err != nil {
					log.Fatal(err)
				}
				<-done
			},
		},
		{
			Name:  "fmt",
			Usage: "Give it a source file and it will style it properly",
			Action: func(c *cli.Context) {
				file := c.Args().Get(0)
				fmt.Println("watch ", file, "and dependencys and run instantly")
			},
		},
		{
			Name:  "example",
			Usage: "Compile the example",
			Action: func(c *cli.Context) {
				data, err := ioutil.ReadFile("example.run")
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(data)
			},
		},
	}
	return app
}
