# CEI Scraper (Reverse Engineered)

Documentação das URL's usadas pelo CEI no seu website.

URL base:

- `https://cei.b3.com.br/CEI_Responsivo/`

## Login

<details>

<summary>Página de Login</summary>

<p>

__Descrição:__ Página utilizada para capturar dados que serão utilizados como input do `Formulário de Login`. Os dados são:
- Variáveis internas do CEI para controle de sessão.

__URL:__ `https://cei.b3.com.br/CEI_Responsivo`

__Método:__ GET

__Tipo do Retorno:__ `text/html`
 
__Scrap (Exemplo):__ 

Capturar o atributo `value` dos elementos abaixo:

```html
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwUKMTc3NDI2MTA1OA9kFgJmD2QWAgIDD2QWAgIBD2QWAgIIDxYCHgdWaXNpYmxlaGRk7P4AjuBFTAmTPK6r/26AJZjS3WI=" />

<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="803C878C" />

<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAASpMZlRQVEkIJsV6kw/uC9KdHQiWNJPAzoojF2W6rb9pVH0ZHo6j+VFVNKHCgI4S+fS7Ixw1xxNt6QIpIQNh+THL5njjjIDpf92JIAZu/Tu1Az5Buc=" />
```

</p>

</details>

<details>

<summary>Formulário de Login</summary>

<p>

__Descrição:__ Irá permitir o acesso, dentro da sessão, nas demais URL's utilizadas.

__URL:__ `https://cei.b3.com.br/CEI_Responsivo`

__Método:__ POST

__Form Data (Exemplo):__

```json
{
    "ctl00$ContentPlaceHolder1$txtLogin": "12345678901",
    "ctl00$ContentPlaceHolder1$txtSenha": "senha123",
    "__VIEWSTATE": "/wEPDwUKMTc3NDI2MTA1OA9kFgJmD2QWAgIDD2QWAgIBD2QWAgIIDxYCHgdWaXNpYmxlaGRk7P4AjuBFTAmTPK6r/26AJZjS3WI=",
    "__VIEWSTATEGENERATOR": "803C878C",
    "__EVENTVALIDATION": "/wEdAASpMZlRQVEkIJsV6kw/uC9KdHQiWNJPAzoojF2W6rb9pVH0ZHo6j+VFVNKHCgI4S+fS7Ixw1xxNt6QIpIQNh+THL5njjjIDpf92JIAZu/Tu1Az5Buc=",
    "__EVENTTARGET": "",
    "__EVENTARGUMENT": "",
    "__ASYNCPOST": "false",
    "ctl00$ContentPlaceHolder1$btnLogar": "Entrar",
    "ctl00$ContentPlaceHolder1$smLoad": "ctl00$ContentPlaceHolder1$UpdatePanel1|ctl00$ContentPlaceHolder1$btnLogar"
}
```

- `ctl00$ContentPlaceHolder1$txtLogin`: `string`
  > CPF do usuário
- `ctl00$ContentPlaceHolder1$txtSenha`: `string`
  > Senha do usuário
- `__VIEWSTATE`: `string`
  > Adquirida na `Página de Login`
- `__VIEWSTATEGENERATOR`: `string`
  > Adquirida na `Página de Login`
- `__EVENTTARGET`: `string`
  > Adquirida na `Página de Login`

__Tipo do Retorno:__ `text/html`

</p>

</details>

## Negociação de Ativos

<details>

<summary>Página de Negociação de Ativos</summary>

<p>

__Descrição:__ Página utilizada para capturar dados que serão utilizados como input da `Escolha da Instituição`. Os dados são:
- Variáveis internas do CEI para controle de sessão.
- Instituições Financeiras do usuário.

__URL:__ `https://cei.b3.com.br/CEI_Responsivo/negociacao-de-ativos.aspx`

__Método:__ GET

__Tipo do Retorno:__ `text/html`
 
__Scrap (Exemplo):__ 

Capturar o atributo `value` dos elementos abaixo:

