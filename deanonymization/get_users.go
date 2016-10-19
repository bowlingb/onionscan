package deanonymization

import (
	"github.com/s-rah/onionscan/config"
	"github.com/s-rah/onionscan/report"
	"regexp"
	"fmt"
	"strings"
)

func GetUsers(osreport *report.OnionScanReport, anonreport *report.AnonymityReport, osc *config.OnionScanConfig) {

	config, ok := osc.CrawlConfigs[osreport.HiddenService]
	if ok {
		for uri, id := range osreport.Crawls {
			crawlRecord, _ := osc.Database.GetCrawlRecord(id)
			if strings.Contains(crawlRecord.Page.Headers.Get("Content-Type"), "text/html") {

				r := regexp.MustCompile(config.UserPage.TriggerRegex)
				result := r.FindAllStringSubmatch(uri, 1)
				if len(result) == 1 {
					osc.LogInfo(fmt.Sprintf("Found Users Page: %s", result[0][1]))
					r = regexp.MustCompile(config.UserPage.NameRegex)
					result = r.FindAllStringSubmatch(crawlRecord.Page.Snapshot, 1)
					if len(result) == 1 {
						osc.LogInfo(fmt.Sprintf("Found User Name: %s", result[0][1]))
					}

					r = regexp.MustCompile(config.UserPage.PositionRegex)
					result = r.FindAllStringSubmatch(crawlRecord.Page.Snapshot, 1)
					if len(result) == 1 {
						osc.LogInfo(fmt.Sprintf("Found User Position: %s", result[0][1]))
					}
				}
			}
		}
	}
}
