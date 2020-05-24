# API B3 (Reverse engineered)

Documentação das API's usadas pela B3 no seu website.

Coleçao de URL's base das API's usadas:

- `http://cotacao.b3.com.br/mds/api/v1/`
- `https://arquivos.b3.com.br/apinegocios/`

## [Market Data](http://www.b3.com.br/pt_br/market-data-e-indices/)

### Índices
TODO

### Cotações

#### [Ações](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/)

<details>
<summary>Informações sobre o pregão mais recente</summary>
<p>

__Descrição:__ Retorna dados do pregão mais recente.

__URL:__ `http://cotacao.b3.com.br/mds/api/v1/TradingFloorInfo`

__Método:__ `GET`

__Tipo de retorno:__ `application/json`

__Exemplo:__
```shell
curl http://cotacao.b3.com.br/mds/api/v1/TradingFloorInfo
```
```json
{
  "BizSts": {
    "cd": "OK"
  },
  "Msg": {
    "dtTm": "2020-04-24T18:23:06Z"
  },
  "TradgFlr": {
    "TradgFlrSts": {
      "desc": "E"
    },
    "grssAmt": 37812848269.99,
    "qty": 5891015
  }
}
```

__Descrição de campos:__
- `.BizSts.cd`: `string`
  > Algum tipo de indicador de estado da API, talvez indique que a API conseguiu responder com sucesso
- `.Msg.dtTm`: `string`
  > Data do pregão
- `.TradgFlr.TradgFlrSts.desc`: `string`
  > Estado do pregão. Possíveis valores:
  >  - `E`: Pregão encerrado
  >  - _outros_: Pregão em andamento
- `.TradgFlr.grssAmt`: `float?`
  > Volume do pregão
- `.TradgFlr.qty`: `int?`
  > Quantidade (?) do pregão

</p>
</details>

<details>
<summary>Ativos com maiores flutuações de preço no pregão mais recente</summary>
<p>

__Descrição:__ Retorna dados dos 5 ativos com maiores altas e dos 5 ativos com maiores baixas do pregão mais recente.

__URL:__ `http://cotacao.b3.com.br/mds/api/v1/InstrumentPriceFluctuation/ibov`

__Método:__ `GET`

__Tipo de retorno:__ `application/json`

__Exemplo:__
```shell
curl http://cotacao.b3.com.br/mds/api/v1/InstrumentPriceFluctuation/ibov
```
```json
{
  "BizSts": {
    "cd": "OK"
  },
  "Msg": {
    "dtTm": "2020-04-26 20:08:36"
  },
  "SctyHghstIncrLst": [
    {
      "SctyQtn": {
        "curPrc": 38.99,
        "prcFlcn": 6.646608315098468
      },
      "symb": "SUZB3",
      "desc": "SUZANO S.A. ON      NM"
    },
    {
      "SctyQtn": {
        "curPrc": 17.3,
        "prcFlcn": 2.0047169811320753
      },
      "symb": "KLBN11",
      "desc": "KLABIN S/A  UNT     N2"
    },
    {
      "SctyQtn": {
        "curPrc": 29.85,
        "prcFlcn": 1.530612244897959
      },
      "symb": "BRAP4",
      "desc": "BRADESPAR   PN      N1"
    },
    {
      "SctyQtn": {
        "curPrc": 43.95,
        "prcFlcn": 1.011261778901402
      },
      "symb": "VALE3",
      "desc": "VALE        ON      NM"
    },
    {
      "SctyQtn": {
        "curPrc": 10.5,
        "prcFlcn": 0.28653295128939826
      },
      "symb": "MRFG3",
      "desc": "MARFRIG     ON      NM"
    }
  ],
  "SctyHghstDrpLst": [
    {
      "SctyQtn": {
        "curPrc": 14.26,
        "prcFlcn": -12.889431887599267
      },
      "symb": "AZUL4",
      "desc": "AZUL        PN      N2"
    },
    {
      "SctyQtn": {
        "curPrc": 25.05,
        "prcFlcn": -12.748171368861025
      },
      "symb": "ELET6",
      "desc": "ELETROBRAS  PNB     N1"
    },
    {
      "SctyQtn": {
        "curPrc": 12.22,
        "prcFlcn": -12.21264367816092
      },
      "symb": "CVCB3",
      "desc": "CVC BRASIL  ON      NM"
    },
    {
      "SctyQtn": {
        "curPrc": 22.11,
        "prcFlcn": -11.736526946107784
      },
      "symb": "ELET3",
      "desc": "ELETROBRAS  ON      N1"
    },
    {
      "SctyQtn": {
        "curPrc": 6.5,
        "prcFlcn": -11.684782608695652
      },
      "symb": "VVAR3",
      "desc": "VIAVAREJO   ON      NM"
    }
  ]
}
```