```html
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwUKLTI2MTA0ODczNg9kFgJmD2QWAgIDD2QWCAIBDw8WAh4EVGV4dAUcRkxBVklPIFJJQkVJUk8gVEVJWEVJUkEgTkVUT2RkAgMPZBYCAgEPFgIeB1Zpc2libGVoZAIFD2QWBAIBDw8WAh8ABRZOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCAw9kFgICAw8PFgIfAAUzIC8gRXh0cmF0b3MgZSBJbmZvcm1hdGl2b3MgLyBOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCBw9kFgICBQ9kFgJmD2QWCAIDDxBkEBUGCVNlbGVjaW9uZRYxOTAgLSBXQVJSRU4gQ1ZNQyBMVERBIDMwOCAtIENMRUFSIENPUlJFVE9SQSAtIEdSVVBPIFhQGjkzIC0gTk9WQSBGVVRVUkEgQ1RWTSBMVERBIzM4NiAtIFJJQ08gSU5WRVNUSU1FTlRPUyAtIEdSVVBPIFhQHzkwIC0gRUFTWU5WRVNUIC0gVElUVUxPIENWIFMuQS4VBgItMQMxOTADMzA4AjkzAzM4NgI5MBQrAwZnZ2dnZ2cWAWZkAgUPEA8WAh4HRW5hYmxlZGhkEBUBDFNlbGVjaW9uZS4uLhUBATAUKwMBZxYBZmQCBw8PFgIfAWdkFgQCBQ8PFgIfAAUKMDUvMTEvMjAxOGRkAgcPDxYCHwAFCjI0LzA0LzIwMjBkZAIJD2QWCAIBDw8WAh8ABQowNS8xMS8yMDE4ZGQCAw8PFgIfAAUKMjQvMDQvMjAyMGRkAgUPDxYCHwAFCjA1LzExLzIwMThkZAIHDw8WAh8ABQoyNC8wNC8yMDIwZGRkbUuDiMMPsTsDi7AF/W7eMjFd83s=" />

<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="B345DEBA" />

<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAAzWqnmr/YC0E8PGCYgOisxPkdMq+p9EbLSDyFDpFsJg32RR0+cq1RgYtJkxz8vozrSFp0V+sAU3qQ1xD0H4ZlandPwdZycbAgmz2rhziqC/I4rKgXGVnMYDWSeXmAgyzFJ+T4cvEJcdEKVEObT+JUFSVx7WZDZAC7VL6Ff5PQQ8MglkmHVWFVVrRRDXkWFySGBqNF2aESdCPXyA43vLkvgLAYyn3p774f+hILGCk2jveex99A477mR6kRNOW1balkiyu79flJiXkRgOzNyRwuCSascvhw==" />
```

Capturar o atributo `value` de todas as instituições financeiras do usuário no `select` abaixo:

```html
<select name="ctl00$ContentPlaceHolder1$ddlAgentes" id="ctl00_ContentPlaceHolder1_ddlAgentes">
    <option selected="selected" value="-1">Selecione</option>
    <option value="190">190 - WARREN CVMC LTDA</option>
    <option value="308">308 - CLEAR CORRETORA - GRUPO XP</option>
    <option value="93">93 - NOVA FUTURA CTVM LTDA</option>
    <option value="386">386 - RICO INVESTIMENTOS - GRUPO XP</option>
    <option value="90">90 - EASYNVEST - TITULO CV S.A.</option>
</select>
```

</p>

</details>

<details>

<summary>Escolha da Instituição</summary>

<p>

__Descrição:__ Página utilizada para capturar dados que serão utilizados como input da `Escolha da Conta`. Os dados são:
- Variáveis internas do CEI para controle de sessão.
- Contas de uma Instituição Financeira do usuário.

__URL:__ `https://cei.b3.com.br/CEI_Responsivo/negociacao-de-ativos.aspx`

__Método:__ POST

__Form Data (Exemplo):__

