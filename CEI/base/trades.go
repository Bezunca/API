package base

import (
	"../utils"
	"github.com/antchfx/htmlquery"
	"log"
	"net/url"
)

var CeiBaseUrl = "https://cei.b3.com.br/CEI_Responsivo/"

var DateBegin string
var DateEnd string

type Trade struct {
	Date        string
	Action      string
	Market      string
	Expiration  string
	Symbol      string
	Description string
	Amount      int
	Price       float32
	FullPrice   float32
	PriceFactor int
}

var userTrades []Trade

func login(cpf, password string) bool {

	log.Printf("------ login( %s , %s )", cpf, password)

	log.Printf("\t(Get): %s", CeiBaseUrl)
	loginPage, err := utils.GetPage(CeiBaseUrl)
	if err != nil {
		log.Printf("\t(Get): %v", err)
		return false

	} else {

		viewState := htmlquery.FindOne(loginPage, "//input[@id='__VIEWSTATE']")
		viewStateValue := htmlquery.SelectAttr(viewState, "value")

		viewStateGenerator := htmlquery.FindOne(loginPage, "//input[@id='__VIEWSTATEGENERATOR']")
		viewStateGeneratorValue := htmlquery.SelectAttr(viewStateGenerator, "value")

		eventValidation := htmlquery.FindOne(loginPage, "//input[@id='__EVENTVALIDATION']")
		eventValidationValue := htmlquery.SelectAttr(eventValidation, "value")

		loginPayload := url.Values{
			"ctl00$ContentPlaceHolder1$txtLogin": {cpf},
			"ctl00$ContentPlaceHolder1$txtSenha": {password},
			"__VIEWSTATE":                        {viewStateValue},
			"__VIEWSTATEGENERATOR":               {viewStateGeneratorValue},
			"__EVENTVALIDATION":                  {eventValidationValue},
			"ctl00$ContentPlaceHolder1$btnLogar": {"Entrar"},
			"ctl00$ContentPlaceHolder1$smLoad":   {"ctl00$ContentPlaceHolder1$UpdatePanel1|ctl00$ContentPlaceHolder1$btnLogar"},
			"__EVENTTARGET":                      {""},
			"__EVENTARGUMENT":                    {""},
			"__ASYNCPOST":                        {"false"},
		}

		log.Printf("\t(Post): %s", CeiBaseUrl)
		homePage, err := utils.PostPage(CeiBaseUrl, loginPayload)
		if err != nil {
			log.Printf("\t(Post): %v", err)
			return false
		} else {

			homeTitle := htmlquery.FindOne(homePage, "//a[@href='home.aspx']")
			if homeTitle != nil {
				log.Printf("\tLogin: SUCCESS")
				return true
			} else {
				log.Printf("\tLogin: FAILED")
				return false
			}
		}
	}
}

func getAgents() {

	log.Printf("------ getAgents()")

	log.Printf("\t(Get): %s", CeiBaseUrl+"negociacao-de-ativos.aspx")
	page, err := utils.GetPage(CeiBaseUrl + "negociacao-de-ativos.aspx")
	if err != nil {
		log.Printf("\t(Get): %v\n", err)
	} else {

		viewState := htmlquery.FindOne(page, "//input[@id='__VIEWSTATE']")
		viewStateValue := htmlquery.SelectAttr(viewState, "value")

		viewStateGenerator := htmlquery.FindOne(page, "//input[@id='__VIEWSTATEGENERATOR']")
		viewStateGeneratorValue := htmlquery.SelectAttr(viewStateGenerator, "value")

		eventValidation := htmlquery.FindOne(page, "//input[@id='__EVENTVALIDATION']")
		eventValidationValue := htmlquery.SelectAttr(eventValidation, "value")

		DateBeginSpan := htmlquery.FindOne(page, "//span[@id='ctl00_ContentPlaceHolder1_lblPeriodoInicialBolsa']")
		DateBegin = htmlquery.InnerText(DateBeginSpan)

		DateEndSpan := htmlquery.FindOne(page, "//span[@id='ctl00_ContentPlaceHolder1_lblPeriodoFinalBolsa']")
		DateEnd = htmlquery.InnerText(DateEndSpan)

		agentPayload := url.Values{
			"ctl00$ContentPlaceHolder1$ddlContas":       {"0"},
			"ctl00$ContentPlaceHolder1$txtDataDeBolsa":  {DateBegin},
			"ctl00$ContentPlaceHolder1$txtDataAteBolsa": {DateEnd},
			"__VIEWSTATE":          {viewStateValue},
			"__VIEWSTATEGENERATOR": {viewStateGeneratorValue},
			"__EVENTVALIDATION":    {eventValidationValue},
			"__EVENTTARGET":        {"ctl00$ContentPlaceHolder1$ddlAgentes"},
			"ctl00$ContentPlaceHolder1$ToolkitScriptManager1":             {"ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$ddlAgentes"},
			"ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField": {""},
			"ctl00$ContentPlaceHolder1$hdnPDF_EXCEL":                      {""},
			"__EVENTARGUMENT":                                             {""},
			"__LASTFOCUS":                                                 {""},
			"__ASYNCPOST":                                                 {"false"},
		}

		agents_select := htmlquery.FindOne(page, "//select[@name='ctl00$ContentPlaceHolder1$ddlAgentes']")
		agents := htmlquery.Find(agents_select, "//option")
		log.Printf("\t%d Agents Found", len(agents))

		for _, agent := range agents {
			agentValue := htmlquery.SelectAttr(agent, "value")
			if agentValue != "-1" {
				getAgentAccounts(agentValue, agentPayload)
			}
		}
	}
}