__Descrição de campos:__
- `.BizSts.cd`: `string`
  > Algum tipo de indicador de estado da API, talvez indique que a API conseguiu responder com sucesso
- `.Msg.dtTm`: `string`
  > Data correspondente a quando os dados foram consultados (Empiracamente parece ser só um Date.now() no lado do servidor)
- `.SctyHghstIncrLst`: `Object[5]`
  > Lista dos 5 ativos com as maiores altas nesse pregão(?)
- `.SctyHghstIncrLst[].SctyQtn.curPrc`: `float`
  > Preço atual de um dos ativos com maior alta nesse pregão(?)
- `.SctyHghstIncrLst[].SctyQtn.prcFlcn`: `float`
  > Flutuação no preço de um dos ativos com maior alta desde o ínicio do pregão(?)
- `.SctyHghstIncrLst[].symb`: `string`
  > Simbolo de um dos ativos com maior alta nesse pregão(?)
- `.SctyHghstIncrLst[].desc`: `string`
  > Descrição de um dos ativos com maior alta nesse pregão(?)
- `.SctyHghstDrpLst`: `Object[5]`
  > Lista dos 5 ativos com as maiores baixas nesse pregão(?)
- `.SctyHghstDrpLst[].SctyQtn.curPrc`: `float`
  > Preço atual de um dos ativos com maior baixas nesse pregão(?)
- `.SctyHghstDrpLst[].SctyQtn.prcFlcn`: `float`
  > Flutuação no preço de um dos ativos com maior baixas desde o ínicio do pregão(?)
- `.SctyHghstDrpLst[].symb`: `string`
  > Simbolo de um dos ativos com maior baixas nesse pregão(?)
- `.SctyHghstDrpLst[].desc`: `string`
  > Descrição de um dos ativos com maior baixas nesse pregão(?)

</p>
</details>

<details>
<summary>Ativos mais negociados no pregão mais recente</summary>
<p>

__Descrição:__ Retorna dados dos 5 ativos mais negociados do pregão mais recente.

__URL:__ `http://cotacao.b3.com.br/mds/api/v1/InstrumentTradeVolume/vista`

__Método:__ `GET`

__Tipo de retorno:__ `application/json`

__Exemplo:__
```shell
curl http://cotacao.b3.com.br/mds/api/v1/InstrumentTradeVolume/vista
```
```json
{
  "BizSts": {
    "cd": "OK"
  },
  "Msg": {
    "dtTm": "2020-04-26 20:46:21"
  },
  "Volume": [
    {
      "grossAmt": 2575586340,
      "pricVal": 16.2,
      "scty": {
        "symb": "PETR4",
        "desc": "PETROBRAS   PN      N2"
      }
    },
    {
      "grossAmt": 2379933959,
      "pricVal": 43.95,
      "scty": {
        "symb": "VALE3",
        "desc": "VALE        ON      NM"
      }
    },
    {
      "grossAmt": 1853217460,
      "pricVal": 24.78,
      "scty": {
        "symb": "BBAS3",
        "desc": "BRASIL      ON      NM"
      }
    },
    {
      "grossAmt": 1432158691,
      "pricVal": 21.18,
      "scty": {
        "symb": "ITUB4",
        "desc": "ITAUUNIBANCOPN      N1"
      }
    },
    {
      "grossAmt": 1377093003,
      "pricVal": 6.5,
      "scty": {
        "symb": "VVAR3",
        "desc": "VIAVAREJO   ON      NM"
      }
    }
  ]
}
```

