package main

import (
	// "github.com/gin-gonic/autotls"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/resty.v1"
)

type AWSIPRanges struct {
	SyncToken    string      `json:"syncToken"`
	CreateDate   string      `json:"createDate"`
	Prefixes     []AWSPrefix `json:"prefixes"`
	Ipv6Prefixes []AWSPrefix `json:"ipv6_prefixes"`
}

type AWSPrefix struct {
	Ipv6Prefix         string `json:"ipv6_prefix,omitempty"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
	IPPrefix           string `json:"ip_prefix,omitempty"`
}

type GoogleIPRanges struct {
	SyncToken    string         `json:"syncToken"`
	CreationTime string         `json:"creationTime"`
	Prefixes     []GooglePrefix `json:"prefixes"`
}

type GooglePrefix struct {
	Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
}

type AzureIPRanges struct {
	ChangeNumber int64   `json:"changeNumber"`
	Cloud        string  `json:"cloud"`
	Values       []Value `json:"values"`
}

type Value struct {
	Name       string     `json:"name"`
	ID         string     `json:"id"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	ChangeNumber    int64    `json:"changeNumber"`
	Region          string   `json:"region"`
	RegionID        int64    `json:"regionId"`
	Platform        string   `json:"platform"`
	SystemService   string   `json:"systemService"`
	AddressPrefixes []string `json:"addressPrefixes"`
	NetworkFeatures []string `json:"networkFeatures"`
}

var (
	awsip     AWSIPRanges
	googleip  GoogleIPRanges
	azureip   AzureIPRanges
	awsURL    = "https://ip-ranges.amazonaws.com/ip-ranges.json"
	googleURL = "https://www.gstatic.com/ipranges/goog.json"
	azureURL  = "https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20220214.json"
)

func main() {
	r := gin.Default()
	client := resty.New()

	r.GET("/aws", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "aws",
			"url":  awsURL,
		})
	})

	r.GET("/aws/ranges/:iptype", func(c *gin.Context) {
		iptype := c.Param("iptype")
		awsv4 := []string{}
		awsv6 := []string{}
		awsall := []string{}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(awsURL)

		if err != nil {
			c.JSON(200, gin.H{
				"error": fmt.Sprintf("unable to connect to AWS - %s", err),
			})
		}

		if err := json.Unmarshal([]byte(resp.String()), &awsip); err != nil {
			c.JSON(200, gin.H{
				"error": fmt.Sprintf("JSON parse error on IP info - %s", err),
			})
		}

		if iptype == "4" {
			for _, iprange := range awsip.Prefixes {
				awsv4 = append(awsv4, iprange.IPPrefix)
			}

			c.String(200, strings.Join(awsv4, "\n"))
		}

		if iptype == "6" {
			for _, iprange6 := range awsip.Ipv6Prefixes {
				awsv6 = append(awsv6, iprange6.Ipv6Prefix)
			}

			c.String(200, strings.Join(awsv6, "\n"))
		}

		if iptype == "all" {
			for _, iprange := range awsip.Prefixes {
				awsall = append(awsall, iprange.IPPrefix)
			}

			for _, iprange6 := range awsip.Ipv6Prefixes {
				awsall = append(awsall, iprange6.Ipv6Prefix)
			}

			c.String(200, strings.Join(awsall, "\n"))
		}

	})

	r.GET("/google", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "google",
			"url":  googleURL,
		})
	})

	r.GET("/google/ranges/:iptype", func(c *gin.Context) {
		iptype := c.Param("iptype")
		googlev4 := []string{}
		googlev6 := []string{}
		googleall := []string{}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(googleURL)

		if err != nil {
			c.JSON(200, gin.H{
				"error": fmt.Sprintf("unable to connect to Google - %s", err),
			})
		}

		if err := json.Unmarshal([]byte(resp.String()), &googleip); err != nil {
			c.JSON(200, gin.H{
				"error": fmt.Sprintf("JSON parse error on IP info - %s", err),
			})
		}

		for _, iprange := range googleip.Prefixes {
			if iptype == "4" || iptype == "all" {
				if len(iprange.Ipv4Prefix) > 0 && len(iprange.Ipv6Prefix) <= 0 {
					googlev4 = append(googlev4, iprange.Ipv4Prefix)
					googleall = append(googleall, iprange.Ipv4Prefix)
				}
			}

			if iptype == "6" || iptype == "all" {
				if len(iprange.Ipv4Prefix) <= 0 && len(iprange.Ipv6Prefix) > 0 {
					googlev6 = append(googlev6, iprange.Ipv6Prefix)
					googleall = append(googleall, iprange.Ipv6Prefix)
				}
			}
		}

		if iptype == "4" {
			c.String(200, strings.Join(googlev4, "\n"))
		}

		if iptype == "6" {
			c.String(200, strings.Join(googlev6, "\n"))
		}

		if iptype == "all" {
			c.String(200, strings.Join(googleall, "\n"))
		}

	})

	r.GET("/azure", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "azure",
			"url":  azureURL,
		})
	})

	r.GET("/azure/ranges/:iptype", func(c *gin.Context) {
		iptype := c.Param("iptype")
		azurev4 := []string{}
		azurev6 := []string{}
		azureall := []string{}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(azureURL)

		if err != nil {
			c.JSON(200, gin.H{
				"error": fmt.Sprintf("unable to connect to Azure - %s", err),
			})
		}

		if err := json.Unmarshal([]byte(resp.String()), &azureip); err != nil {
			c.JSON(200, gin.H{
				"error": fmt.Sprintf("JSON parse error on IP info - %s", err),
			})
		}

		for _, iprange := range azureip.Values {
			for _, prefixes := range iprange.Properties.AddressPrefixes {
				if iptype == "4" || iptype == "all" {
					if IsIPv4(prefixes) {
						azurev4 = append(azurev4, prefixes)
						azureall = append(azureall, prefixes)
					}
				}

				if iptype == "6" || iptype == "all" {
					if IsIPv6(prefixes) {
						azurev6 = append(azurev6, prefixes)
						azureall = append(azureall, prefixes)
					}
				}
			}
		}

		if iptype == "4" {
			c.String(200, strings.Join(azurev4, "\n"))
		}

		if iptype == "6" {
			c.String(200, strings.Join(azurev6, "\n"))
		}

		if iptype == "all" {
			c.String(200, strings.Join(azureall, "\n"))
		}
	})

	r.Run(":8080")
}

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func sliceToString(slice []string) string {
	var str string

	for _, item := range slice {
		str += fmt.Sprintf("%s, ", item)
	}

	return strings.TrimRight(str, ", ")
}

func stringToSlice(str string) []string {
	var slice []string

	list := strings.FieldsFunc(str, func(r rune) bool { return strings.ContainsRune(",;", r) })
	for _, item := range list {
		slice = append(slice, strings.TrimSpace(item))
	}

	return slice
}
