# Dados históricos (API B3)

Documentação sobre como interpretar os arquivos de dados históricos contidos dentro dos zips disponibilizados pela B3.

Esses dados históricos disponibilizam as cotações diárias dos ativos a partir de 1986.

Coleção de URLs base das APIs usadas:

- `http://www.b3.com.br/en_us/market-data-and-indices/data-services/market-data/historical-data/equities/historical-quotes/`

## [Estrutura do arquivo](http://www.b3.com.br/data/files/65/50/AD/26/29C8B51095EE46B5790D8AA8/HistoricalQuotations_B3.pdf)

### Índices
TODO

### Nomenclatura dos arquivos

Dentro do zip chamado `COTAHIST.AAAAA.ZIP` haverá um único arquivo chamado ` COTAHIST.AAAA.TXT`. Em ambos os casos, os
últimos quatro `A`s (`AAAA`) representam o ano dos dados contidos no arquivo (e.g. ` COTAHIST.A2019.TXT`).  

### Conteúdo de um arquivo

Cada arquivo desse tipo será formado por uma linha de cabeçalho, seguida de linhas de dados propriamente dito. Por fim a
última linha é composta por um rodapé (quase idêntico ao cabeçalho).

### Conteúdo de uma linha de cabeçalho

A linha de cabeçalho se dá no seguinte formato:

```text
00COTAHIST.2019BOVESPA 20191230                                                                                                                                                                                                                      
```

É importante levar em conta os espaços pois a formatação do arquivo exige que todas as linhas possuam a mesma quantidade
de caracteres. A divisão é feita por posição de caracter, em vez de conter algum tipo de divisor de campos.

A tabela a seguir descreve cada campo do cabeçalho:

| Campo                      | Valor                      | Tipo   | Tamanho (caracteres) | Offset (caracteres) |
|----------------------------|----------------------------|--------|----------------------|---------------------|
| Tipo de registro (TIPREG)  | Fixo `00`                  |  Int   | 2                    | 0                   |
| Nome do arquivo            | Fixo `COTAHIST.AAAA`       | String | 13                   | 2                   |
| Fonte dos dados            | Fixo `BOVESPA`             | String | 8                    | 15                  |
| Data de geração do arquivo | Data no formato `YYYYMMDD` |  Int   | 8                    | 23                  |
| Reservado                  | Em branco de propósito     | String | 214                  | 31                  |

Dessa forma a decomposição do exemplo acima se daria da seguinte forma:

| Campo                      | Valor                    |
|----------------------------|--------------------------|
| Tipo de registro (TIPREG)  | `00`                     |
| Nome do arquivo            | `COTAHIST.2019`          |
| Fonte dos dados            | `BOVESPA `               |
| Data de geração do arquivo | `20191230`               |
| Reservado                  | 214 caracteres de espaço |

### Conteúdo de uma linha de dado

Após o cabeçalho começam as linhas com conteúdo propriamente dito. O arquivo ` COTAHIST.A2019.TXT`, por exemplo, contem
779232 linhas. Segue abaixo um exemplo de linha de conteúdo. 

```text
012019012402ITSA4       010ITAUSA      PN      N1   R$  000000000130000000000013350000000001300000000000131500000000013200000000001319000000000132040515000000000038730300000000050950264600000000000000009999123100000010000000000000BRITSAACNPR7389
```

Todas as linhas de conteúdo seguem a seguinte formatação:

**Atenção:** Os campos marcados como `Fixed Point` possuem uma quantidade fixa de dígitos antes e depois da vírgula. A
vírgula não é aparece no texto. Por tanto o um campo do tipo `Fixed Point (%11.2f)` indica um número com 11 dígitos
antes da vírgula e 2 dígitos após a vírgula (e.g. o número `00123456789.00` apareceria como `0012345678900`).  

