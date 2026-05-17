<template>
  <div class="min-h-screen bg-gray-100">
    <nav class="bg-blue-800 text-white shadow-lg">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <span class="font-bold text-xl tracking-tight">📦 EstoqueMaster</span>
            <div class="ml-10 flex items-baseline space-x-4">
              <router-link to="/" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-blue-700" active-class="bg-blue-900" exact-active-class="bg-blue-900">Dashboard</router-link>
              <router-link to="/products" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-blue-700" active-class="bg-blue-900">Produtos</router-link>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <main>
      <router-view />
    </main>
  </div>
</template>
