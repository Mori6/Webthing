//golang implementation of belkin wemo upnp RCE
package wemo

import (
	"fmt"

)

func is_vulnerable(ip string, port string) bool { 
	var vuln_paths[] string = "/setupservice.xml", "/firmwareupdate.xml", "/remoteaccess.xml", "/setup.xml"
	for _, i := range vuln_paths {	
		url := ar request string = "http://" + "ip" + ":" + port + i
		client := &http.Client{}
		sRequestContent := generateRequestContent(postalCode)
		requestContent := []byte(sRequestContent)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestContent))
		if err != nil {
			return nil,err
		}
	
		req.Header.Add("SOAPAction", url)
		req.Header.Add("Content-Type", "text/xml; charset=utf-8")
		req.Header.Add("Accept", "text/xml")
		resp, err := client.Do(req)
		if err != nil {
			return nil,err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, errors.New("Error Respose " + resp.Status)
		}
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

}
}

func wemo() { 




}