```json
{
    "ctl00$ContentPlaceHolder1$ddlAgentes": "308",
    "ctl00$ContentPlaceHolder1$txtDataDeBolsa": "05/11/2018",
    "ctl00$ContentPlaceHolder1$txtDataAteBolsa": "24/04/2020",
    "__VIEWSTATE": "/wEPDwUKLTI2MTA0ODczNg9kFgJmD2QWAgIDD2QWCAIBDw8WAh4EVGV4dAUcRkxBVklPIFJJQkVJUk8gVEVJWEVJUkEgTkVUT2RkAgMPZBYCAgEPFgIeB1Zpc2libGVoZAIFD2QWBAIBDw8WAh8ABRZOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCAw9kFgICAw8PFgIfAAUzIC8gRXh0cmF0b3MgZSBJbmZvcm1hdGl2b3MgLyBOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCBw9kFgICBQ9kFgJmD2QWCAIDDxBkEBUGCVNlbGVjaW9uZRYxOTAgLSBXQVJSRU4gQ1ZNQyBMVERBIDMwOCAtIENMRUFSIENPUlJFVE9SQSAtIEdSVVBPIFhQGjkzIC0gTk9WQSBGVVRVUkEgQ1RWTSBMVERBIzM4NiAtIFJJQ08gSU5WRVNUSU1FTlRPUyAtIEdSVVBPIFhQHzkwIC0gRUFTWU5WRVNUIC0gVElUVUxPIENWIFMuQS4VBgItMQMxOTADMzA4AjkzAzM4NgI5MBQrAwZnZ2dnZ2cWAWZkAgUPEA8WAh4HRW5hYmxlZGhkEBUBDFNlbGVjaW9uZS4uLhUBATAUKwMBZxYBZmQCBw8PFgIfAWdkFgQCBQ8PFgIfAAUKMDUvMTEvMjAxOGRkAgcPDxYCHwAFCjI0LzA0LzIwMjBkZAIJD2QWCAIBDw8WAh8ABQowNS8xMS8yMDE4ZGQCAw8PFgIfAAUKMjQvMDQvMjAyMGRkAgUPDxYCHwAFCjA1LzExLzIwMThkZAIHDw8WAh8ABQoyNC8wNC8yMDIwZGRkbUuDiMMPsTsDi7AF/W7eMjFd83s=",
    "__VIEWSTATEGENERATOR": "B345DEBA",
    "__EVENTVALIDATION": "/wEdAAzWqnmr/YC0E8PGCYgOisxPkdMq+p9EbLSDyFDpFsJg32RR0+cq1RgYtJkxz8vozrSFp0V+sAU3qQ1xD0H4ZlandPwdZycbAgmz2rhziqC/I4rKgXGVnMYDWSeXmAgyzFJ+T4cvEJcdEKVEObT+JUFSVx7WZDZAC7VL6Ff5PQQ8MglkmHVWFVVrRRDXkWFySGBqNF2aESdCPXyA43vLkvgLAYyn3p774f+hILGCk2jveex99A477mR6kRNOW1balkiyu79flJiXkRgOzNyRwuCSascvhw==", 
    "ctl00$ContentPlaceHolder1$ddlContas": "0",
    "ctl00$ContentPlaceHolder1$ToolkitScriptManager1": "ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$ddlAgentes",
    "ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField": "",
    "ctl00$ContentPlaceHolder1$hdnPDF_EXCEL": "",
    "__EVENTTARGET": "ctl00$ContentPlaceHolder1$ddlAgentes",
    "__EVENTARGUMENT": "",
    "__LASTFOCUS": "",
    "__ASYNCPOST": "false"
}
```

- `ctl00$ContentPlaceHolder1$ddlAgentes`: `string`
  > `value` da Instituição escolhida. Adquirida na `Página de Negociação de Ativos`
- `ctl00$ContentPlaceHolder1$txtDataDeBolsa`: `string`
  > Data de Início do intervalo da consulta.
- `ctl00$ContentPlaceHolder1$txtDataAteBolsa`: `string`
  > Data de Término do intervalo da consulta.
- `__VIEWSTATE`: `string`
  > Adquirida na `Página de Negociação de Ativos`
- `__VIEWSTATEGENERATOR`: `string`
  > Adquirida na `Página de Negociação de Ativos`
- `__EVENTTARGET`: `string`
  > Adquirida na `Página de Negociação de Ativos`

__Tipo do Retorno:__ `text/html`

__Scrap (Exemplo):__ 

Capturar o atributo `value` dos elementos abaixo:

