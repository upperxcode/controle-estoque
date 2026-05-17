# Design System: controle-estoque

## 1. Domain and Scope (PRD)
**Problem:** controle de estoque
**Functional:** estoque, entrada, saída, painel de gerenciamento visual (UI)
**Non-Functional:** performance, segurança, interface ágil e responsiva

## 2. Architecture and Engineering
- **Base:** Client-Server / CRUD
- **Persistence:** sql
- **Philosophies:** kiss, dry, solid, Separation of Concerns (Frontend/Backend desacoplados)

## 3. Design Patterns
- **Design Patterns (GoF):** factory
- **Data Patterns:** repository
- **Frontend Patterns:** Component-Based Architecture, Store/State Management

## 4. Opinative Stack and Dependencies
- **Backend Stack:** Go Backend (Gin/Sqlx)
- **Backend Plugins:** github.com/gin-gonic/gin, github.com/jmoiron/sqlx
- **Frontend Stack:** Vue.js 3 (Composition API) + Vite
- **Frontend Plugins:** vue-router (navegação), pinia (gerenciamento de estado), axios (cliente HTTP), tailwindcss (estilização rápida)

### Stack Configuration
- O Backend expõe a API na porta `8080`.
- O Frontend roda de forma independente (ex: porta `5173` via Vite) e utiliza um proxy reverso no ambiente de desenvolvimento para evitar problemas de CORS ao consumir a API em `/api/v1`.

## 5. GOLDEN RULES
> [!IMPORTANT]
> 1. **Strict Decoupling:** O frontend em Vue não deve conter regras de negócio complexas. Toda a validação crítica e cálculo de estoque deve ocorrer no backend em Go.
> 2. **Componentização CRUD:** Siga o princípio DRY criando componentes Vue genéricos e reutilizáveis para tabelas de listagem (DataTables), modais de formulário e botões de ação.
> 3. **Single Source of Truth no UI:** Use o Pinia para gerenciar o estado global do inventário no frontend, minimizando chamadas repetidas à API.

## 6. API Contract / Communication
rest (JSON payloads exclusivamente)

## 7. Customization Details and Subtleties
uses sqlite. O projeto deve ser estruturado em dois diretórios raízes principais para isolar os ambientes: `/backend` (código Go) e `/frontend` (código Vue).

## 8. Implementation Instructions
Follow project structure. Inicie o backend com `go mod init` dentro da pasta `/backend` e o frontend com `npm create vite@latest frontend -- --template vue` na raiz.


## 9. Stock entry screen
Frontend screen for entering stock entry data.

## 10. stock exit screen
Frontend screen for entering stock outflow transactions.



---
📊 **Architecture Health: 100%**
