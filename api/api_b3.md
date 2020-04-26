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

##### Dados do pregão

Retorna dados do último pregão
> URL: `http://cotacao.b3.com.br/mds/api/v1/TradingFloorInfo`
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> ```



##### Flutuação de preço

Retorna dados dos 5 ativos com maiores altas e maiores baixas do ultimo pregão
> URL: `http://cotacao.b3.com.br/mds/api/v1/InstrumentPriceFluctuation/ibov`
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> ```


Retorna dados dos 5 ativos com maior volume de ações comercializados
> URL: `http://cotacao.b3.com.br/mds/api/v1/InstrumentTradeVolume/vista`
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> ```

#### [Derivativos](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/derivativos.htm)
TODO

#### [Índices](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/indices.htm])
TODO

#### [Outros Ativos](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/outros-ativos.htm)
TODO

#### [Cotações](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/cotacoes/)

> De acordo com testes as api's abaixo só retornam dados dos últimos 10ish dias excluindo o dia atual

Consulta possíveis códigos de ativos
> URL: `https://arquivos.b3.com.br/apinegocios/tickersymbol?q=<QUERY>`
>
> Método: GET
>
> QueryString:
> - q: `string` - texto a ser usado na busca de ativos semelhante
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> `

Recupera cotações de dado ativo em certo dia
> URL: `https://arquivos.b3.com.br/apinegocios/ticker/<CODE>/<YEAR>-<MONTH>-<DAY>`
>
> Parêmetros de URL:
> - CODE: `string` - nome do ativo a ser pesquisado
> - YEAR: `number` - ano a ser usado na busca
> - MONTH: `number` - mês a ser usado na busca
> - DAY: `number` - dias a ser usado na busca
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> `

Retorna lista de datas nas quais estão disponíveis históricos de cotações (Últimos 10 dias)
> URL: `https://arquivos.b3.com.br/apinegocios/dates`
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> ```

Retorna URL para acessar zip contendo histórico de cotações de todos os ativos dado um certo dia
> URL: `https://arquivos.b3.com.br/apinegocios/tickercsv/<YEAR>-<MONTH>-<DAY>`
>
> Parêmetros de URL:
> - YEAR: `number` - ano a ser usado na busca
> - MONTH: `number` - mês a ser usado na busca
> - DAY: `number` - dias a ser usado na busca
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> `

Retorna URL para acessar zip contendo histórico de cotações de certo ativos dado um certo dia
> URL: `https://arquivos.b3.com.br/apinegocios/tickercsv/<CODE>/<YEAR>-<MONTH>-<DAY>`
>
> Parêmetros de URL:
> - CODE: `string` - nome do ativo a ser pesquisado
> - YEAR: `number` - ano a ser usado na busca
> - MONTH: `number` - mês a ser usado na busca
> - DAY: `number` - dias a ser usado na busca
>
> Método: GET
>
> Tipo de retorno: `application/json`
> 
> Exemplo:
> ```
> TODO
> ```
> Descrição de campos:
> ```
> TODO
> `

#### [Renda fixa](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/cotacoes/renda-fixa/)
TODO

### [Histórico](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/historico/)
TODO

### [Consultas](http://www.b3.com.br/pt_br/market-data-e-indices/servicos-de-dados/market-data/consultas/)
TODO