| Campo                                                                                           | Valor                                                                                                                                                                                                                                                                                                       | Tipo                 | Tamanho (caracteres) | Offset (caracteres) |
|-------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|----------------------|---------------------|
| Tipo de registro (TIPREG)                                                                       | Fixo `01`                                                                                                                                                                                                                                                                                                   | Número               | 2                    | 0                   |
| Data da coleta                                                                                  | Data no formato `YYYYMMDD`                                                                                                                                                                                                                                                                                  | Int                  | 8                    | 2                   |
| Código BDI (CODBDI)                                                                             | Índice usado para classificar os ativos no informativo diário. Verificar tabela Códigos BDI.                                                                                                                                                                                                                | String               | 2                    | 10                  |
| Ticker (CODNEG)                                                                                 | Símbolo do ativo (ticker symbol)                                                                                                                                                                                                                                                                            | String               | 12                   | 12                  |
| Tipo de mercado (TPMERC)                                                                        | Índice do tipo de mercado que o ativo está registrado, de acordo com a tabela de Tipos de Mercado                                                                                                                                                                                                           | Int                  | 3                    | 24                  |
| Abreviação da empresa (NOMRES)                                                                  | Abreviação do nome da empresa associada ao ativo                                                                                                                                                                                                                                                            | String               | 12                   | 27                  |
| Especificação do ativo (ESPECI)                                                                 | Tipo do ativo (ON, PN, BDR, direito de subscrição, etc). Verificar a tabela Especificações de Ativos.                                                                                                                                                                                                       | String               | 10                   | 39                  |
| Tempo em dias para expiração no mercado futuro (PRAZOT)                                         | Quantidade de dias para expiração do contrato no mercado futuro (não sei se vale para opções também).                                                                                                                                                                                                       | String               | 3                    | 49                  |
| Moeda de referência (MODREF)                                                                    | Moeda utilizada para negociar o ativo em questão                                                                                                                                                                                                                                                            | String               | 4                    | 52                  |
| Preço de abertura (PREABE)                                                                      | Preço do ativo ao abrir o mercado (i.e. preço após a realização da primeira compra do pregão)                                                                                                                                                                                                               | Fixed Point (%11.2f) | 13                   | 56                  |
| Preço máximo (PREMAX)                                                                           | Preço máximo do ativo nesse pregão                                                                                                                                                                                                                                                                          | Fixed Point (%11.2f) | 13                   | 69                  |
| Preço mínimo (PREMIN)                                                                           | Preço mínimo do ativo nesse pregão                                                                                                                                                                                                                                                                          | Fixed Point (%11.2f) | 13                   | 82                  |
| Preço médio (PREMED)                                                                            | Média dos preços de um ativo durante o pregão                                                                                                                                                                                                                                                               | Fixed Point (%11.2f) | 13                   | 95                  |
| Preço de fechamento (PREULT)                                                                    | Preço do ativo ao fechar o mercado (i.e. preço após a realização da última compra do pregão)                                                                                                                                                                                                                | Fixed Point (%11.2f) | 13                   | 108                 |
| Preço da melhor oferta de compra (PREOFC)                                                       | Preço da melhor oferta de compra (bid) do pregão                                                                                                                                                                                                                                                            | Fixed Point (%11.2f) | 13                   | 121                 |
| Preço da melhor oferta de venda (PREOFV)                                                        | Preço da melhor oferta de venda (ask) do pregão                                                                                                                                                                                                                                                             | Fixed Point (%11.2f) | 13                   | 134                 |
| Total de trades (TOTNEG)                                                                        | Número de trades realizados com esse ativo durante o pregão desse dia                                                                                                                                                                                                                                       | Int                  | 5                    | 147                 |
| Quantidade total negociada (QUATOT)                                                             | Quantidade de unidades do ativo que foram negociadas (lembrando que um único trade pode conter 100 unidades de um ativo, por exemplo)                                                                                                                                                                       | Int                  | 18                   | 152                 |
| Volume total negociado (VOLTOT)                                                                 | Volume total do ativo negociado (em moeda fiat)                                                                                                                                                                                                                                                             | Fixed Point (%16.2f) | 18                   | 170                 |
| Strike de opções ou quantidade de contratos no mercado futuro (PREEXE)                          | Se estivermos falando de opções, é o preço de exercício (strike) da opção. Se estivermos falando de contratos futuros é a quantidade de contratos                                                                                                                                                           | Fixed Point (%11.2f) | 13                   | 188                 |
| Strike ou quantidade de contratos de opções ou indicador de correção do mercado futuro (INDOPC) | Se estivermos falando de contratos futuros, é o índice de correção de acordo com a tabela Correção de Contratos. Se estivermos falando de opções, pode ser o strike ou a quantidade de contratos(?) Não faz sentido visto que esse campo tem tamanho 1.                                                     | Int                  | 1                    | 201                 |
| Data de vencimento (DATVEN)                                                                     | Data de vencimento da opção ou do contrato futuro no formato `YYYYMMDD`                                                                                                                                                                                                                                     | Int                  | 8                    | 202                 |
| Fator de cotação (FATCOT)                                                                       | Não entendi muito bem, descrição original: "`1` = unit quote; `1000` = quote per lot of one thousand"                                                                                                                                                                                                       | Int                  | 7                    | 210                 |
| Strike de opções em pontos ou quantidade de contratos em pontos de contratos futuros (PTOEXE)   | Não tenho conhecimento suficiente pra traduzir isso fielmente, original: "For those referenced in US Dollars, each point corresponds to the value, in currency, of one one-hundreth of the interbank commercial dollar's average closing rate of the previous day, or, in other words, 1 point = 1/100 USD" | Fixed Point (%7.6f)  | 13                   | 217                 |
| Código do ativo no sistema ISIN ou código interno do ativo (CODISI)                             | Código do papel no sistema ISIN desde 05/15/1995                                                                                                                                                                                                                                                            | String               | 12                   | 230                 |
| Número de distribuição do ativo (DISMES)                                                        | Número sequencial do ativo de acordo com a lei vigente (whatever that means...)                                                                                                                                                                                                                             | Int                  | 3                    | 242                 |