__Descrição de campos:__
- `.BizSts.cd`: `string`
  > Algum tipo de indicador de estado da API, talvez indique que a API conseguiu responder com sucesso
- `.Msg.dtTm`: `string`
  > Data correspondente a quando os dados foram consultados (Empiracamente parece ser só um Date.now() no lado do servidor)
- `.Volume`: `Object[5]`
  > Lista dos 5 ativos com maior flutuação de volume nesse pregão(?)
- `.Volume[].grossAmt`: `int`
  > Volume negociado de um dos ativos com maior flutuação de volume nesse pregão(?)
- `.Volume[].pricVal`: `int`
  > Preço de um dos ativos com maior flutuação de volume nesse pregão(?)
- `.Volume[].scty.symb`: `string`
  > Simbolo de um dos ativos com maior flutuação de volume nesse pregão(?)
- `.Volume[].scty.desc`: `string`
  > Descrição de um dos ativos com maior flutuação de volume nesse pregão(?)


</p>
</details>

#### [Derivativos](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/derivativos.htm)
TODO

#### [Índices](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/indices.htm])
TODO

#### [Outros Ativos](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/outros-ativos.htm)
TODO

#### [Cotações](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/cotacoes/)

> De acordo com testes as api's abaixo só retornam dados dos últimos 10ish dias excluindo o dia atual

<details>
<summary>Consulta de datas com dados disponíveis</summary>
<p>

__Descrição:__ Retorna lista de datas nas quais estão disponíveis os dados da cotações.
Parece não ser acurado. Existem datas fora dos períodos retornados nessa rota com dados disponíveis
Normalmente retorna lista contendo os ultimos dez dias úteis.

__URL:__ `https://arquivos.b3.com.br/apinegocios/dates`

__Método:__ `GET`

__Tipo de retorno:__ `application/json`

__Exemplo:__
```shell
curl https://arquivos.b3.com.br/apinegocios/dates
```
```json
[
  "2020-04-24",
  "2020-04-23",
  "2020-04-22",
  "2020-04-20",
  "2020-04-17",
  "2020-04-16",
  "2020-04-15",
  "2020-04-14",
  "2020-04-13",
  "2020-04-09"
]
```

__Descrição de campos:__
- `.`: `string[]`
  > Lista contendo datas com dados de cotação disponíveis

</p>
</details>

<details>
<summary>Consulta de simbolos de ativos</summary>
<p>

__Descrição:__ Dado um texto de busca, retorna uma lista de simbolos de ativos semelhantes.

__URL:__ `https://arquivos.b3.com.br/apinegocios/tickersymbol?q=<QUERY>`

__QueryString:__
- q: `string`
  > texto a ser usado na busca de ativos semelhante

__Método:__ `GET`

__Tipo de retorno:__ `application/json`

__Exemplo:__
```shell
curl https://arquivos.b3.com.br/apinegocios/tickersymbol?q=SU
```
```json
[
  "SULA11",
  "SULA11F",
  "SULA11T",
  "SULA3",
  "SULA3F",
  "SULA4",
  "SULA4F",
  "SULAC75",
  "SULAD35",
  "SULAD35E",
  "SULAD500",
  "SULAD570",
  "SULAE401",
  "SULAE438",
  "SULAE476",
  "SULAH437",
  "SULAH505",
  "SULAI660",
  "SULAO64",
  "SULAQ217"
]
```
__Descrição de campos:__
- `.`: `string[]`
  > Lista de simbolos de ativos semelhantes ao texto enviado

</p>
</details>

<details>
<summary>Cotações de ativo durante um dia</summary>
<p>

__Descrição:__ Dado um ativo e a data de um dia, com distância de aproximadamente 10 dias do presente, retorna a cotação dele durante aquele dia.

