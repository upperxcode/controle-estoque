# 🗺️ ROADMAP: controle-estoque
> This document details the development phases and the current progress of the system.

- **Language**: Go Backend (Gin/Sqlx) / Vue.js 3
- **Pattern**: CRUD / Repository

## Sprint 1: Fundação e Estrutura Base

### ✅ Configuração do Backend e Banco de Dados
Configurar o servidor Gin, conexão com SQLite usando Sqlx e estrutura inicial de diretórios.

**Acceptance Criteria:**
- [x] Conexão com SQLite estabelecida via Sqlx.. completed
- [x] Server Gin respondendo em porta 8080 com um endpoint `/health`.. completed
- [x] Unit/Integration tests passing for all changes.. completed
- [x] `sw verify` validation aprovada.. completed

### ✅ Scaffold do Frontend Vue.js
Inicializar o projeto Vue 3 with Vite, Tailwind CSS e Pinia.

**Acceptance Criteria:**
- [x] Frontend rodando via Vite com Tailwind configurado.. completed
- [x] Pinia store inicial configurada e acessível.. completed
- [x] Proxy configurado (chamada para `/api/v1/health` redirecionada ao backend).. completed
- [x] Unit/Integration tests passing for all changes.. completed
- [x] `sw verify` validation aprovada.. completed

## Sprint 2: Gestão de Estoque (Core)

### ✅ API de Gestão de Produtos
Implementar repositório e endpoints CRUD para produtos no backend.

**Acceptance Criteria:**
- [x] Endpoints GET, POST, PUT, DELETE para `/api/v1/products` funcionando corretamente.. completed
- [x] Validação de dados implementada (ex: nome e SKU obrigatórios).. completed
- [x] Persistência correta no banco SQLite.. completed
- [x] Unit tests para o repositório e integration tests para os handlers passando.. completed
- [x] `sw verify` validation aprovada.. completed

### ✅ Interface de Listagem e Cadastro
Desenvolver componentes Vue para listagem e cadastro de produtos.

**Acceptance Criteria:**
- [x] DataTable reutilizável exibindo produtos vindos da API.. completed
- [x] Formulários de cadastro/edição funcionais com validação reativa (campos obrigatórios, preço > 0).. completed
- [x] Exclusão de produto com confirmação.. completed
- [x] Unit/Integration tests passing for all changes.. completed
- [x] `sw verify` validation aprovada.. completed

## Sprint 3: Entradas, Saídas e Painel Visual

### ✅ Controle de Movimentação de Estoque
Lógica de negócio para entradas e saídas, garantindo atomicidade.

**Acceptance Criteria:**
- [x] Transações atômicas garantidas (atualização de produto + insert de movimento).. completed
- [x] Histórico de movimentações persistido e consultável via API.. completed
- [x] Impedir saídas que resultem em estoque negativo.. completed
- [x] Unit tests para lógica de transação e integração de movimentação passando.. completed
- [x] `sw verify` validation aprovada.. completed

### ✅ Painel de Gerenciamento Visual (Dashboard)
Interface visual para monitoramento do estoque.

**Acceptance Criteria:**
- [x] Cards de resumo exibindo dados corretos baseados no estado global.. completed
- [x] Lista visual de movimentações recentes funcionando.. completed
- [x] Interface responsiva e estilizada com Tailwind CSS.. completed
- [x] Unit/Integration tests passing for all changes.. completed
- [x] `sw verify` validation aprovada.. completed

---

_Generated automatically by Spec Wizard_