Dessa forma, o exemplo dado acima ficaria:

| Campo                                                                                           | Valor                   | Interpretação                   |
|-------------------------------------------------------------------------------------------------|-------------------------|---------------------------------|
| Tipo de registro (TIPREG)                                                                       | `01`                    | Cotação diária                  |
| Data da coleta                                                                                  | `20190124`              | 24/01/2019                      |
| Código BDI (CODBDI)                                                                             | `02`                    | Lote (100 unidades)             |
| Ticker (CODNEG)                                                                                 | `ITSA4       `          | ITSA4                           |
| Tipo de mercado (TPMERC)                                                                        | `010`                   | Espécie/spot (cash)             |
| Abreviação da empresa (NOMRES)                                                                  | `ITAUSA      `          | Itaúsa                          |
| Especificação do ativo (ESPECI)                                                                 | `PN      N1`            | Preferencial                    |
| Tempo em dias para expiração no mercado futuro (PRAZOT)                                         | *3 caracteres em branco | N/A                             |
| Moeda de referência (MODREF)                                                                    | `R$  `                  | Real brasileiro (BRL)           |
| Preço de abertura (PREABE)                                                                      | `0000000001300`         | R$13,00                         |
| Preço máximo (PREMAX)                                                                           | `0000000001335`         | R$13,35                         |
| Preço mínimo (PREMIN)                                                                           | `0000000001300`         | R$13,00                         |
| Preço médio (PREMED)                                                                            | `0000000001315`         | R$13,15                         |
| Preço de fechamento (PREULT)                                                                    | `0000000001320`         | R$13,20                         |
| Preço da melhor oferta de compra (PREOFC)                                                       | `0000000001319`         | R$13,19                         |
| Preço da melhor oferta de venda (PREOFV)                                                        | `0000000001320`         | R$13,20                         |
| Total de trades (TOTNEG)                                                                        | `40515`                 | 40.515 lotes negociados         |
| Quantidade total negociada (QUATOT)                                                             | `000000000038730300`    | 38.730.300 unidades negociadas  |
| Volume total negociado (VOLTOT)                                                                 | `000000050950264600`    | R$509.502.646,00 negociados     |
| Strike de opções ou quantidade de contratos no mercado futuro (PREEXE)                          | `0000000000000`         | N/A                             |
| Strike ou quantidade de contratos de opções ou indicador de correção do mercado futuro (INDOPC) | `0`                     | N/A                             |
| Data de vencimento (DATVEN)                                                                     | `99991231`              | N/A                             |
| Fator de cotação (FATCOT)                                                                       | `0000001`               | 1 unidade de lote com 100 ações |
| Strike de opções em pontos ou quantidade de contratos em pontos de contratos futuros (PTOEXE)   | `0000000000000`         | N/A                             |
| Código do ativo no sistema ISIN ou código interno do ativo (CODISI)                             | `BRITSAACNPR7`          | BRITSAACNPR7                    |
| Número de distribuição do ativo (DISMES)                                                        | `389`                   | 389                             |


### Rodapé

Ao final de cada arquivo está um rodapé. O rodapé contem exatamente os mesmos campos que o cabeçalho acrescido no número
de registros (número de linhas) do arquivo.

