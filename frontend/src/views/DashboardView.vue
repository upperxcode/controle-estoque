<template>
  <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
    <h2 class="text-3xl font-bold text-gray-900 mb-8">Painel de Controle</h2>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
      <StatCard title="Total de Itens" :value="store.totalItems" type="blue" />
      <StatCard title="Baixo Estoque" :value="store.lowStockCount" type="red" />
      <StatCard title="Valor do Inventário" :value="`R$ ${store.totalValue.toFixed(2)}`" type="green" />
    </div>

    <!-- Recent Activity -->
    <div class="bg-white shadow rounded-lg p-6">
      <h3 class="text-lg font-semibold text-gray-800 mb-4 border-b pb-2">Movimentações Recentes</h3>

      <div v-if="store.movements.length > 0" class="space-y-4">
        <div v-for="m in store.movements.slice(0, 5)" :key="m.id" class="flex justify-between items-center p-3 hover:bg-gray-50 rounded">
          <div class="flex items-center">
            <span :class="m.type === 'IN' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 rounded text-xs font-bold mr-3">
              {{ m.type }}
            </span>
            <div>
              <p class="text-sm font-medium text-gray-900">Produto #{{ m.product_id }}</p>
              <p class="text-xs text-gray-500">{{ m.note || 'Sem observações' }}</p>
            </div>
          </div>
          <div class="text-right">
            <p class="text-sm font-bold" :class="m.type === 'IN' ? 'text-green-600' : 'text-red-600'">
              {{ m.type === 'IN' ? '+' : '-' }}{{ m.quantity }}
            </p>
            <p class="text-xs text-gray-400">{{ new Date(m.created_at).toLocaleDateString() }}</p>
          </div>
        </div>
      </div>
      <p v-else class="text-gray-500 text-center py-4">Nenhuma movimentação registrada.</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useInventoryStore } from '../stores/inventory'
import StatCard from '../components/dashboard/StatCard.vue'

const store = useInventoryStore()

onMounted(async () => {
  await store.fetchProducts()
  await store.fetchRecentMovements()
})
</script>