```html
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwUKLTI2MTA0ODczNg9kFgJmD2QWAgIDD2QWCAIBDw8WAh4EVGV4dAUcRkxBVklPIFJJQkVJUk8gVEVJWEVJUkEgTkVUT2RkAgMPZBYCAgEPFgIeB1Zpc2libGVoZAIFD2QWBAIBDw8WAh8ABRZOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCAw9kFgICAw8PFgIfAAUzIC8gRXh0cmF0b3MgZSBJbmZvcm1hdGl2b3MgLyBOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCBw9kFgICBQ9kFgJmD2QWCAIDDxBkEBUGCVNlbGVjaW9uZRYxOTAgLSBXQVJSRU4gQ1ZNQyBMVERBIDMwOCAtIENMRUFSIENPUlJFVE9SQSAtIEdSVVBPIFhQGjkzIC0gTk9WQSBGVVRVUkEgQ1RWTSBMVERBIzM4NiAtIFJJQ08gSU5WRVNUSU1FTlRPUyAtIEdSVVBPIFhQHzkwIC0gRUFTWU5WRVNUIC0gVElUVUxPIENWIFMuQS4VBgItMQMxOTADMzA4AjkzAzM4NgI5MBQrAwZnZ2dnZ2cWAQIBZAIFDxAPFgIeB0VuYWJsZWRoZBAVAQcxMDQ5MTc1FQEHMTA0OTE3NRQrAwFnFgFmZAIHDw8WAh8BZ2QWBAIFDw8WAh8ABQowNS8xMS8yMDE4ZGQCBw8PFgIfAAUKMjQvMDQvMjAyMGRkAgkPZBYIAgEPDxYCHwAFCjA1LzExLzIwMThkZAIDDw8WAh8ABQoyNC8wNC8yMDIwZGQCBQ8PFgIfAAUKMDUvMTEvMjAxOGRkAgcPDxYCHwAFCjI0LzA0LzIwMjBkZGTUtYXZXOr8LIUehtEoKws476RhAw==" />

<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="B345DEBA" />

<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAAzyyzbtePOqzC88tnuVBCp3kdMq+p9EbLSDyFDpFsJg32RR0+cq1RgYtJkxz8vozrSFp0V+sAU3qQ1xD0H4ZlandPwdZycbAgmz2rhziqC/I4rKgXGVnMYDWSeXmAgyzFJ+T4cvEJcdEKVEObT+JUFSVx7WZDZAC7VL6Ff5PQQ8Ms95LlC7dkfH1zfWJcIo92xqNF2aESdCPXyA43vLkvgLAYyn3p774f+hILGCk2jveex99A477mR6kRNOW1balkgbPMuems7B81VXuzoM2X8ty/G7VA==" />
```

Capturar o atributo `value` de todas as contas da instituição financeira escolhida no `select` abaixo:

```html
<select name="ctl00$ContentPlaceHolder1$ddlContas" id="ctl00_ContentPlaceHolder1_ddlContas">
    <option value="1178885">1178885</option>
</select>
```

</p>

</details>

<details>

<summary>Escolha da Conta</summary>

<p>

__Descrição:__ Página utilizada para capturar as `Negociações de Ativos`.

__URL:__ `https://cei.b3.com.br/CEI_Responsivo/negociacao-de-ativos.aspx`

__Método:__ POST

__Form Data (Exemplo):__

```json
{
    "ctl00$ContentPlaceHolder1$ddlAgentes": "308",
    "ctl00$ContentPlaceHolder1$ddlContas": "1178885",
    "ctl00$ContentPlaceHolder1$txtDataDeBolsa": "05/11/2018",
    "ctl00$ContentPlaceHolder1$txtDataAteBolsa": "24/04/2020",
    "__VIEWSTATE": "/wEPDwUKLTI2MTA0ODczNg9kFgJmD2QWAgIDD2QWCAIBDw8WAh4EVGV4dAUcRkxBVklPIFJJQkVJUk8gVEVJWEVJUkEgTkVUT2RkAgMPZBYCAgEPFgIeB1Zpc2libGVoZAIFD2QWBAIBDw8WAh8ABRZOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCAw9kFgICAw8PFgIfAAUzIC8gRXh0cmF0b3MgZSBJbmZvcm1hdGl2b3MgLyBOZWdvY2lhw6fDo28gZGUgYXRpdm9zZGQCBw9kFgICBQ9kFgJmD2QWCAIDDxBkEBUGCVNlbGVjaW9uZRYxOTAgLSBXQVJSRU4gQ1ZNQyBMVERBIDMwOCAtIENMRUFSIENPUlJFVE9SQSAtIEdSVVBPIFhQGjkzIC0gTk9WQSBGVVRVUkEgQ1RWTSBMVERBIzM4NiAtIFJJQ08gSU5WRVNUSU1FTlRPUyAtIEdSVVBPIFhQHzkwIC0gRUFTWU5WRVNUIC0gVElUVUxPIENWIFMuQS4VBgItMQMxOTADMzA4AjkzAzM4NgI5MBQrAwZnZ2dnZ2cWAQIBZAIFDxAPFgIeB0VuYWJsZWRoZBAVAQcxMDQ5MTc1FQEHMTA0OTE3NRQrAwFnFgFmZAIHDw8WAh8BZ2QWBAIFDw8WAh8ABQowNS8xMS8yMDE4ZGQCBw8PFgIfAAUKMjQvMDQvMjAyMGRkAgkPZBYIAgEPDxYCHwAFCjA1LzExLzIwMThkZAIDDw8WAh8ABQoyNC8wNC8yMDIwZGQCBQ8PFgIfAAUKMDUvMTEvMjAxOGRkAgcPDxYCHwAFCjI0LzA0LzIwMjBkZGTUtYXZXOr8LIUehtEoKws476RhAw==",
    "__VIEWSTATEGENERATOR": "B345DEBA",
    "__EVENTVALIDATION": "/wEdAAzyyzbtePOqzC88tnuVBCp3kdMq+p9EbLSDyFDpFsJg32RR0+cq1RgYtJkxz8vozrSFp0V+sAU3qQ1xD0H4ZlandPwdZycbAgmz2rhziqC/I4rKgXGVnMYDWSeXmAgyzFJ+T4cvEJcdEKVEObT+JUFSVx7WZDZAC7VL6Ff5PQQ8Ms95LlC7dkfH1zfWJcIo92xqNF2aESdCPXyA43vLkvgLAYyn3p774f+hILGCk2jveex99A477mR6kRNOW1balkgbPMuems7B81VXuzoM2X8ty/G7VA==",
    "ctl00$ContentPlaceHolder1$ToolkitScriptManager1": "ctl00$ContentPlaceHolder1$updFiltro|ctl00$ContentPlaceHolder1$btnConsultar",
    "ctl00_ContentPlaceHolder1_ToolkitScriptManager1_HiddenField": "",
    "ctl00$ContentPlaceHolder1$hdnPDF_EXCEL": "", 
    "__ASYNCPOST": "false",
    "ctl00$ContentPlaceHolder1$btnConsultar": "Consultar",
    "__EVENTTARGET": "",
    "__EVENTARGUMENT": "",
    "__LASTFOCUS": "",
}
```

