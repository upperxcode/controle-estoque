# Technical Design: Telas de Entrada e Saída

## Component Architecture
- **MovementForm.vue**:
  - Props: `type` ('IN' | 'OUT'), `products` (Array para o Select).
  - Data: `productId`, `quantity`, `notes`.
  - Logic: 
    - Se `type === 'OUT'`, validar se `quantity <= product.quantity`.
    - Emitir evento `submit` com o payload da movimentação.

- **MovementsView.vue**:
  - Utiliza o `useInventoryStore` para buscar produtos (caso não carregados).
  - Renderiza o `MovementForm` baseado no parâmetro da rota.

## Routing Strategy
- `/movements/:type` onde type é validado como 'in' ou 'out'.
- Breadcrumbs ou navegação lateral para fácil acesso.

## State Updates
- A action `registerMovement` já realiza o "optimistic update" ou o ajuste manual do estado local após a chamada bem-sucedida à API.