| Campo                      | Valor                             | Tipo   | Tamanho (caracteres) | Offset (caracteres) |
|----------------------------|-----------------------------------|--------|----------------------|---------------------|
| Tipo de registro (TIPREG)  | Fixo `99`                         |  Int   | 2                    | 0                   |
| Nome do arquivo            | Fixo `COTAHIST.AAAA`              | String | 13                   | 2                   |
| Fonte dos dados            | Fixo `BOVESPA`                    | String | 8                    | 15                  |
| Data de geração do arquivo | Data no formato `YYYYMMDD`        |  Int   | 8                    | 23                  |
| Total de registros         | Número de linhas total do arquivo |  Int   | 8                    | 23                  |
| Reservado                  | Em branco de propósito            | String | 214                  | 31                  |


### Tabelas auxiliares

#### Códigos BDI

|Código | Descrição |
|-------|-----------|
|02 | ROUND LOT |
|05 | BMFBOVESPA REGULATIONS SANCTION |
|06 | STOCKS OF COS. UNDER REORGANIZATION |
|07 | EXTRAJUDICIAL RECOVERY |
|08 | JUDICIAL RECOVERY |
|09 | TEMPORARY ESPECIAL MANAGEMENT |
|10 | RIGHTS AND RECEIPTS |
|11 | INTERVENTION |
|12 | REAL ESTATE FUNDS |
|14 | INVESTMENT CERTIFICATES / DEBENTURES / PUBLIC DEBT SECURITIES |
|18 | BONDS |
|22 | BONUSES (PRIVATE) |
|26 | POLICIES / BONUSES / PUBLIC SECURITIES |
|32 | EXERCISE OF INDEX CALL OPTIONS |
|33 | EXERCISE OF INDEX PUT OPTIONS |
|38 | EXERCISE OF CALL OPTIONS |
|42 | EXERCISE OF PUT OPTIONS |
|46 | AUCTION OF NON-QUOTED SECURITIES |
|48 | PRIVATIZATION AUCTION |
|49 | AUCTION OF ECONOMICAL RECOVERY FUND OF ESPIRITO SANTO STATE |
|50 | AUCTION |
|51 | FINOR AUCTION |
|52 | FINAM AUCTION |
|53 | FISET AUCTION |
|54 | AUCTION OF SHARES IN ARREARS |
|56 | SALES BY COURT ORDER |
|58 | OTHERS |
|60 | SHARE SWAP |
|61 | GOAL |
|62 | TERM |
|66 | DEBENTURES WITH MATURITY DATES OF UP TO 3 YEARS |
|68 | DEBENTURES WITH MATURITY DATES GREATER THAN 3 YEARS |
|70 | FORWARD WITH CONTINUOUS MOVEMENT |
|71 | FORWARD WITH GAIN RETENTION |
|74 | INDEX CALL OPTIONS |
|75 | INDEX PUT OPTIONS |
|78 | CALL OPTIONS |
|82 | PUT OPTIONS |
|83 | DEBENTURES AND PROMISSORY NOTES BOVESPAFIX |
|84 | DEBENTURES AND PROMISSORY NOTES SOMAFIX |
|90 | REGISTERED TERM VIEW |
|96 | FACTIONARY |
|99 | GRAND TOTAL |

#### Especificações de ativos

