# 🏆 Desafio Técnico – EMV Processor

## 📌 Visão Geral
Este projeto implementa um **módulo básico de processamento de transações EMV**, simulando a comunicação entre um **terminal de pagamento (POS)** e um **cartão de crédito/débito com chip**.

A solução foi desenvolvida em **Go**, com foco em **boas práticas, separação de responsabilidades, testabilidade e clareza de código**, simulando um cenário real do mercado de pagamentos.

---

## 🎯 Objetivo do Projeto
- Processar uma transação EMV a partir de dados TLV simulados
- Validar informações essenciais do cartão
- Simular autorização junto a um gateway de pagamento
- Registrar o resultado da transação

---

## 🧠 Conceitos Utilizados
- **EMV (Europay, Mastercard, Visa)**
- **TLV (Tag-Length-Value)**
- **Clean Architecture / Arquitetura Hexagonal**
- **Testes unitários**

---

## 🧱 Arquitetura
O projeto segue uma arquitetura inspirada em **Clean Architecture**, separando claramente as responsabilidades:

### 📂 Camadas

- **Domain**  
  Contém as regras de negócio puras, validações e entidades do sistema.

- **Use Case**  
  Responsável por orquestrar o fluxo da transação EMV.

- **Ports**  
  Define contratos (interfaces) para dependências externas, como gateway e persistência.

- **Infra**  
  Implementações técnicas: parser EMV, gateway mock e repositórios.

- **CMD**  
  Ponto de entrada da aplicação.

Essa separação facilita manutenção, testes e evolução do sistema.

---

## 🔄 Fluxo da Transação
1. Leitura simulada dos dados do cartão em formato TLV
2. Decodificação dos dados EMV (PAN, validade e CVM)
3. Validação das informações:
   - PAN válido (13 a 19 dígitos + algoritmo de Luhn)
   - Cartão não expirado
   - CVM suportado
4. Simulação de autorização via gateway de pagamento
5. Persistência da transação
6. Retorno do resultado (APROVADA ou REJEITADA)

---

## 🧾 TLVs Simulados
| Tag  | Descrição |
|-----|----------|
| 5A  | PAN (Primary Account Number) |
| 5F24 | Data de validade do cartão |
| 9F34 | CVM (Cardholder Verification Method) |

---

## 🛠️ Tecnologias Utilizadas
- **Go 1.22+**
- Go Modules
- Arquitetura Clean / Hexagonal
- Testes unitários com `testing`

---

## ▶️ Como Executar o Projeto

### Pré-requisitos
- Go 1.22 ou superior
- GoLand ou VS Code

### Executar aplicação
```bash
go mod tidy
go run cmd/app/main.go
```

Ao executar, o sistema processa uma transação simulada e exibe o resultado no terminal.

---

## 🧪 Testes Unitários
Os testes foram implementados para validar as regras de negócio e o fluxo principal da aplicação.

### Camadas testadas
- Validação de PAN (Luhn)
- Validação de data de validade
- Caso de uso de processamento da transação

### Executar testes
```bash
go test ./...
```

---

## 🛡️ Tratamento de Erros
- Dados TLV inválidos
- PAN fora do padrão ou inválido
- Cartão expirado
- CVM não suportado

Os erros são tratados e retornados de forma clara pelo caso de uso.

---

## 📁 Estrutura do Projeto
```
emv-processor/
├── cmd/
│   └── app/main.go
├── internal/
│   ├── domain/
│   ├── usecase/
│   ├── ports/
│   └── infra/
├── go.mod
└── README.md
```

---

## 📈 Possíveis Evoluções
- Implementar parser TLV completo conforme especificação EMV
- Persistência em banco SQLite
- Simulação de códigos de resposta do gateway
- Exposição do processamento via API REST

---

## ✅ Considerações Finais
Este projeto foi desenvolvido com foco em **qualidade de código, organização e aderência a boas práticas**, simulando de forma simples um fluxo real de processamento EMV.

Ele demonstra capacidade de estruturar soluções escaláveis, testáveis e alinhadas ao contexto de **meios de pagamento**.

