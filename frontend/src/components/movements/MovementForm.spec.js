import { mount } from '@vue/test-utils'
import { describe, it, expect } from 'vitest'
import MovementForm from './MovementForm.vue'

describe('MovementForm.vue', () => {
  const products = [
    { id: 1, name: 'Product 1', sku: 'SKU1', quantity: 10, price: 100 },
    { id: 2, name: 'Product 2', sku: 'SKU2', quantity: 5, price: 50 }
  ]

  it('renders correctly for IN type', () => {
    const wrapper = mount(MovementForm, {
      props: { type: 'IN', products: products }
    })
    expect(wrapper.find('h3').text()).toBe('Registrar Entrada')
  })

  it('renders correctly for OUT type', () => {
    const wrapper = mount(MovementForm, {
      props: { type: 'OUT', products: products }
    })
    expect(wrapper.find('h3').text()).toBe('Registrar Saída')
  })

  it('disables submit button when form is invalid', async () => {
    const wrapper = mount(MovementForm, {
      props: { type: 'IN', products: products }
    })
    const submitBtn = wrapper.find('button[type="submit"]')
    expect(submitBtn.element.disabled).toBe(true)

    // Fill product_id
    await wrapper.find('select').setValue(1)
    expect(submitBtn.element.disabled).toBe(false)
  })

  it('validates maximum quantity for OUT type', async () => {
    const wrapper = mount(MovementForm, {
      props: { type: 'OUT', products: products }
    })

    // Select product with quantity 10
    await wrapper.find('select').setValue(1)

    const quantityInput = wrapper.find('input[type="number"]')
    await quantityInput.setValue(11)

    const submitBtn = wrapper.find('button[type="submit"]')
    expect(submitBtn.element.disabled).toBe(true)

    await quantityInput.setValue(5)
    expect(submitBtn.element.disabled).toBe(false)
  })

  it('emits submit event with correct payload', async () => {
    const wrapper = mount(MovementForm, {
      props: { type: 'IN', products: products }
    })

    await wrapper.find('select').setValue(1)
    await wrapper.find('input[type="number"]').setValue(5)
    await wrapper.find('textarea').setValue('Some notes')

    await wrapper.find('form').trigger('submit.prevent')

    expect(wrapper.emitted()).toHaveProperty('submit')
    expect(wrapper.emitted().submit[0][0]).toEqual({
      product_id: 1,
      quantity: 5,
      notes: 'Some notes',
      type: 'IN'
    })
  })
})
