# go-openapi-driven-template

## Introdução
Este repositório é um modelo de arquitetura escrito em Go. A escolha desse modelo é forçar o uso do conceito de API First (ou Design First). Os controllers(handlers) desse projeto são gerados a partir do arquivo `openapi.yaml`.

Primeiramente é necessário definir o modelo das APIs utilizando a especificação [OpenAPI 3.0](https://swagger.io/docs/specification/about/). Após a definição do modelo da API, são gerados os códigos contendo os modelos e as interfaces a serem implementadas no handler.

## Configurando a aplicação
WIP

## Rodando a aplicação
Gerar os arquivos de API:
```
make build
```

Rodar a aplicação:
```
make run
```

Caso ocorra erro de `missing go.sum...`, rode o comando abaixo:
```
make fix
```