# Implementation Proposal: feat-movements-ui

1. **Step 1: Prep & Routing**
   - Registrar as rotas no `router/index.js`.
   - Criar o componente base `MovementForm.vue`.

2. **Step 2: Component Logic**
   - Implementar o seletor de produtos (Select search seria ideal, mas usaremos Select simples inicialmente para agilidade).
   - Adicionar validações de formulário.

3. **Step 3: Integration**
   - Criar a View e conectar com a Store.
   - Garantir que `fetchProducts` seja chamado se a lista estiver vazia.

4. **Step 4: Testing & Validation**
   - Escrever testes para o componente de formulário.
   - Rodar `npm run test` e `sw verify`.
