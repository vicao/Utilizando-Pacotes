# Aula de Pacotes em Go

Este repositório contém exemplos práticos de utilização de diversos pacotes fundamentais da biblioteca padrão da linguagem Go (Golang). O projeto foi estruturado de forma modular para demonstrar funcionalidades específicas de cada pacote.

## 🚀 Como executar

Certifique-se de ter o Go instalado em sua máquina (versão 1.23 ou superior).

1. Navegue até a pasta raiz do projeto.
2. Execute o arquivo principal:

```bash
go run menu.go
```

## 📦 Pacotes Abordados

O código explora os seguintes tópicos através de pacotes locais:

- **iopckg**: Demonstrações de entrada e saída (`io`, `io/ioutil`), incluindo escrita no console e criação de arquivos simples.
- **stringpckg**: Manipulação de textos com o pacote `strings` (verificação de conteúdo, contagem, etc).
- **bytespckg**: Operações com slices de bytes (`bytes`), como busca de índices.
- **ospckg**: Interação com o Sistema Operacional (`os`), cobrindo criação de diretórios, leitura de pastas e escrita de arquivos.
- **path_filepathpckg**: Manipulação de caminhos de arquivos e diretórios (`path/filepath`), incluindo a função `Walk` para percorrer árvores de diretórios.
- **errorpckg**: Implementação de erros personalizados utilizando a interface `error` e structs.
- **container_listpckg**: Uso de estruturas de dados como listas encadeadas (`container/list`).
- **sortpckg**: Algoritmos de ordenação de dados (`sort`).
- **hash_cryptopckg**: Exemplos de hashing e criptografia utilizando `hash/crc32` e `crypto/sha1`.

## 🛠 Estrutura

A estrutura de pastas do projeto segue o padrão de organização de pacotes em Go, onde cada diretório representa um pacote isolado:

```text
Utilizando Pacotes/
├── go.mod              # Definição do módulo e dependências
├── menu.go             # Ponto de entrada (main) que executa os exemplos
├── bytespckg/          # Exemplos com pacote bytes
├── container_listpckg/ # Exemplos com container/list
├── errorpckg/          # Exemplos de tratamento de erros
├── hash_cryptopckg/    # Exemplos de hash e criptografia
├── iopckg/             # Exemplos de I/O
├── ospckg/             # Exemplos de interação com OS
├── path_filepathpckg/  # Exemplos de manipulação de caminhos
├── sortpckg/           # Exemplos de ordenação
└── stringpckg/         # Exemplos de manipulação de strings
```

---

*Projeto desenvolvido para fins de estudo e treinamento em Go.*