- `ctl00$ContentPlaceHolder1$ddlAgentes`: `string`
  > `value` da Instituição escolhida. Adquirida na `Página de Negociação de Ativos`
- `ctl00$ContentPlaceHolder1$ddlContas`: `string`
  > `value` da Conta escolhida. Adquirida na `Escolha da Instituição`
- `ctl00$ContentPlaceHolder1$txtDataDeBolsa`: `string`
  > Data de Início do intervalo da consulta.
- `ctl00$ContentPlaceHolder1$txtDataAteBolsa`: `string`
  > Data de Término do intervalo da consulta.
- `__VIEWSTATE`: `string`
  > Adquirida na `Escolha da Instituição`
- `__VIEWSTATEGENERATOR`: `string`
  > Adquirida na `Escolha da Instituição`
- `__EVENTTARGET`: `string`
  > Adquirida na `Escolha da Instituição`

__Tipo do Retorno:__ `text/html`

__Scrap (Exemplo):__ 

Capturar os atributos desejados de todas as Negociações de Ativos no `tbody` abaixo:

```html
<tbody>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl00_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl00_lblPrazoVencimento"></span>
        </td>
        <td>
            B3SA3F
        </td>
        <td>
            B3 ON NM
        </td>
        <td class="text-right">
            2
        </td>
        <td class="text-right">
            49,75
        </td>
        <td class="text-right">
            99,50
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl01_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl01_lblPrazoVencimento"></span>
        </td>
        <td>
            CNTO3F
        </td>
        <td>
            CENTAURO ON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            48,72
        </td>
        <td class="text-right">
            48,72
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl02_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl02_lblPrazoVencimento"></span>
        </td>
        <td>
            CNTO3F
        </td>
        <td>
            CENTAURO ON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            48,00
        </td>
        <td class="text-right">
            48,00
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl03_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl03_lblPrazoVencimento"></span>
        </td>
        <td>
            CVCB3F
        </td>
        <td>
            CVC BRASIL ON NM
        </td>
        <td class="text-right">
            3
        </td>
        <td class="text-right">
            23,11
        </td>
        <td class="text-right">
            69,33
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl04_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl04_lblPrazoVencimento"></span>
        </td>
        <td>
            ENBR3F
        </td>
        <td>
            ENERGIAS BR ON NM
        </td>
        <td class="text-right">
            3
        </td>
        <td class="text-right">
            20,85
        </td>
        <td class="text-right">
            62,55
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl05_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl05_lblPrazoVencimento"></span>
        </td>
        <td>
            FLRY3F
        </td>
        <td>
            FLEURY ON ED NM
        </td>
        <td class="text-right">
            2
        </td>
        <td class="text-right">
            31,00
        </td>
        <td class="text-right">
            62,00
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl06_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl06_lblPrazoVencimento"></span>
        </td>
        <td>
            HYPE3F
        </td>
        <td>
            HYPERA ON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            41,88
        </td>
        <td class="text-right">
            41,88
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl07_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl07_lblPrazoVencimento"></span>
        </td>
        <td>
            ITUB4F
        </td>
        <td>
            ITAUUNIBANCOPN EDJ N1
        </td>
        <td class="text-right">
            3
        </td>
        <td class="text-right">
            31,54
        </td>
        <td class="text-right">
            94,62
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl08_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl08_lblPrazoVencimento"></span>
        </td>
        <td>
            LREN3F
        </td>
        <td>
            LOJAS RENNERON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            52,74
        </td>
        <td class="text-right">
            52,74
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl09_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl09_lblPrazoVencimento"></span>
        </td>
        <td>
            MGLU3F
        </td>
        <td>
            MAGAZ LUIZA ON NM
        </td>
        <td class="text-right">
            2
        </td>
        <td class="text-right">
            53,54
        </td>
        <td class="text-right">
            107,08
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl10_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl10_lblPrazoVencimento"></span>
        </td>
        <td>
            TAEE11F
        </td>
        <td>
            TAESA UNT N2
        </td>
        <td class="text-right">
            2
        </td>
        <td class="text-right">
            30,35
        </td>
        <td class="text-right">
            60,70
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl11_lblDataNegocio">03/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl11_lblPrazoVencimento"></span>
        </td>
        <td>
            WEGE3F
        </td>
        <td>
            WEG ON ED NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            48,74
        </td>
        <td class="text-right">
            48,74
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl12_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl12_lblPrazoVencimento"></span>
        </td>
        <td>
            B3SA3F
        </td>
        <td>
            B3 ON EDJ NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            37,47
        </td>
        <td class="text-right">
            37,47
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl13_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl13_lblPrazoVencimento"></span>
        </td>
        <td>
            CVCB3F
        </td>
        <td>
            CVC BRASIL ON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            12,75
        </td>
        <td class="text-right">
            12,75
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl14_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl14_lblPrazoVencimento"></span>
        </td>
        <td>
            ENBR3F
        </td>
        <td>
            ENERGIAS BR ON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            16,25
        </td>
        <td class="text-right">
            16,25
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl15_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl15_lblPrazoVencimento"></span>
        </td>
        <td>
            FLRY3F
        </td>
        <td>
            FLEURY ON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            21,04
        </td>
        <td class="text-right">
            21,04
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl16_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl16_lblPrazoVencimento"></span>
        </td>
        <td>
            HYPE3F
        </td>
        <td>
            HYPERA ON EJ NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            30,36
        </td>
        <td class="text-right">
            30,36
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl17_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl17_lblPrazoVencimento"></span>
        </td>
        <td>
            IVVB11F
        </td>
        <td>
            ISHARE SP500CI
        </td>
        <td class="text-right">
            2
        </td>
        <td class="text-right">
            145,58
        </td>
        <td class="text-right">
            291,16
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl18_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl18_lblPrazoVencimento"></span>
        </td>
        <td>
            LREN3F
        </td>
        <td>
            LOJAS RENNERON NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            35,87
        </td>
        <td class="text-right">
            35,87
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl19_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl19_lblPrazoVencimento"></span>
        </td>
        <td>
            TAEE11F
        </td>
        <td>
            TAESA UNT N2
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            25,76
        </td>
        <td class="text-right">
            25,76
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

    <tr>
        <td class="text-center">
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl20_lblDataNegocio">30/03/2020</span>
        </td>
        <td class="text-left">
            C
        </td>
        <td>
            Merc. Fracionário
        </td>
        <td>
            <span id="ctl00_ContentPlaceHolder1_rptAgenteBolsa_ctl00_rptContaBolsa_ctl00_rptAnaBolsa_ctl20_lblPrazoVencimento"></span>
        </td>
        <td>
            WEGE3F
        </td>
        <td>
            WEG ON EJ NM
        </td>
        <td class="text-right">
            1
        </td>
        <td class="text-right">
            34,13
        </td>
        <td class="text-right">
            34,13
        </td>
        <td class="text-right">
            1
        </td>
    </tr>

</tbody>
```

</p>

</details>