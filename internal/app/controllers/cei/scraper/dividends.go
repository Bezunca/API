package scraper

import (
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"github.com/antchfx/htmlquery"
	"log"
	"net/url"
	"time"
)

const dividendsUrl = "ConsultarProventos.aspx"

func getAccountDividends(agent, account string, payloadList []map[string]string) map[string][]models.Dividend {

	log.Printf("------ getAccountDividends( %s , %s )", agent, account)
	log.Printf("\t(Post): %s", ceiBaseUrl+dividendsUrl)

	payload := url.Values{
		"ctl00$ContentPlaceHolder1$ddlAgentes":                        {agent},
		"ctl00$ContentPlaceHolder1$ddlContas":                         {account},
		"ctl00$ContentPlaceHolder1$btnConsultar":                      {"Consultar"},
		"ctl00$ContentPlaceHolder1$ToolkitScriptManager1":             {"ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$btnConsultar"},
		"ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField": {""},
		"ctl00$ContentPlaceHolder1$hdnPDF_EXCEL":                      {""},
		"__EVENTTARGET":                                               {""},
		"__EVENTARGUMENT":                                             {""},
		"__LASTFOCUS":                                                 {""},
		"__ASYNCPOST":                                                 {"false"},
	}

	for _, payloadItem := range payloadList {
		payload.Set(payloadItem["form_key"], payloadItem["form_value"])
	}

	page := utils.PostPage(ceiBaseUrl+dividendsUrl, payload)

	var userCreditedDividends []models.Dividend
	var userProvisionedDividends []models.Dividend

	dividendsTables := htmlquery.Find(page, "//table[@class='responsive']")
	for _, table := range dividendsTables {

		tableTbody := htmlquery.FindOne(table, "//tbody")

		if tableTbody != nil {
			dividends := htmlquery.Find(tableTbody, "//tr")

			for _, dividend := range dividends {

				dInfos := htmlquery.Find(dividend, "//td")

				parsedDividend := models.Dividend{
					utils.CleanString(htmlquery.InnerText(dInfos[0])) + " " + utils.CleanString(htmlquery.InnerText(dInfos[1])),
					utils.CleanString(htmlquery.InnerText(dInfos[2])),
					utils.CleanString(htmlquery.InnerText(dInfos[3])),
					utils.CleanString(htmlquery.InnerText(dInfos[4])),
					utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(dInfos[5]))),
					utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(dInfos[6]))),
					utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(dInfos[7]))),
					utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(dInfos[8]))),
				}

				provisionDate, err := time.Parse("01/02/2006", utils.DateBrToUs(parsedDividend.Date))
				utils.Check(err)
				currentDate := time.Now()

				if parsedDividend.Date != "01/01/0001" && provisionDate.Before(currentDate) {
					userCreditedDividends = append(userCreditedDividends, parsedDividend)
				} else {
					userProvisionedDividends = append(userProvisionedDividends, parsedDividend)
				}
			}
		}
	}

	return map[string][]models.Dividend{
		"credited":    userCreditedDividends,
		"provisioned": userProvisionedDividends,
	}
}

func GetUserDividends(cpf, password string) map[string][]models.Dividend {
	if login(cpf, password) {

		var scrapList []map[string]string
		agents, agentPayload := getAgents(dividendsUrl, scrapList)

		scrapList = []map[string]string{
			{
				"html_path": "//span[@id='ctl00_ContentPlaceHolder1_lblPeriodoFinal']",
				"html_attr": "inner_text",
				"form_key":  "ctl00$ContentPlaceHolder1$txtData",
			},
		}
		accounts, accountPayload := getAgentAccounts(dividendsUrl, agents[0], agentPayload, scrapList)

		return getAccountDividends(agents[0], accounts[0], accountPayload)
	} else {
		return map[string][]models.Dividend{}
	}
}
