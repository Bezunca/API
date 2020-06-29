package scraper

import (
	"bezuncapi/internal/utils"
	"github.com/antchfx/htmlquery"
	"log"
	"net/url"
)

const ceiBaseUrl = "https://cei.b3.com.br/CEI_Responsivo/"

func login(cpf, password string) bool {

	log.Printf("------ login( %s , %s )", cpf, password)
	log.Printf("\t(Get): %s", ceiBaseUrl)

	loginPage := utils.GetPage(ceiBaseUrl)

	viewStateValue := htmlquery.SelectAttr(htmlquery.FindOne(loginPage, "//input[@id='__VIEWSTATE']"), "value")

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

	log.Printf("\t(Post): %s", ceiBaseUrl)
	homePage := utils.PostPage(ceiBaseUrl, loginPayload)

	homeTitle := htmlquery.FindOne(homePage, "//a[@href='home.aspx']")
	if homeTitle != nil {
		log.Printf("\tLogin: SUCCESS")
		return true
	} else {
		log.Printf("\tLogin: FAILED")
		return false
	}
}

func getAgents(pageUrl string, scrapList []map[string]string) ([]string, []map[string]string) {

	log.Printf("------ getAgents( %s )", pageUrl)

	log.Printf("\t(Get): %s", ceiBaseUrl+pageUrl)
	page := utils.GetPage(ceiBaseUrl + pageUrl)

	scrapList = append(scrapList,
		map[string]string{
			"html_path": "//input[@id='__VIEWSTATE']",
			"html_attr": "value",
			"form_key":  "__VIEWSTATE"},
		map[string]string{
			"html_path": "//input[@id='__VIEWSTATEGENERATOR']",
			"html_attr": "value",
			"form_key":  "__EVENTVALIDATION"},
		map[string]string{
			"html_path": "//input[@id='__EVENTVALIDATION']",
			"html_attr": "value",
			"form_key":  "__EVENTVALIDATION"})

	var sessionData []map[string]string
	for _, scrapItem := range scrapList {

		item := htmlquery.FindOne(page, scrapItem["html_path"])

		itemValue := ""
		if scrapItem["html_attr"] == "inner_text" {
			itemValue = htmlquery.InnerText(item)
		} else {
			itemValue = htmlquery.SelectAttr(item, scrapItem["html_attr"])
		}

		sessionData = append(sessionData, map[string]string{
			"form_key":   scrapItem["form_key"],
			"form_value": itemValue})
	}

	agentsSelect := htmlquery.FindOne(page, "//select[@name='ctl00$ContentPlaceHolder1$ddlAgentes']")
	agentsOptions := htmlquery.Find(agentsSelect, "//option")

	var agents []string
	for _, agent := range agentsOptions {
		agentValue := htmlquery.SelectAttr(agent, "value")
		agents = append(agents, agentValue)
	}

	log.Printf("\t%d Agents Found", len(agents))
	return agents, sessionData
}

func getAgentAccounts(pageUrl string, agent string, payloadList, scrapList []map[string]string) ([]string, []map[string]string) {

	log.Printf("------ getAgentAccounts( %s )", agent)
	log.Printf("\t(Post): %s", ceiBaseUrl+pageUrl)

	payload := url.Values{
		"ctl00$ContentPlaceHolder1$ddlAgentes":                        {agent},
		"ctl00$ContentPlaceHolder1$ddlContas":                         {"0"},
		"__EVENTTARGET":                                               {"ctl00$ContentPlaceHolder1$ddlAgentes"},
		"ctl00$ContentPlaceHolder1$ToolkitScriptManager1":             {"ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$ddlAgentes"},
		"ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField": {""},
		"ctl00$ContentPlaceHolder1$hdnPDF_EXCEL":                      {""},
		"__EVENTARGUMENT":                                             {""},
		"__LASTFOCUS":                                                 {""},
		"__ASYNCPOST":                                                 {"false"},
	}

	for _, payloadItem := range payloadList {
		payload.Set(payloadItem["form_key"], payloadItem["form_value"])
	}

	page := utils.PostPage(ceiBaseUrl+pageUrl, payload)

	scrapList = append(scrapList,
		map[string]string{
			"html_path": "//input[@id='__VIEWSTATE']",
			"html_attr": "value",
			"form_key":  "__VIEWSTATE"},
		map[string]string{
			"html_path": "//input[@id='__VIEWSTATEGENERATOR']",
			"html_attr": "value",
			"form_key":  "__EVENTVALIDATION"},
		map[string]string{
			"html_path": "//input[@id='__EVENTVALIDATION']",
			"html_attr": "value",
			"form_key":  "__EVENTVALIDATION"})

	var sessionData []map[string]string
	for _, scrapItem := range scrapList {
		item := htmlquery.FindOne(page, scrapItem["html_path"])

		itemValue := ""
		if scrapItem["html_attr"] == "inner_text" {
			itemValue = htmlquery.InnerText(item)
		} else {
			itemValue = htmlquery.SelectAttr(item, scrapItem["html_attr"])
		}

		sessionData = append(sessionData, map[string]string{"form_key": scrapItem["form_key"], "form_value": itemValue})
	}

	accountsSelect := htmlquery.FindOne(page, "//select[@name='ctl00$ContentPlaceHolder1$ddlContas']")
	accountsOption := htmlquery.Find(accountsSelect, "//option")

	var accounts []string
	for _, account := range accountsOption {
		accountValue := htmlquery.SelectAttr(account, "value")
		accounts = append(accounts, accountValue)
	}

	log.Printf("\tAgent %s: %d Accounts Found", agent, len(accounts))
	return accounts, sessionData
}
