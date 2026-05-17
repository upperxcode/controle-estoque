<template>
  <form @submit.prevent="handleSubmit" class="space-y-4 bg-gray-50 p-6 rounded-lg border">
    <h3 class="text-lg font-semibold text-gray-700">
      {{ initialData ? 'Editar Produto' : 'Novo Produto' }}
    </h3>

    <div>
      <label class="block text-sm font-medium text-gray-700">Nome</label>
      <input v-model="form.name" type="text" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border" />
    </div>

    <div class="grid grid-cols-2 gap-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">SKU</label>
        <input v-model="form.sku" type="text" required class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Preço</label>
        <input v-model.number="form.price" type="number" step="0.01" required min="0.01" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border" />
      </div>
    </div>

    <div>
      <label class="block text-sm font-medium text-gray-700">Descrição</label>
      <textarea v-model="form.description" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm p-2 border"></textarea>
    </div>

    <div class="flex justify-end space-x-3 mt-6">
      <button type="button" @click="$emit('cancel')" class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
        Cancelar
      </button>
      <button type="submit" :disabled="loading" class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50">
        {{ loading ? 'Salvando...' : 'Salvar' }}
      </button>
    </div>
  </form>
</template>

<script setup>
import { reactive, onMounted } from 'vue'

const props = defineProps({
  initialData: {
    type: Object,
    default: null
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit', 'cancel'])

const form = reactive({
  name: '',
  description: '',
  sku: '',
  price: 0,
  quantity: 0
})

onMounted(() => {
  if (props.initialData) {
    Object.assign(form, props.initialData)
  }
})

const handleSubmit = () => {
  emit('submit', { ...form })
}
</script>
