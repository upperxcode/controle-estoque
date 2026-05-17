<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold text-gray-900">Gestão de Estoque</h2>
        <button v-if="!showForm" @click="openCreateForm" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 font-medium">
          Adicionar Produto
        </button>
      </div>

      <div v-if="showForm" class="mb-8 max-w-2xl mx-auto">
        <ProductForm
          :initialData="editingProduct"
          :loading="store.loading"
          @submit="handleSave"
          @cancel="closeForm"
        />
      </div>

      <DataTable
        :headers="headers"
        :items="store.products"
        :loading="store.loading"
      >
        <template #cell-price="{ item }">
          R$ {{ item.price.toFixed(2) }}
        </template>
        <template #actions="{ item }">
          <button @click="openEditForm(item)" class="text-blue-600 hover:text-blue-900 mr-4 font-medium">Editar</button>
          <button @click="handleDelete(item.id)" class="text-red-600 hover:text-red-900 font-medium">Deletar</button>
        </template>
      </DataTable>

      <div v-if="store.error" class="mt-4 p-4 bg-red-100 text-red-700 rounded-md">
        {{ store.error }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useInventoryStore } from '../stores/inventory'
import DataTable from '../components/common/DataTable.vue'
import ProductForm from '../components/products/ProductForm.vue'

const store = useInventoryStore()
const showForm = ref(false)
const editingProduct = ref(null)

const headers = [
  { label: 'Nome', key: 'name' },
  { label: 'SKU', key: 'sku' },
  { label: 'Qtd', key: 'quantity' },
  { label: 'Preço', key: 'price' }
]

onMounted(() => {
  store.fetchProducts()
})

const openCreateForm = () => {
  editingProduct.ref = null
  showForm.value = true
}

const openEditForm = (product) => {
  editingProduct.value = { ...product }
  showForm.value = true
}

const closeForm = () => {
  showForm.value = false
  editingProduct.value = null
}

const handleSave = async (payload) => {
  try {
    if (editingProduct.value) {
      await store.updateProduct(editingProduct.value.id, payload)
    } else {
      await store.addProduct(payload)
    }
    closeForm()
  } catch (err) {
    // Error handled by store
  }
}

const handleDelete = async (id) => {
  if (confirm('Tem certeza que deseja remover este produto?')) {
    try {
      await store.deleteProduct(id)
    } catch (err) {
      // Error handled by store
    }
  }
}
</script>