func getAgentAccounts(agent string, agentPayload url.Values) {

	log.Printf("------ getAgentAccounts( %s )", agent)

	log.Printf("\t(Post): %s", CeiBaseUrl+"negociacao-de-ativos.aspx")
	agentPayload.Set("ctl00$ContentPlaceHolder1$ddlAgentes", agent)
	page, err := utils.PostPage(CeiBaseUrl+"negociacao-de-ativos.aspx", agentPayload)
	if err != nil {
		log.Printf("\t(Post): %v", err)
	} else {

		viewState := htmlquery.FindOne(page, "//input[@id='__VIEWSTATE']")
		viewStateValue := htmlquery.SelectAttr(viewState, "value")

		viewStateGenerator := htmlquery.FindOne(page, "//input[@id='__VIEWSTATEGENERATOR']")
		viewStateGeneratorValue := htmlquery.SelectAttr(viewStateGenerator, "value")

		eventValidation := htmlquery.FindOne(page, "//input[@id='__EVENTVALIDATION']")
		eventValidationValue := htmlquery.SelectAttr(eventValidation, "value")

		accountPayload := url.Values{
			"ctl00$ContentPlaceHolder1$ddlAgentes":      {agent},
			"ctl00$ContentPlaceHolder1$txtDataDeBolsa":  {DateBegin},
			"ctl00$ContentPlaceHolder1$txtDataAteBolsa": {DateEnd},
			"__VIEWSTATE":                                                 {viewStateValue},
			"__VIEWSTATEGENERATOR":                                        {viewStateGeneratorValue},
			"__EVENTVALIDATION":                                           {eventValidationValue},
			"ctl00$ContentPlaceHolder1$btnConsultar":                      {"Consultar"},
			"ctl00$ContentPlaceHolder1$ToolkitScriptManager1":             {"ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$btnConsultar"},
			"ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField": {""},
			"ctl00$ContentPlaceHolder1$hdnPDF_EXCEL":                      {""},
			"__EVENTTARGET":                                               {""},
			"__EVENTARGUMENT":                                             {""},
			"__LASTFOCUS":                                                 {""},
			"__ASYNCPOST":                                                 {"false"},
		}

		accounts_select := htmlquery.FindOne(page, "//select[@name='ctl00$ContentPlaceHolder1$ddlContas']")
		accounts := htmlquery.Find(accounts_select, "//option")

		log.Printf("\tAgent %s: %d Accounts Found", agent, len(accounts))

		for _, account := range accounts {
			accountValue := htmlquery.SelectAttr(account, "value")
			if accountValue != "-1" {
				getAccountTrades(agent, accountValue, accountPayload)
			}
		}

	}
}

func getAccountTrades(agent, account string, accountPayload url.Values) {

	log.Printf("------ getAccountTrades( %s , %s )", agent, account)

	log.Printf("\t(Post): %s", CeiBaseUrl+"negociacao-de-ativos.aspx")
	accountPayload.Set("ctl00$ContentPlaceHolder1$ddlContas", account)
	page, err := utils.PostPage(CeiBaseUrl+"negociacao-de-ativos.aspx", accountPayload)
	if err != nil {
		log.Printf("\t(Post): %v", err)
	} else {

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
					utils.StringToInt(utils.CleanString(htmlquery.InnerText(tInfos[6]))),
					utils.StringToFloat32(utils.CleanString(htmlquery.InnerText(tInfos[7]))),
					utils.StringToFloat32(utils.CleanString(htmlquery.InnerText(tInfos[8]))),
					utils.StringToInt(utils.CleanString(htmlquery.InnerText(tInfos[9]))),
				}

				userTrades = append(userTrades, parsedTrade)
			}

			log.Printf("\tAgent %s, Account %s:  %d Transactions Found", agent, account, len(trades))
		} else {
			log.Printf("\tAgent %s, Account %s:  %d Transactions Found", agent, account, 0)
		}

	}
}

func GetRawUserTrades(cpf, password string) []Trade {
	if login(cpf, password) {
		getAgents()
		if len(userTrades) > 0 {
			return userTrades
		} else {
			return nil
		}
	} else {
		return nil
	}
}
