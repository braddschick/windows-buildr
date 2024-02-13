package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"windows-buildr/types"
)

func main() {
	name := flag.String("n", "", "Name of the Software")
	vers := flag.String("v", "", "Version of the Software v0.0.0")
	icon := flag.String("i", "", "Path to ICON PNG with -w prepended ../icon.png")
	desc := flag.String("d", "", "Description of the Software")
	workDir := flag.String("w", "", "Working Directory")
	oFile := flag.String("o", "", "Filename for the output resource file with -w prepended defaults is rsrc_windows_amd64.syso")
	company := flag.String("c", "", "Company Name of the Software")
	copy := flag.String("r", "", "Copyright of the Software")
	flag.Parse()

	if *name == "" || *vers == "" || *icon == "" || *desc == "" || *company == "" || *copy == "" || *workDir == "" {
		log.Fatalln("All flags are required")
		flag.PrintDefaults()
	}
	if *oFile == "" {
		*oFile = "rsrc_windows_amd64.syso"
	}

	if _, err := os.Stat(*workDir); os.IsNotExist(err) {
		log.Fatalln("Working Directory does not exist")
		fmt.Println("")
		fmt.Println(*workDir + " does not exist")
	}

	wi := types.WindowsInfo{
		Name:        *name,
		Version:     *vers,
		IconPath:    path.Join(*workDir, *icon),
		Description: *desc,
		SysPath:     path.Join(*workDir, *oFile),
		CompanyName: *company,
		CopyRight:   *copy,
	}
	if err := wi.Write(); err != nil {
		log.Fatalln(err)
		fmt.Println("")
		fmt.Println("failed to write... " + err.Error())
	}
}
