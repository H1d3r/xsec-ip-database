/*

Copyright (c) 2017 xsec.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package feeds

import (
	"xsec-evil-ips/models"
	"xsec-evil-ips/util"

	"io/ioutil"
	"strings"
	"regexp"
)

func FetchFromEmergingthreats() (evilIps models.EvilIps, err error) {
	url := "http://rules.emergingthreats.net/open/suricata/rules/botcc.rules"
	src := "emergingthreats.net"
	desc := "potential malware site"
	check := "alert ip"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock
	reg := regexp.MustCompile(regexPattern)

	resp, err := util.GetPage(url)
	if err == nil {
		ret, err := ioutil.ReadAll(resp)
		if err == nil {
			lines := strings.Split(string(ret), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "#") || !strings.Contains(line, check) {
					continue
				}
				ret := reg.FindAllString(line, -1)
				evilIps.Ips = append(evilIps.Ips, ret...)
			}
		}
	}
	return evilIps, err
}

func FetchIpsFromEmergingthreats() (evilIps models.EvilIps, err error) {
	url := "http://rules.emergingthreats.net/open/suricata/rules/compromised-ips.txt"
	src := "emergingthreats.net"
	desc := "compromised ips"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	resp, err := util.GetPage(url)
	if err == nil {
		ret, err := ioutil.ReadAll(resp)
		if err == nil {
			ips := strings.Split(string(ret), "\n")
			evilIps.Ips = append(evilIps.Ips, ips...)
		}
	}
	return evilIps, err
}