__URL:__ `https://arquivos.b3.com.br/apinegocios/ticker/:code/:year-:month-:day`

__Parâmetros:__
- `code`: `string`
  > Simbolo do ativo a ser pesquisado
- `year`: `number`
  > Ano a ser usado na busca
- `month`: `number`
  > Mês a ser usado na busca
- `day`: `number`
  > Dia a ser usado na busca

__Método:__ `GET`

__Tipo de retorno:__ `application/json`

__Exemplo:__
```shell
curl https://arquivos.b3.com.br/apinegocios/ticker/SULA3/2020-04-24
```
```json
{
  "name": "SULA3",
  "friendlyName": "SULA3",
  "columns": [
    {
      "name": "TickerSymbol",
      "friendlyName": "Papel",
      "type": 5,
      "columnAlignment": 1,
      "valueAlignment": 2
    },
    {
      "name": "Quantity",
      "friendlyName": "Quantidade",
      "type": 2,
      "columnAlignment": 1,
      "valueAlignment": 1
    },
    {
      "name": "Price",
      "friendlyName": "Preço",
      "type": 1,
      "columnAlignment": 1,
      "valueAlignment": 3
    },
    {
      "name": "TradeId",
      "friendlyName": "Número do negócio",
      "type": 2,
      "columnAlignment": 1,
      "valueAlignment": 1
    },
    {
      "name": "EntryDate",
      "friendlyName": "Data de referência",
      "type": 4,
      "format": "DD/MM/YYYY",
      "columnAlignment": 1,
      "valueAlignment": 1
    },
    {
      "name": "EntryTime",
      "friendlyName": "Hora",
      "type": 5,
      "columnAlignment": 1,
      "valueAlignment": 1
    }
  ],
  "values": [
    [
      "SULA3",
      300,
      20.9,
      690,
      "2020-04-24",
      "17:05"
    ],
    [
      "SULA3",
      100,
      20.9,
      680,
      "2020-04-24",
      "16:41"
    ],
    [
      "SULA3",
      100,
      20.9,
      670,
      "2020-04-24",
      "16:41"
    ],
    [
      "SULA3",
      100,
      20.9,
      660,
      "2020-04-24",
      "16:40"
    ],
    [
      "SULA3",
      100,
      20.9,
      650,
      "2020-04-24",
      "16:39"
    ],
    [
      "SULA3",
      100,
      20.9,
      640,
      "2020-04-24",
      "16:39"
    ],
    [
      "SULA3",
      100,
      20.9,
      630,
      "2020-04-24",
      "16:38"
    ],
    [
      "SULA3",
      100,
      20.9,
      620,
      "2020-04-24",
      "16:38"
    ],
    [
      "SULA3",
      100,
      20.9,
      610,
      "2020-04-24",
      "16:38"
    ],
    [
      "SULA3",
      100,
      20.72,
      600,
      "2020-04-24",
      "16:37"
    ],
    [
      "SULA3",
      100,
      20.9,
      590,
      "2020-04-24",
      "16:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      580,
      "2020-04-24",
      "15:52"
    ],
    [
      "SULA3",
      100,
      21.23,
      570,
      "2020-04-24",
      "15:52"
    ],
    [
      "SULA3",
      200,
      21.23,
      560,
      "2020-04-24",
      "15:51"
    ],
    [
      "SULA3",
      100,
      21.23,
      550,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      540,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      530,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      520,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      510,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      500,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      490,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.23,
      480,
      "2020-04-24",
      "15:22"
    ],
    [
      "SULA3",
      100,
      21.03,
      470,
      "2020-04-24",
      "15:07"
    ],
    [
      "SULA3",
      100,
      21.03,
      460,
      "2020-04-24",
      "14:36"
    ],
    [
      "SULA3",
      100,
      20.98,
      450,
      "2020-04-24",
      "14:21"
    ],
    [
      "SULA3",
      100,
      19.42,
      440,
      "2020-04-24",
      "13:46"
    ],
    [
      "SULA3",
      100,
      19.42,
      430,
      "2020-04-24",
      "13:46"
    ],
    [
      "SULA3",
      100,
      19.42,
      420,
      "2020-04-24",
      "13:30"
    ],
    [
      "SULA3",
      100,
      19.62,
      410,
      "2020-04-24",
      "13:26"
    ],
    [
      "SULA3",
      100,
      19.62,
      400,
      "2020-04-24",
      "13:21"
    ],
    [
      "SULA3",
      100,
      19.62,
      390,
      "2020-04-24",
      "12:33"
    ],
    [
      "SULA3",
      100,
      18.99,
      380,
      "2020-04-24",
      "12:24"
    ],
    [
      "SULA3",
      100,
      18.99,
      370,
      "2020-04-24",
      "12:13"
    ],
    [
      "SULA3",
      100,
      19.21,
      360,
      "2020-04-24",
      "12:13"
    ],
    [
      "SULA3",
      100,
      19.21,
      350,
      "2020-04-24",
      "12:13"
    ],
    [
      "SULA3",
      100,
      19.62,
      340,
      "2020-04-24",
      "12:02"
    ],
    [
      "SULA3",
      100,
      19.5,
      330,
      "2020-04-24",
      "11:56"
    ],
    [
      "SULA3",
      100,
      18.9,
      320,
      "2020-04-24",
      "11:42"
    ],
    [
      "SULA3",
      100,
      18.95,
      310,
      "2020-04-24",
      "11:42"
    ],
    [
      "SULA3",
      100,
      18.96,
      300,
      "2020-04-24",
      "11:42"
    ],
    [
      "SULA3",
      100,
      19,
      290,
      "2020-04-24",
      "11:42"
    ],
    [
      "SULA3",
      100,
      19,
      280,
      "2020-04-24",
      "11:42"
    ],
    [
      "SULA3",
      100,
      19.34,
      270,
      "2020-04-24",
      "11:41"
    ],
    [
      "SULA3",
      100,
      18.96,
      260,
      "2020-04-24",
      "11:22"
    ],
    [
      "SULA3",
      100,
      18.99,
      250,
      "2020-04-24",
      "11:20"
    ],
    [
      "SULA3",
      100,
      18.99,
      240,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      18.77,
      230,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      18.8,
      220,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      18.77,
      210,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      18.77,
      200,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      18.77,
      190,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      18.8,
      180,
      "2020-04-24",
      "11:15"
    ],
    [
      "SULA3",
      100,
      19.31,
      170,
      "2020-04-24",
      "11:13"
    ],
    [
      "SULA3",
      600,
      19.9,
      160,
      "2020-04-24",
      "10:58"
    ],
    [
      "SULA3",
      100,
      19.9,
      150,
      "2020-04-24",
      "10:54"
    ],
    [
      "SULA3",
      100,
      19.9,
      140,
      "2020-04-24",
      "10:54"
    ],
    [
      "SULA3",
      100,
      19.9,
      130,
      "2020-04-24",
      "10:54"
    ],
    [
      "SULA3",
      100,
      19.9,
      120,
      "2020-04-24",
      "10:54"
    ],
    [
      "SULA3",
      100,
      19.9,
      110,
      "2020-04-24",
      "10:54"
    ],
    [
      "SULA3",
      100,
      20,
      100,
      "2020-04-24",
      "10:54"
    ],
    [
      "SULA3",
      100,
      21.27,
      90,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      21.27,
      80,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      21.27,
      70,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      21.27,
      60,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      21.27,
      50,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      20.24,
      40,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      20.24,
      30,
      "2020-04-24",
      "10:47"
    ],
    [
      "SULA3",
      100,
      21.27,
      20,
      "2020-04-24",
      "10:25"
    ],
    [
      "SULA3",
      100,
      20.1,
      10,
      "2020-04-24",
      "10:20"
    ]
  ]
}
```

