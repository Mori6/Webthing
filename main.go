package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"runtime"

	"io/ioutil"
	"strings"
)

//configure pool options
func configure_xmrig(url string, user string, pass string) { 
	fmt.Println("configuring json script...")
	inputfile, err := ioutil.ReadFile("xmrig/src/core/config/Config_default.h")
	if nil != err { 
		fmt.Println("could'nt read file...")
		os.Exit(1)
	}

	fileslines := strings.Split(string(inputfile), "\n")

	for i, line := range fileslines {
		if strings.Contains(line, "url") {
			fileslines[i] = "\t\t\t\"url\":" + url
		}
	}

	for i, line := range fileslines {
		if strings.Contains(line, "user") {
			fileslines[i] = "\t\t\t\"user\"" + user
		}
	}

	for i, line := range fileslines {
		if strings.Contains(line, "pass") {
			fileslines[i] = "\t\t\t\"pass\"" + pass
		}
	}

	
	output := strings.Join(fileslines, "\n")
        err = ioutil.WriteFile("xmrig/src/core/config/Config_default.h", []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
		}

	fmt.Println("done configuring...")
	

}

//compile a new version of xmrig 
func generate_bin() { 
	fmt.Println("building...")
	_, err := exec.Command("mkdir", "xmrig/build").Output()
	_, err2 := exec.Command("cmake",  "xmrig").Output()
	_, err3 := exec.Command("make", "xmrig/build").Output()
	if err != nil || err2 != nil || err3 != nil{
		fmt.Println("make sure that cmake and make are installed...")
		os.Exit(1)
	}

	if _, err := os.Stat("xmrig"); err != nil { 
		fmt.Println("error: could not build...")
		fmt.Println("make sure all of the dependencies are installed: https://github.com/xmrig/xmrig/wiki/Ubuntu-Build")
		os.Exit(1)
	}

	fmt.Println("done building xmrig...")

	

}

func main() { 
	if runtime.GOOS == "windows" { 
		fmt.Println("Please run on GNU/linux system")
	}else {
		if _, err := os.Stat("xmrig"); err != nil { 
			if os.IsExist(err) { 
				fmt.Println("repo already cloned...")
			}else {
				fmt.Println("cloning...")
				_, err := exec.Command("git", "clone", "https://github.com/xmrig/xmrig").Output()
				if err != nil {
					fmt.Println("git not found...please install git first")
					os.Exit(1)
				}
				configure_xmrig("donate.v2.xmrig.com:3333", "mott", "abc123")
				generate_bin()
				fmt.Println("checking tor proxy...")
				main_tor()
			}
		}
	}
	

}