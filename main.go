package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/template"

	gcf "github.com/zhaizhonghao/connectionHelper/services/ccf"
)

var tpl *template.Template

func main() {

	numberOfOrgs := flag.Int("numberOfOrgs", 2, "please input the number of organizations")
	flag.Parse()

	fmt.Printf("generate generate-ccp.sh for %d organzations\n", *numberOfOrgs)

	organizations := []gcf.Organization{}

	for i := 1; i <= *numberOfOrgs; i++ {
		organization := gcf.Organization{}
		organization.OrgNum = strconv.Itoa(i)
		organization.OrgName = "org" + strconv.Itoa(i)
		organization.CAPort = strconv.Itoa(7+i-1) + "054"
		organization.P0Port = strconv.Itoa(7+(i-1)*2) + "051"
		organization.P1Port = strconv.Itoa(8+(i-1)*2) + "051"
		organizations = append(organizations, organization)
	}
	configInfo := gcf.ConfigInfo{}
	configInfo.Organizations = organizations

	//Generate generate-ccp.sh
	tpl = template.Must(template.ParseGlob("templates/*.sh"))
	file, err := os.Create("generate-ccp.sh")
	if err != nil {
		fmt.Println("Fail to create file!")
	}
	err = gcf.GenerateCCPTemplate(configInfo, tpl, file)
	if err != nil {
		fmt.Println("Fail to generate generate-ccp.sh", err)
	}
}
