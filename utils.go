package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"time"

	"github.com/dlintw/goconf"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var props *goconf.ConfigFile

func loadProps() *goconf.ConfigFile {

	fmt.Println("loading properties")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	check(err)
	prop, err := goconf.ReadConfigFile(filepath.Join(dir, "props.cfg"))
	check(err)

	fmt.Println("done loading props")
	return prop
}

func getTime() string {
	return fmt.Sprintf("%v:%02d:%02d",
		time.Now().Hour(), time.Now().Minute(),
		time.Now().Second())
}

// GetNumOfProcesses - gets the properties from the given server
func GetNumOfProcesses() (int, error) {
	processes, err := exec.Command("sh", "-c", "ps aux | grep aprocess | grep -v grep | wc -l").Output()
	p, _ := strconv.Atoi(fmt.Sprintf("%v", processes))

	return p, err
}

// KillAllProcesses - Kills the schedulers
func KillAllProcesses() error {
	_, err := exec.Command("sh", "-c", "ps -ef | grep aprocess | grep -v grep | awk '{print $2}' | xargs kill -9").Output()
	return err
}

// ResetFiles - Kills the schedulers
func ResetFiles() error {
	_, err := exec.Command("sh", "-c", "./reset.sh").Output()
	return err
}
