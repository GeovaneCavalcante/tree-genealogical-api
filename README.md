
# tree-genealogical-api

Este projeto é uma API para construir e gerenciar árvores genealógicas. Ele permite que você crie, atualize, delete e consulte informações sobre pessoas e suas relações familiares.

## Pré-requisitos

Para executar este projeto, você precisará ter o seguinte instalado em sua máquina:

- Docker (para execução com Docker)
- Go (para execução sem Docker ou em modo desenvolvimento)
- [reflex](https://github.com/cespare/reflex) para execução em modo de desenvolvimento

## Como Executar

### Com Docker

```bash
make docker-build-image
make docker-run
```

### Sem Docker

```bash
make run
```

### Modo Desenvolvimento

```bash
make dev
```

## Utilizando a Aplicação

- Importe a collection do Postman que está no diretório `docs/postman`.
- Acesse o Swagger da aplicação em [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) ou utilize a [demo online](https://tree-genealogical-api-mszsnkdbwq-uc.a.run.app/docs/index.html).

## Famílias Cadastradas

Existem duas famílias cadastradas. A primeira contém os membros: Martin, Anastasia, Phoebe, Advik, Sonny, Ann, Dunny, Bruce, Clark, Eric, Jacqueline, Ariel, Melody. A segunda está definida no arquivo `database/database.go`.

## Rotas da Aplicação

- `/api/v1/person` - BREAD do recurso de pessoa.
- `/api/v1/relationship` - BREAD do recurso de relacionamento.
- `/api/v1/familytree`:
  - `GET /members/{personName}` - Retorna a árvore genealógica de uma pessoa.
  - `GET /relationship/{firstPersonName}/{secondPersonName}` - Retorna o relacionamento entre duas pessoas.
  - `GET /kinship/distance/{firstPersonName}/{secondPersonName}` - Retorna a distância de parentesco entre duas pessoas.

A API aceita JSON, XML e também YAML, mas o Swagger não suporta YAML.
Consulte a documentação para mais informações. 

## Limites e Extensões

Não existe limite de profundidade na árvore genealógica. O mapeamento de relacionamentos existe somente até bisavó. Qualquer parente não mapeado será adicionado como `Unknown Relation`. Para adicionar novos mapeamentos, atualize `kinshipTypes` e `rulesParents` no arquivo `pkg/genealogy/genealogy.go`.

Exemplo de adição de tataravó:

Declaração do novo tipo.
```go
var kinshipTypes = map[string]map[string]string{
    ...,
    greatGreatGrandSon: {"F": "GreatGreatGranddaughter", "M": "GreatGreatGrandson"},
}
```


Adicione uma nova regra de mapeamento na constante `rulesParents`, essa regra são de ancestrais diretos, (mãe, pai, vó, etc.)

É necessário mapear o tipo de relacionamento antecessor ao novo tipo, exemplo (bizavó é antecessor de tataravó, etc.)
```go
rulesParents := map[string]string{
    ...,
    greatGrandfather: greatGreatGrandSon,
}
```


ou voce pode adicionar a regra de decententes e parentes indiretos na constante `rulesChild`, exemplo: (filho, tia, sobrinho, etc.)
```go
rulesChild := map[string]string{
    ...,
    brother:          nephew,
}
```



### Executando Testes e Gerando Relatórios de Cobertura

Para executar os testes unitários da aplicação, utilize o comando:

```bash
make test
```

Para executar testes e gerar relatórios de cobertura, use:

```bash
make test-coverage
```

Este comando produz dois arquivos no diretório raiz: `coverage.out` e `coverage.html`, onde você pode visualizar os detalhes da cobertura de testes.

### Gerando Mocks

Para fins de teste, você pode gerar mocks com:

```bash
make build-mocks
```
Esse comando usa a biblioteca `go.uber.org/mock/mockgen` para criar mocks para a aplicação.

### Gerando Documentação do Swagger

Para gerar ou atualizar a documentação do Swagger para a API, execute:

```bash
make swagger
```

Esse comando utiliza a biblioteca `github.com/swaggo/swag/cmd/swag` para gerar automaticamente a documentação do Swagger com base no código.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para enviar pull requests, relatar problemas ou sugerir melhorias.

Obrigado por conferir a API Genealógica de Árvores!
