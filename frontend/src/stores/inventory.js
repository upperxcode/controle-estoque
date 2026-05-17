import { defineStore } from "pinia";
import productService from "../services/productService";
import movementService from "../services/movementService";

export const useInventoryStore = defineStore("inventory", {
  state: () => ({
    products: [],
    loading: false,
    error: null,
  }),
  actions: {
    async fetchProducts() {
      this.loading = true;
      try {
        const response = await productService.getAll();
        this.products = response.data;
      } catch (err) {
        this.error = "Erro ao carregar produtos";
        console.error(err);
      } finally {
        this.loading = false;
      }
    },
    async addProduct(product) {
      this.loading = true;
      try {
        const response = await productService.create(product);
        this.products.unshift(response.data);
        return response.data;
      } catch (err) {
        this.error = "Erro ao adicionar produto";
        throw err;
      } finally {
        this.loading = false;
      }
    },
    async updateProduct(id, product) {
      this.loading = true;
      try {
        const response = await productService.update(id, product);
        const index = this.products.findIndex((p) => p.id === id);
        if (index !== -1) {
          this.products[index] = response.data;
        }
        return response.data;
      } catch (err) {
        this.error = "Erro ao atualizar produto";
        throw err;
      } finally {
        this.loading = false;
      }
    },
    async deleteProduct(id) {
      this.loading = true;
      try {
        await productService.delete(id);
        this.products = this.products.filter((p) => p.id !== id);
      } catch (err) {
        this.error = "Erro ao deletar produto";
        throw err;
      } finally {
        this.loading = false;
      }
    },
    async registerMovement(movement) {
      this.loading = true;
      try {
        await movementService.register(movement);
        // Atualizar saldo localmente
        const product = this.products.find((p) => p.id === movement.product_id);
        if (product) {
          if (movement.type === "IN") {
            product.quantity += movement.quantity;
          } else {
            product.quantity -= movement.quantity;
          }
        }
      } catch (err) {
        this.error =
          err.response?.data?.error || "Erro ao registrar movimentação";
        throw err;
      } finally {
        this.loading = false;
      }
    },
  },
});