__Descrição de campos:__
- `.name`: `string`
  > Nome do ativo
- `.friendlyName`: `string`
  > Nome amigável do ativo (Empiricamente parece ser a mesma coisa do `.name`)
- `.columns`: `Object[N]`
  > Lista determinando as colunas da tabela de cotações
- `.columns[].name`: `string`
  > Nome da coluna
- `.columns[].friendlyName`: `string`
  > Nome amigável da coluna (Empiricamente parece ser a versão traduzida de `.columns[].name`)
- `.columns[].type`: `int`
  > Tipo de dado da coluna. Valores possíveis (Deduzidos):
  >   - `1`: `float`
  >   - `2`: `int`
  >   - `3`: UNKNOWN
  >   - `4`: `date`
  >   - `5`: `string`
  >   - ...
- `.columns[].columnAlignment`: `int`
  > Alinhamento da coluna (???)
- `.columns[].valueAlignment`: `int`
  > Alinhamento do valor (???)
- `.values`: `Any[N][...]`
  > Lista das linhas contendo os valores correspondentes as colunas determinadas em `.columns` (Valores da tabela)

</p>
</details>

<details>
<summary>Baixar zip do histórico de cotações diário</summary>
<p>

__Descrição:__ Retorna URL para acessar zip contendo histórico de cotações de todos os ativos de dado dia.

