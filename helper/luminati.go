package helper

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var ZoneRegister = &Luminati{
	Account:  "",
	Zone:     "",
	Password: "",
}

var ZoneProxy1 = &Luminati{
	Account:  "",
	Zone:     "",
	Password: "",
}

var lumToken = ""

type Luminati struct {
	Account  string
	Zone     string
	Password string
}

func (p *Luminati) GetIps(country string) []string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://brightdata.com/api/zone/route_ips?zone=%s&country=%s", p.Zone, country), nil)
	req.Header.Set("Authorization", "Bearer "+lumToken)
	resp, err := client.Do(req)
	if err == nil {
		ips, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return strings.Split(string(ips), "\n")
	} else {
		return []string{}
	}
}

func (p *Luminati) GetPorxy(args string) string {
	return fmt.Sprintf("http://brd-customer-%s-zone-%s%s:%s@zproxy.lum-superproxy.io:22225", p.Account, p.Zone, args, p.Password)
}

func (p *Luminati) GetPorxyByCountry(country string) string {
	//ips := p.GetIps("id")
	ips := p.GetIps(country)
	rand.Seed(time.Now().Unix())
	if len(ips) < 2 {
		fmt.Println("can't get ip in", country)
		return fmt.Sprintf("http://brd-customer-%s-zone-%s-country-us:%s@zproxy.lum-superproxy.io:22225", p.Account, p.Zone, p.Password)
	}
	ip := ips[rand.Intn(len(ips))]
	return fmt.Sprintf("http://brd-customer-%s-zone-%s-ip-%s:%s@zproxy.lum-superproxy.io:22225", p.Account, p.Zone, ip, p.Password)
}
