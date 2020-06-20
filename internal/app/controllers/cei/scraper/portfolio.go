package scraper

import (
	"bezuncapi/internal/utils"
	"log"
	"net/url"

	"github.com/antchfx/htmlquery"
	"github.com/shopspring/decimal"
)

const portfolioUrl = "ConsultarCarteiraAtivos.aspx"

func getAccountPortfolio(agent, account string, payloadList []map[string]string) []Asset {

	log.Printf("------ getAccountPortfolio( %s , %s )", agent, account)
	log.Printf("\t(Post): %s", ceiBaseUrl+portfolioUrl)

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

	page := utils.PostPage(ceiBaseUrl+portfolioUrl, payload)

	var userPortfolio []Asset

	portfolioTables := htmlquery.Find(page, "//table[@class='Responsive']")
	for _, table := range portfolioTables {

		tableTbody := htmlquery.FindOne(table, "//tbody")

		if tableTbody != nil {

			assets := htmlquery.Find(tableTbody, "//tr")

			for _, asset := range assets {

				aInfos := htmlquery.Find(asset, "//td")

				averagePrice, _ := decimal.NewFromString("0")
				parsedAsset := Asset{
					utils.CleanString(htmlquery.InnerText(aInfos[2])),
					utils.CleanString(htmlquery.InnerText(aInfos[0])) + " " + utils.CleanString(htmlquery.InnerText(aInfos[1])),
					"",
					utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(aInfos[5]))),
					averagePrice,
				}

				userPortfolio = append(userPortfolio, parsedAsset)
			}
		}
	}

	return userPortfolio
}

func GetUserPortfolio(cpf, password string) []Asset {
	if login(cpf, password) {
		var scrapList []map[string]string
		agents, agentPayload := getAgents(portfolioUrl, scrapList)

		scrapList = []map[string]string{
			{
				"html_path": "//span[@id='ctl00_ContentPlaceHolder1_lblPeriodoFinal']",
				"html_attr": "inner_text",
				"form_key":  "ctl00$ContentPlaceHolder1$txtData",
			},
		}
		accounts, accountPayload := getAgentAccounts(portfolioUrl, agents[0], agentPayload, scrapList)

		return getAccountPortfolio(agents[0], accounts[0], accountPayload)
	} else {
		return []Asset{}
	}
}
