import Modal from './Modal.vue'
import { ref } from 'vue'
import Button from '../Button/Button.vue'

export default {
  title: 'UI/Modal',
  component: Modal,
  parameters: {
    docs: {
      description: {
        component: 'A reusable modal dialog component with backdrop, animations, and accessibility features.'
      }
    }
  }
}

// Basic modal template
const Template = (args) => ({
  components: { Modal, Button },
  setup() {
    const isOpen = ref(false)
    return { isOpen, args }
  },
  template: `
    <div>
      <Button @click="isOpen = true" variant="primary">Open Modal</Button>
      <Modal v-model="isOpen">
        <template #title>
          {{ args.title || 'Modal Title' }}
        </template>

        <div style="padding: 1rem 0;">
          <p>This is the modal content. You can put any content here.</p>
          <p>The modal includes:</p>
          <ul>
            <li>Backdrop overlay with click-to-close</li>
            <li>Escape key to close</li>
            <li>Body scroll prevention</li>
            <li>Smooth animations</li>
            <li>Accessibility features</li>
          </ul>
        </div>

        <template #footer>
          <Button @click="isOpen = false" variant="primary">OK</Button>
          <Button @click="isOpen = false" variant="secondary">Cancel</Button>
        </template>
      </Modal>
    </div>
  `
})

export const Default = Template.bind({})
Default.args = {
  title: 'Basic Modal'
}
Default.parameters = {
  docs: {
    description: {
      story: 'A basic modal with title, content, and footer actions.'
    }
  }
}

// Template for modal without footer
const WithoutFooterTemplate = (args) => ({
  components: { Modal, Button },
  setup() {
    const isOpen = ref(false)
    return { isOpen, args }
  },
  template: `
    <div>
      <Button @click="isOpen = true" variant="primary">Open Modal</Button>
      <Modal v-model="isOpen">
        <template #title>
          {{ args.title }}
        </template>

        <div style="padding: 1rem 0;">
          <p>This modal has no footer actions. Users can close it by:</p>
          <ul>
            <li>Clicking the X button</li>
            <li>Clicking the backdrop</li>
            <li>Pressing the Escape key</li>
          </ul>
        </div>
      </Modal>
    </div>
  `
})

export const WithoutFooter = WithoutFooterTemplate.bind({})
WithoutFooter.args = {
  title: 'Modal Without Footer'
}
WithoutFooter.parameters = {
  docs: {
    description: {
      story: 'A modal without a footer section. The footer slot is optional.'
    }
  }
}
