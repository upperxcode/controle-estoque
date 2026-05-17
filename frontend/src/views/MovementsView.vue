<template>
  <div class="max-w-4xl mx-auto py-8 px-4">
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Movimentação de Estoque</h1>
      <p class="text-gray-600">Registre entradas ou saídas de produtos no inventário.</p>
    </div>

    <div v-if="store.loading && !store.products.length" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
      <p class="mt-4 text-gray-500">Carregando produtos...</p>
    </div>

    <div v-else>
      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="flex border-b">
          <router-link
            to="/movements/in"
            class="flex-1 py-4 text-center font-medium hover:bg-gray-50"
            :class="[movementType === 'IN' ? 'text-blue-600 border-b-2 border-blue-600' : 'text-gray-500']"
          >
            Entrada
          </router-link>
          <router-link
            to="/movements/out"
            class="flex-1 py-4 text-center font-medium hover:bg-gray-50"
            :class="[movementType === 'OUT' ? 'text-blue-600 border-b-2 border-blue-600' : 'text-gray-500']"
          >
            Saída
          </router-link>
        </div>

        <div class="p-6">
          <div v-if="successMessage" class="mb-4 p-4 bg-green-100 text-green-700 rounded-md">
            {{ successMessage }}
          </div>
          <div v-if="store.error" class="mb-4 p-4 bg-red-100 text-red-700 rounded-md">
            {{ store.error }}
          </div>

          <MovementForm
            :key="movementType"
            :type="movementType"
            :products="store.products"
            :loading="store.loading"
            @submit="handleMovement"
            @cancel="goBack"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useInventoryStore } from '../stores/inventory'
import MovementForm from '../components/movements/MovementForm.vue'

const route = useRoute()
const router = useRouter()
const store = useInventoryStore()

const successMessage = ref('')

const movementType = computed(() => {
  return route.params.type === 'out' ? 'OUT' : 'IN'
})

onMounted(async () => {
  if (store.products.length === 0) {
    await store.fetchProducts()
  }
})

const handleMovement = async (payload) => {
  try {
    await store.registerMovement(payload)
    successMessage.value = `Movimentação de ${payload.type === 'IN' ? 'entrada' : 'saída'} registrada com sucesso!`

    // Limpar mensagem após 3 segundos
    setTimeout(() => {
      successMessage.value = ''
    }, 3000)
  } catch (err) {
    // Erro já é tratado na store
  }
}

const goBack = () => {
  router.push('/products')
}
</script>
