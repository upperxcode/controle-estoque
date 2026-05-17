<template>
  <div class="bg-white p-6 rounded-lg shadow-md border-l-4" :class="colorClass">
    <div class="flex items-center">
      <div class="flex-1">
        <h3 class="text-sm font-medium text-gray-500 uppercase">{{ title }}</h3>
        <p class="text-2xl font-bold text-gray-900 mt-1">{{ value }}</p>
      </div>
      <div class="text-3xl opacity-20">
        <slot name="icon"></slot>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  title: String,
  value: [String, Number],
  type: {
    type: String,
    default: 'blue'
  }
})

const colorClass = computed(() => {
  const colors = {
    blue: 'border-blue-600',
    red: 'border-red-600',
    green: 'border-green-600',
    yellow: 'border-yellow-600'
  }
  return colors[props.type] || colors.blue
})
</script>
