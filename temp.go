package main

import (
	"fmt"
	"os"
)



func main()  {
	tempFile := "/sys/class/thermal/thermal_zone0/temp"
	fd, err := os.OpenFile(tempFile)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(-1)
	}

}
