package main

import (
	"fmt"
	"io/ioutil"
	"os"
)



func main()  {
	tempFile := "/sys/class/thermal/thermal_zone0/temp"
	temp, err := ioutil.ReadFile(tempFile)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(-1)
	}
	reading := string(temp)
	full := reading[0:2]
	partial := reading[2:]
	fmt.Printf("%s.%s", full, partial)
}
