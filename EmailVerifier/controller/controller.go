package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type Domain struct {
	Domain string `json:"domain"`
}

func (c *Domain) isEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.Domain == ""
}

func CheckDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Email Field is Empty")
	}

	var domain Domain

	_ = json.NewDecoder(r.Body).Decode(&domain)

	if domain.isEmpty() {
		json.NewEncoder(w).Encode("JSON: Please send some data")
	}

	var hasMX, hasSPF, hasDMRC bool
	var spfRecord, dmarcRecord string

	hasMX = MX(domain)
	hasSPF, spfRecord = SPF(domain)
	hasDMRC, dmarcRecord = DMRC(domain)

	if hasMX && hasSPF && hasDMRC {
		json.NewEncoder(w).Encode("domain find was successfull!")
	}

	fmt.Printf("%v %v %v %v %v %v", domain, hasMX, hasSPF, hasDMRC, spfRecord, dmarcRecord)
}

func MX(domain Domain) bool {
	var hasMX bool
	mxRecords, err := net.LookupMX(domain.Domain)

	if err != nil {
		log.Fatal("1", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	return hasMX
}

func SPF(domain Domain) (bool, string) {
	var hasSPF bool
	var spfRecord string
	txtRecords, err := net.LookupTXT(domain.Domain)

	if err != nil {
		log.Fatal("2", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	return hasSPF, spfRecord
}

func DMRC(domain Domain) (bool, string) {
	var hasDMRC bool
	var dmarcRecord string

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain.Domain)

	if err != nil {
		log.Fatal("3", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMRC = true
			dmarcRecord = record
			break
		}
	}

	return hasDMRC, dmarcRecord

}
