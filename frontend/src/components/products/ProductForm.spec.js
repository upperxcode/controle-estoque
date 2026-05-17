import { mount } from '@vue/test-utils'
import { describe, it, expect } from 'vitest'
import ProductForm from './ProductForm.vue'

describe('ProductForm.vue', () => {
  it('should render correct title for new product', () => {
    const wrapper = mount(ProductForm, {
      props: { initialData: null }
    })
    expect(wrapper.find('h3').text()).toBe('Novo Produto')
  })

  it('should render correct title for editing product', () => {
    const wrapper = mount(ProductForm, {
      props: { initialData: { name: 'Test' } }
    })
    expect(wrapper.find('h3').text()).toBe('Editar Produto')
  })

  it('emits cancel event when cancel button is clicked', async () => {
    const wrapper = mount(ProductForm)
    await wrapper.find('button[type="button"]').trigger('click')
    expect(wrapper.emitted()).toHaveProperty('cancel')
  })
})
