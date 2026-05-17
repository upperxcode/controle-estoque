import { setActivePinia, createPinia } from 'pinia'
import { describe, it, expect, beforeEach } from 'vitest'
import { useInventoryStore } from './inventory'

describe('Inventory Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  it('should have empty products initially', () => {
    const store = useInventoryStore()
    expect(store.products).toEqual([])
  })

  it('should have loading false initially', () => {
    const store = useInventoryStore()
    expect(store.loading).toBe(false)
  })
})
