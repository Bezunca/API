import requests
from bs4 import BeautifulSoup
import re

CPF = "CPF"
SENHA = "SENHA"
CEI_BASE_URL = 'https://cei.b3.com.br'

#FIX DATE RANGE
DATE_BEGIN = '05/11/2018'
DATE_END = '24/04/2020'


def clean_string(s):
    if s:
        return re.sub(' +', ' ', s.strip().replace("\n",""))
    else:
        return ""

class CEI_Crawler():

    def __init__(self):
        self.login()
        self.transactions = {}

    def login(self):

        self.session = requests.Session()

        page = self.session.get(CEI_BASE_URL, verify=False)
        soup = BeautifulSoup(page.content, 'lxml')

        form = soup.find('form')

        view_state = form.select_one("input[id='__VIEWSTATE']").get('value')
        view_state_generator = form.select_one("input[id='__VIEWSTATEGENERATOR']").get('value')
        event_validation = form.select_one("input[id='__EVENTVALIDATION']").get('value')

        data = {
            'ctl00$ContentPlaceHolder1$txtLogin': CPF,
            'ctl00$ContentPlaceHolder1$txtSenha': SENHA,
            '__VIEWSTATE': view_state,
            '__VIEWSTATEGENERATOR': view_state_generator,
            '__EVENTVALIDATION': event_validation,
            '__EVENTTARGET': '',
            '__EVENTARGUMENT': '',
            '__ASYNCPOST': 'false',
            'ctl00$ContentPlaceHolder1$btnLogar': "Entrar",
            'ctl00$ContentPlaceHolder1$smLoad': 'ctl00$ContentPlaceHolder1$UpdatePanel1|ctl00$ContentPlaceHolder1$btnLogar',
        }
    
        self.session.post(CEI_BASE_URL + "/CEI_Responsivo/", data=data, verify=False)

        print("------ Login Done")


    def craw_transactions(self):
        self.get_agents()

    def get_agents(self):

        url = "/CEI_Responsivo/negociacao-de-ativos.aspx"

        page = self.session.get(CEI_BASE_URL + url, verify=False)
        
        soup = BeautifulSoup(page.content, 'lxml')

        form = soup.find('form')

        view_state = form.select_one("input[id='__VIEWSTATE']").get('value')
        view_state_generator = form.select_one("input[id='__VIEWSTATEGENERATOR']").get('value')
        event_validation = form.select_one("input[id='__EVENTVALIDATION']").get('value')

        data = {
            'ctl00$ContentPlaceHolder1$ddlAgentes': None,
            'ctl00$ContentPlaceHolder1$ddlContas': '0',
            'ctl00$ContentPlaceHolder1$txtDataDeBolsa': DATE_BEGIN,
            'ctl00$ContentPlaceHolder1$txtDataAteBolsa': DATE_END,
            '__VIEWSTATE': view_state,
            '__VIEWSTATEGENERATOR': view_state_generator,
            '__EVENTVALIDATION': event_validation, 
            'ctl00$ContentPlaceHolder1$ToolkitScriptManager1': 'ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$ddlAgentes',
            'ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField': '',
            'ctl00$ContentPlaceHolder1$hdnPDF_EXCEL': '',
            '__EVENTTARGET': 'ctl00$ContentPlaceHolder1$ddlAgentes',
            '__EVENTARGUMENT': '',
            '__LASTFOCUS': '',
            '__ASYNCPOST': 'false'
        }

        agents = soup.find("select", {"name": "ctl00$ContentPlaceHolder1$ddlAgentes"}).findAll("option")
        print("------ {} Agents Found".format(len(agents)))
        for agent in agents:

            self.get_agent_accounts(agent, data)
            

    def get_agent_accounts(self, agent, data):

        url = "/CEI_Responsivo/negociacao-de-ativos.aspx"         

        agent_value = agent.get('value') 
            
        if int(agent_value) != -1:

            self.transactions[agent_value] = {}

            data['ctl00$ContentPlaceHolder1$ddlAgentes'] = agent_value

            page = self.session.post(CEI_BASE_URL + url, data=data, verify=False)

            soup = BeautifulSoup(page.content, 'lxml')

            form = soup.find('form')

            view_state = form.select_one("input[id='__VIEWSTATE']").get('value')
            view_state_generator = form.select_one("input[id='__VIEWSTATEGENERATOR']").get('value')
            event_validation = form.select_one("input[id='__EVENTVALIDATION']").get('value')

            data = {
                'ctl00$ContentPlaceHolder1$ToolkitScriptManager1': 'ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$btnConsultar',
                'ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField': '',
                'ctl00$ContentPlaceHolder1$hdnPDF_EXCEL': '',
                'ctl00$ContentPlaceHolder1$ddlAgentes': agent_value,
                'ctl00$ContentPlaceHolder1$ddlContas': None,
                'ctl00$ContentPlaceHolder1$txtDataDeBolsa': DATE_BEGIN,
                'ctl00$ContentPlaceHolder1$txtDataAteBolsa': DATE_END,
                '__EVENTTARGET': '',
                '__EVENTARGUMENT': '',
                '__LASTFOCUS': '',
                '__VIEWSTATE': view_state,
                '__VIEWSTATEGENERATOR': view_state_generator,
                '__EVENTVALIDATION': event_validation, 
                '__ASYNCPOST': 'false',
                'ctl00$ContentPlaceHolder1$btnConsultar': 'Consultar'
            }

            accounts = soup.find("select", {"name": "ctl00$ContentPlaceHolder1$ddlContas"}).findAll("option")
            print("------ Agent {} - {} Accounts Found".format(agent_value, len(accounts)))
            for account in accounts:

                self.get_account_transactions(agent_value, account, data)

    def get_account_transactions(self, agent_value, account, data):

        url = "/CEI_Responsivo/negociacao-de-ativos.aspx"         

        account_value = account.get('value') 

        self.transactions[agent_value][account_value] = []

        data['ctl00$ContentPlaceHolder1$ddlContas'] = account_value

        page = self.session.post(CEI_BASE_URL + url, data=data, verify=False)

        soup = BeautifulSoup(page.content, 'lxml')

        try:
            transactions = soup.find("table", {"class": "responsive"}).find("tbody", recursive=False).findAll("tr", recursive=False)

            keys = ["date", "action", "market", "expiration", "symbol", "description", "amount", "price", "full_price", "quotation_factor"]
            parsed_transactions = []

            for i, transaction in enumerate(transactions):

                parsed_transaction = {}

                t_infos = transaction.findAll("td", recursive=False)
                for j, t_info in enumerate(t_infos):
                    parsed_transaction[keys[j]] = clean_string(t_info.get_text())
                
                parsed_transactions.append(parsed_transaction)

            self.transactions[agent_value][account_value] = parsed_transactions

            print("------ Agent {} - Account {} - {} Transactions Found".format(agent_value, account_value, len(parsed_transactions)))
        except Exception as e:
            print("------ Agent {} - Account {} - No Transactions Found".format(agent_value, account_value))

c = CEI_Crawler()
c.craw_transactions()

print(c.transactions)