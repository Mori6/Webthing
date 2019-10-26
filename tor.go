package tor

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"os"
	"os/exec"
	"strings"
	
)

func check_tor_running() bool { 
	output, err := exec.Command("netstat", "-a", "|", "grep", "9050").Output()
	if err != nil { 
		log.Fatal(err)
		return false
	}

	if (strings.Contains(string(output[:]), "9050")) { 
		fmt.Println("tor is already running...")
	}else { 
		fmt.Println("tor is not running...starting it for you..")
		_, err := exec.Command("sudo", "service", "tor", "start").Output()
		if err != nil { 
			fmt.Println("tor not found...install it first")
			return false
		}
	}

	return true

}

func make_TOR_request(urlstr string) string{

	var webUrl string = urlstr
	torProxy := "socks5://127.0.0.1:9050"
	tor_proxyUrl, err := url.Parse(torProxy)
	if err != nil { 
		fmt.Println("could'nt parse tor url")
	}

	torTransport := &http.Transport{
		Proxy: http.ProxyURL(tor_proxyUrl)}
	
	client := &http.Client{
		Transport: torTransport,
		Timeout:time.Second * 5}


	resp, err := client.Get(webUrl)
	if err != nil { 
		fmt.Println("error, could not make request")
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { 
		fmt.Println("error reading body response")
		os.Exit(1)}

	return string(body)
	

}

func main_tor() { 

	check_tor_running()
	var response string = make_TOR_request("https://www.example.com")
	fmt.Println(response)
}