import { setActivePinia, createPinia } from "pinia";
import { describe, it, expect, beforeEach, vi } from "vitest";
import { useInventoryStore } from "./inventory";
import movementService from "../services/movementService";

vi.mock("../services/movementService", () => ({
  default: {
    register: vi.fn(() => Promise.resolve({ data: {} })),
  },
}));

describe("Inventory Store", () => {
  beforeEach(() => {
    setActivePinia(createPinia());
    vi.clearAllMocks();
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

  it("should update product quantity locally after registerMovement (IN)", async () => {
    const store = useInventoryStore();
    store.products = [{ id: 1, quantity: 10 }];

    await store.registerMovement({ product_id: 1, quantity: 5, type: "IN" });

    expect(movementService.register).toHaveBeenCalled();
    expect(store.products[0].quantity).toBe(15);
  });

  it("should update product quantity locally after registerMovement (OUT)", async () => {
    const store = useInventoryStore();
    store.products = [{ id: 1, quantity: 10 }];

    await store.registerMovement({ product_id: 1, quantity: 3, type: "OUT" });

    expect(movementService.register).toHaveBeenCalled();
    expect(store.products[0].quantity).toBe(7);
  });
});
