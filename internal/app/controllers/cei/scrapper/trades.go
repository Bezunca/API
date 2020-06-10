package scrapper

import (
	"bezuncapi/internal/utils"
	"log"
	"net/url"

	"github.com/antchfx/htmlquery"
)

const tradesUrl = "negociacao-de-ativos.aspx"

func getAccountTrades(agent, account string, payloadList []map[string]string, userTrades *[]Trade) {

	log.Printf("------ getAccountTrades( %s , %s )", agent, account)
	log.Printf("\t(Post): %s", ceiBaseUrl+tradesUrl)

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

	page := utils.PostPage(ceiBaseUrl+tradesUrl, payload)

	tradesTable := htmlquery.FindOne(page, "//table[@class='responsive']")
	if tradesTable != nil {

		tradesTbody := htmlquery.FindOne(tradesTable, "//tbody")
		trades := htmlquery.Find(tradesTbody, "//tr")

		for _, trade := range trades {

			tInfos := htmlquery.Find(trade, "//td")

			parsedTrade := Trade{
				utils.CleanString(htmlquery.InnerText(tInfos[0])),
				utils.CleanString(htmlquery.InnerText(tInfos[1])),
				utils.CleanString(htmlquery.InnerText(tInfos[2])),
				utils.CleanString(htmlquery.InnerText(tInfos[3])),
				utils.CleanString(htmlquery.InnerText(tInfos[4])),
				utils.CleanString(htmlquery.InnerText(tInfos[5])),
				utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(tInfos[6]))),
				utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(tInfos[7]))),
				utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(tInfos[8]))),
				utils.StringToDecimal(utils.CleanString(htmlquery.InnerText(tInfos[9]))),
			}

			*userTrades = append(*userTrades, parsedTrade)
		}

		log.Printf("\tAgent %s, Account %s:  %d Transactions Found", agent, account, len(trades))
	} else {
		log.Printf("\tAgent %s, Account %s:  %d Transactions Found", agent, account, 0)
	}
}

func GetUserTrades(cpf, password string) []Trade {

	var userTrades []Trade

	if login(cpf, password) {

		var scrapList []map[string]string

		agents, agentPayload := getAgents(tradesUrl, scrapList)

		for _, agent := range agents {
			if agent != "-1" {

				scrapList = []map[string]string{
					{
						"html_path": "//span[@id='ctl00_ContentPlaceHolder1_lblPeriodoInicialBolsa']",
						"html_attr": "inner_text",
						"form_key":  "ctl00$ContentPlaceHolder1$txtDataDeBolsa",
					},
					{
						"html_path": "//span[@id='ctl00_ContentPlaceHolder1_lblPeriodoFinalBolsa']",
						"html_attr": "inner_text",
						"form_key":  "ctl00$ContentPlaceHolder1$txtDataAteBolsa",
					},
				}

				accounts, accountPayload := getAgentAccounts(tradesUrl, agent, agentPayload, scrapList)

				for _, account := range accounts {
					if account != "-1" {
						getAccountTrades(agent, account, accountPayload, &userTrades)
					}
				}
			}
		}
	}

	if userTrades == nil {
		return []Trade{}
	}
	return userTrades
}
