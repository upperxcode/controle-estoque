<template>
  <form @submit.prevent="handleSubmit" class="space-y-4 bg-gray-50 p-6 rounded-lg border">
    <h3 class="text-lg font-semibold text-gray-700">
      {{ type === 'IN' ? 'Registrar Entrada' : 'Registrar Saída' }}
    </h3>

    <div>
      <label class="block text-sm font-medium text-gray-700">Produto</label>
      <select v-model="form.product_id" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border">
        <option value="" disabled>Selecione um produto</option>
        <option v-for="product in products" :key="product.id" :value="product.id">
          {{ product.name }} (SKU: {{ product.sku }}) - Saldo: {{ product.quantity }}
        </option>
      </select>
    </div>

    <div>
      <label class="block text-sm font-medium text-gray-700">Quantidade</label>
      <input v-model.number="form.quantity" type="number" required min="1" :max="maxQuantity" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border" />
      <p v-if="type === 'OUT' && selectedProduct" class="text-xs text-gray-500 mt-1">
        Saldo disponível: {{ selectedProduct.quantity }}
      </p>
    </div>

    <div>
      <label class="block text-sm font-medium text-gray-700">Notas / Observações</label>
      <textarea v-model="form.notes" rows="3" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border"></textarea>
    </div>

    <div class="flex justify-end space-x-3 mt-6">
      <button type="button" @click="$emit('cancel')" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
        Cancelar
      </button>
      <button type="submit" :disabled="loading || !isValid" class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50">
        {{ loading ? 'Processando...' : 'Confirmar' }}
      </button>
    </div>
  </form>
</template>

<script setup>
import { reactive, computed } from 'vue'

const props = defineProps({
  type: {
    type: String,
    required: true,
    validator: (value) => ['IN', 'OUT'].includes(value)
  },
  products: {
    type: Array,
    required: true
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit', 'cancel'])

const form = reactive({
  product_id: '',
  quantity: 1,
  notes: '',
  type: props.type
})

const selectedProduct = computed(() => {
  return props.products.find(p => p.id === form.product_id)
})

const maxQuantity = computed(() => {
  if (props.type === 'OUT' && selectedProduct.value) {
    return selectedProduct.value.quantity
  }
  return undefined
})

const isValid = computed(() => {
  if (!form.product_id || form.quantity <= 0) return false
  if (props.type === 'OUT' && selectedProduct.value && form.quantity > selectedProduct.value.quantity) return false
  return true
})

const handleSubmit = () => {
  if (isValid.value) {
    emit('submit', { ...form, type: props.type })
  }
}
</script>
