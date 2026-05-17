import { setActivePinia, createPinia } from "pinia";
import { describe, it, expect, beforeEach } from "vitest";
import { useInventoryStore } from "./inventory";

describe("Inventory Store", () => {
  beforeEach(() => {
    setActivePinia(createPinia());
  });

  it("should have empty products initially", () => {
    const store = useInventoryStore();
    expect(store.products).toEqual([]);
  });

  it("should have loading false initially", () => {
    const store = useInventoryStore();
    expect(store.loading).toBe(false);
  });

  it("should calculate getters correctly", () => {
    const store = useInventoryStore();
    store.products = [
      { id: 1, quantity: 10, price: 100 },
      { id: 2, quantity: 2, price: 50 },
      { id: 3, quantity: 5, price: 10 },
    ];

    expect(store.totalItems).toBe(17);
    expect(store.lowStockCount).toBe(1); // Only id 2 has < 5
    expect(store.totalValue).toBe(1150); // (10*100) + (2*50) + (5*10) = 1000 + 100 + 50
  });
});
