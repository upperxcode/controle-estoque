# Backlog: controle-estoque

## Sprint 1: Fundação e Estrutura Base
- [x] **1.1 Configuração do Backend e Banco de Dados**
    - [x] Conexão com SQLite estabelecida via Sqlx
    - [x] Server Gin respondendo em porta 8080
    - [x] Unit/Integration tests passing for all changes
- [x] **1.2 Scaffold do Frontend Vue.js**
    - [x] Frontend rodando via Vite com Tailwind configurado
    - [x] Pinia store inicial configurada
    - [x] Unit/Integration tests passing for all changes

## Sprint 2: Gestão de Estoque (Core)
- [x] **2.1 API de Gestão de Produtos**
    - [x] Endpoints GET, POST, PUT, DELETE para /api/v1/products
    - [x] Validação de dados implementada no Go
    - [x] Unit/Integration tests passing for all changes
- [x] **2.2 Interface de Listagem e Cadastro**
    - [x] DataTable reutilizável consumindo a API
    - [x] Formulários de cadastro com validação reativa
    - [x] Unit/Integration tests passing for all changes

## Sprint 3: Entradas, Saídas e Painel Visual
- [x] **3.1 Controle de Movimentação de Estoque**
    - [x] Transações atômicas para atualização de saldo
    - [x] Histórico de movimentações persistido
    - [x] Unit/Integration tests passing for all changes
- [x] **3.2 Painel de Gerenciamento Visual (Dashboard)**
    - [x] Cards de resumo (total itens, baixo estoque)
    - [x] Gráfico ou lista visual de movimentações recentes
    - [x] Unit/Integration tests passing for all changes