|Código |Descrição                                                                     |
|-------|------------------------------------------------------------------------------|
|BDR    |BDR                                                                           |
|BNS    |SUBSCRIPTION BONUS FOR MISCELLANEOUS                                          |
|BNS B/A|SUBSCRIPTION BONUS FOR PREFERRED SHARES                                       |
|BNS ORD|SUBSCRIPTION BONUS FOR COMMON SHARES                                          |
|BNS P/A|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS A                               |
|BNS P/B|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS B                               |
|BNS P/C|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS C                               |
|BNS P/D|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS D                               |
|BNS P/E|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS E                               |
|BNS P/F|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS F                               |
|BNS P/G|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS G                               |
|BNS P/H|SUBSCRIPTION BONUS FOR PREFERRED SHARES CLASS H                               |
|BNS PRE|SUBSCRIPTION BONUS FOR PREFERRED SHARES                                       |
|CDA    |COMMON SHARE DEPOSIT CERTIFICATE                                              |
|CI     |INVESTMENT FUND                                                               |
|CPA    |ADDITIONAL CONSTRUCTION AND OPERATION POTENTIAL                               |
|DIR    |SUBSCRIPTION RIGHTS – MISCELLANEOUS (BONUS, DEBENTURES, ETC)                  |
|DIR ORD|SUBSCRIPTION RIGHTS TO COMMON SHARES                                          |
|DIR P/A|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS A                               |
|DIR P/B|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS B                               |
|DIR P/C|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS C                               |
|DIR P/D|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS D                               |
|DIR P/E|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS E                               |
|DIR P/F|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS F                               |
|DIR P/G|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS G                               |
|DIR P/H|SUBSCRIPTION RIGHTS TO PREFERRED SHARES CLASS H                               |
|DIR PR |SUBSCRIPTION RIGHTS TO REDEEMABLE PREF. SHARES                                |
|DIR PRA|SUBSCRIPTION RIGHTS TO REDEEMABLE PREF. SHARES CLASS A                        |
|DIR PRB|SUBSCRIPTION RIGHTS TO REDEEMABLE PREF. SHARES CLASS B                        |
|DIR PRC|SUBSCRIPTION RIGHTS TO REDEEMABLE PREF. SHARES CLASS C                        |
|DIR PRE|SUBSCRIPTION RIGHTS TO PREFERRED SHARES                                       |
|LFT    |FINANCIAL TREASURY BILL                                                       |
|M1 REC |RECEIPT OF SUBSCRIPTION TO MISCELLANEOUS                                      |
|ON     |NOMINATIVE COMMON SHARES                                                      |
|ON P   |NOMINATIVE COMMON SHARES WITH DIFFERENTIATED RIGHTS                           |
|ON REC |RECEIPT OF SUBSCRIPTION FOR COMMON SHARES                                     |
|OR     |NOMINATIVE COMMON REDEEMABLE SHARES                                           |
|OR P   |NOMINATIVE COMMON REDEEMABLE SHARES W/ DIFFERENTIATED RIGHTS                  |
|PCD    |CONSOLIDATED DEBT POSITION                                                    |
|PN     |NOMINATIVE PREFERRED SHARES                                                   |
|PN P   |NOMINATIVE PREFERRED SHARES WITH DIFFERENTIATED RIGHTS                        |
|PN REC |RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES                                  |
|PNA    |NOMINATIVE PREFERRED SHARES CLASS A                                           |
|PNA P  |NOMINATIVE PREFERRED SHARES CLASS A W/ DIFFERENTIATED RIGHTS                  |
|PNA REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS A                          |
|PNB    |NOMINATIVE PREFERRED SHARES CLASS B                                           |
|PNB P  |NOMINATIVE PREFERRED SHARES CLASS B W/ DIFFERENTIATED RIGHTS                  |
|PNB REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS B                          |
|PNC    |NOMINATIVE PREFERRED SHARES CLASS C                                           |
|PNC P  |NOMINATIVE PREFERRED SHARES CLASS C W/ DIFFERENTIATED RIGHTS                  |
|PNC REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS C                          |
|PND    |NOMINATIVE PREFERRED SHARES CLASS D                                           |
|PND P  |NOMINATIVE PREFERRED SHARES CLASS D W/ DIFFERENTIATED RIGHTS                  |
|PND REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS D                          |
|PNE    |NOMINATIVE PREFERRED SHARES CLASS                                             |
|PNE P  |NOMINATIVE PREFERRED SHARES CLASS E W/ DIFFERENTIATED RIGHTS                  |
|PNE REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS E                          |
|PNF    |NOMINATIVE PREFERRED SHARES CLASS F                                           |
|PNF P  |NOMINATIVE PREFERRED SHARES CLASS F W/ DIFFERENTIATED RIGHTS                  |
|PNF REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS F                          |
|PNG    |NOMINATIVE PREFERRED SHARES CLASS G                                           |
|PNG P  |NOMINATIVE PREFERRED SHARES CLASS G W/ DIFFERENTIATED RIGHTS                  |
|PNG REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS G                          |
|PNH    |NOMINATIVE PREFERRED SHARES CLASS H                                           |
|PNH P  |NOMINATIVE PREFERRED SHARES CLASS H W/ DIFFERENTIATED RIGHTS                  |
|PNH REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES CLASS H                          |
|PNR    |NOMINATIVE PREFERRED REDEEMABLE SHARES                                        |
|PNV    |NOMINATIVE PREFERRED SHARES WITH VOTING RIGHTS                                |
|PNV P  |NOMINATIVE PREFERRED SHARES CLASS V W/ DIFFERENTIATED RIGHTS                  |
|PNV REC|RECEIPT OF SUBSCRIPTION FOR PREFERRED SHARES W/VOTING RIGHTS                  |
|PR P   |NOMINATIVE PREFERRED REDEEMABLE SHARES W/ DIFFERENTIATED RIGHTS               |
|PRA    |NOMINATIVE PREFERRED SHARES CLASS A REDEEMABLE                                |
|PRA P  |NOMINATIVE PREFERRED SHARES CLASS A REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRA REC|RECEIPT OF SUBSCRIPTION RIGHTS TO REDEEMABLE SHARES CLASS A                   |
|PRB    |NOMINATIVE PREFERRED SHARES CLASS B REDEEMABLE                                |
|PRB P  |NOMINATIVE PREFERRED SHARES CLASS B REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRB REC|RECEIPT OF SUBSCRIPTION RIGHTS TO REDEEMABLE SHARES CLASS B                   |
|PRC    |NOMINATIVE PREFERRED SHARES CLASS C REDEEMABLE                                |
|PRC P  |NOMINATIVE PREFERRED SHARES CLASS C REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRC REC|RECEIPT OF SUBSCRIPTION RIGHTS TO REDEEMABLE SHARES CLASS C                   |
|PRD    |NOMINATIVE PREFERRED SHARES CLASS D REDEEMABLE                                |
|PRD P  |NOMINATIVE PREFERRED SHARES CLASS D REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRE    |NOMINATIVE PREFERRED SHARES CLASS E REDEEMABLE                                |
|PRE P  |NOMINATIVE PREFERRED SHARES CLASS E REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRF    |NOMINATIVE PREFERRED SHARES CLASS F REDEEMABLE                                |
|PRF P  |NOMINATIVE PREFERRED SHARES CLASS F REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRG    |NOMINATIVE PREFERRED SHARES CLASS G REDEEMABLE                                |
|PRG P  |NOMINATIVE PREFERRED SHARES CLASS G REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRH    |NOMINATIVE PREFERRED SHARES CLASS H REDEEMABLE                                |
|PRH P  |NOMINATIVE PREFERRED SHARES CLASS H REDEEMABLE W/ DIFFERENTIATED RIGHTS       |
|PRV    |NOMINATIVE PREFERRED SHARES WITH VOTING RIGHTS REDEEMABLE                     |
|PRV P  |NOMINATIVE PREFERRED SHARES REDEEM. W/ DIFFERENTIATED RIGHTS AND VOTING RIGHTS|
|R      |BASKET OF NOMINATIVE SHARES                                                   |
|REC    |RECEIPT OF SUBSCRIPTION MISCELLANEOUS                                         |
|REC PR |RECEIPT OF SUBSCRIPTION TO REDEEMABLE PREF. SHARES                            |
|RON    |BASKET OF NOMINATIVE COMMON SHARES                                            |
|TPR    |PERPETUAL BONDS WITH VARIABLE INCOME BASED ON ROYALTIES                       |
|UNT    |SHARE DEPOSIT CERTIFICATE – MISCELLANEOUS                                     |
|UNT    |UNITS                                                                         |
|UP     |ROGATORY LETTERS                                                              |
|WRT    |DEBENTURE WARRANTS                                                            |

#### Correção de Contratos

| Código | Sigla | Descrição |
|--------|-------|-----------|
| 1 | US$ | CORRECTION BY THE DOLLAR RATE |
| 2 | TJLP | CORRECTION BY THE TJLP |
| 8 | IGPM | (PROTECTED) CORRECTION BY THE IGP-M - PROTECTED OPTIONS |
| 9 | URV | CORRECTION BY THE URV |


#### Tipos de Mercado

| Código | Descrição                             |
|--------|---------------------------------------|
| 010    | Espécie (cash)                        |
| 012    | Exercício de opções do tipo CALL      |
| 013    | Exercício de opções do tipo PUT       |
| 017    | Leilão                                |
| 020    | Mercado fracionado                    |
| 030    | Mercado a termo                       |
| 050    | Mercado futuro com retenção de ganho  |
| 060    | Mercado futuro com movimento contínuo |
| 070    | Opções do tipo CALL                   |
| 080    | Opções do tipo PUT                    |