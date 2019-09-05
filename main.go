package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type record struct {
	ip    string
	names []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseEntry(entry string) (string, []string) {
	s := strings.Split(entry, " ")
	ip := s[0]
	domain_list := s[1]
	dommain_array := strings.Split(domain_list, ",")
	return ip, dommain_array
}

func main() {
	var dnsFilename string
	var urlFilename string

	flag.StringVar(&dnsFilename, "d", "", "file containing entries like '1.1.1.1 example.com,www.example.com'")
	flag.StringVar(&urlFilename, "u", "", "file containing a list of URLs")

	flag.Parse()

	if dnsFilename == "" || urlFilename == "" {
		flag.PrintDefaults()
		return
	}

	dnsRecords, err := os.Open(dnsFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer dnsRecords.Close()

	file, err := os.Open(urlFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records := make(map[string][]string)

	var ip string
	var domains []string
	sc := bufio.NewScanner(dnsRecords)

	for sc.Scan() {
		ip, domains = ParseEntry(sc.Text())
		records[ip] = domains
	}

	sc = bufio.NewScanner(file)
	for sc.Scan() {
		l := sc.Text()
		// grab IP addresses from each line
		r, _ := regexp.Compile("(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)")
		ip := r.FindString(l)
		domains = records[ip]
		for _, domain := range domains {
			// massdns likes to put a dot at the end of a domain name
			// so let's get rid of it
			if string(domain[len(domain)-1]) == "." {
				domain = domain[:len(domain)-1]
			}
			fmt.Println(strings.ReplaceAll(l, ip, domain))
		}
	}
}