__URL:__ `https://arquivos.b3.com.br/apinegocios/tickercsv/:year-:month-:day`

__Parêmetros:__
- `year`: `number`
  > Ano a ser usado na busca
- `month`: `number`
  > Mês a ser usado na busca
- `day`: `number`
  > Dia a ser usado na busca

__Método:__ `GET`

__Tipo de retorno:__ `text/plain`

__Exemplo:__
```shell
curl https://arquivos.b3.com.br/apinegocios/tickercsv/2020-04-24
```
```
https://up2dataweb.blob.core.windows.net/raptor/TradeIntraday_20200424_1.zip?sv=2019-07-07&ss=b&srt=o&spr=https&se=2020-04-27T00%3A59%3A21Z&sp=r&sig=L4Cv%2BQmTuxAizbrOHu7tj3DOvXLYrqBaZH5yzzNFvdI%3D
```

</p>
</details>

<details>
<summary>Baixar zip do histórico de cotações diário de dado ativo</summary>
<p>

__Descrição:__ Retorna URL para acessar zip contendo histórico de cotações de certo ativos dado um certo dia.

__URL:__ `https://arquivos.b3.com.br/apinegocios/tickercsv/:code/:year-:month-:day`

__Parêmetros:__
- `code`: `string`
  > Simbolo do ativo a ser pesquisado
- `year`: `number`
  > Ano a ser usado na busca
- `month`: `number`
  > Mês a ser usado na busca
- `day`: `number`
  > Dia a ser usado na busca

__Método:__ `GET`

__Tipo de retorno:__ `text/plain`

__Exemplo:__
```shell
curl https://arquivos.b3.com.br/apinegocios/tickercsv/SULA3/2020-04-24
```
```
https://up2dataweb.blob.core.windows.net/raptor/TradeIntraday_SULA3_20200424_1.zip?sv=2019-07-07&ss=b&srt=o&spr=https&se=2020-04-27T01%3A01%3A38Z&sp=r&sig=wIGhSQSvAsZk4pGNlgOYeZezX31k%2FHPiou8AXS3fLIU%3D
```

</p>
</details>

#### [Renda fixa](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/renda-fixa/)
TODO

### [Histórico](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/historico/)
> As rotas abaixo retornam dados relcaionados ao históricos de preço de ações

<details>
<summary>Baixar Zip com cotações de todos os ativos relativo a um dado ano</summary>
<p>

__Descrição:__ Retorna um arquivo Zip contendo um arquivo de texto com dados diários de todos os tickers da bolsa ao longo do ano.

__URL:__ `http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A:year.ZIP`

__Parêmetros:__
- `year`: `number`
  > Ano a ser usado na busca (a partir de 1986).

__Método:__ `GET`

__Tipo de retorno:__ `application/x-zip-compressed`

__Exemplo:__
```shell
curl http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A2020.ZIP
```
A resposta será o binário de um Zip.

</p>
</details>

### [Consultas](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/consultas/)
TODO