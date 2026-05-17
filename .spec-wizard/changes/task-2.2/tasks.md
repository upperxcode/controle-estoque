# 🎯 TASK SPEC: Telas de Entrada e Saída (Frontend)
*(ID: 2.2)*

## 📋 General view
Implementação das interfaces de usuário para registro de movimentações de estoque (entrada e saída). O objetivo é fornecer uma forma rápida e segura de alterar a quantidade de produtos no estoque, garantindo que o saldo seja atualizado visualmente de imediato.

### Sub-tasks
- [x] Criar componente `MovementForm.vue` (reutilizável para entrada e saída).
- [x] Adicionar rotas `/movements/in` e `/movements/out` no Vue Router.
- [x] Criar View `MovementsView.vue` para gerenciar o estado da tela.
- [x] Integrar com a action `registerMovement` do Pinia.
- [x] Implementar validação de estoque negativo no frontend para saídas.

## 🚀 Status: completed

## ✅ Acceptance Criteria
- [x] **2.2.1** View de entrada permite selecionar produto e quantidade.
- [x] **2.2.2** View de saída valida disponibilidade antes de enviar.
- [x] **2.2.3** Lista de movimentos recentes visível para auditoria. (Implementado via integração com a store)
- [x] sw verify validation.

## 📂 linked files
| Arquivo | Descrição |
| :--- | :--- |
| frontend/src/router/index.js | Registro das novas rotas de movimentação. |
| frontend/src/views/MovementsView.vue | Container para as operações de estoque. |
| frontend/src/components/movements/MovementForm.vue | Formulário principal com seletor de produto e quantidade. |
| frontend/src/stores/inventory.js | Garantir que a store suporte a atualização imediata. |

## 📂 test files
| Arquivo | Descrição |
| :--- | :--- |
| frontend/src/components/movements/MovementForm.spec.js | Testes de validação e emissão de eventos do formulário. |
| frontend/src/stores/inventory.spec.js | Testes da action registerMovement e atualização de saldo. |
