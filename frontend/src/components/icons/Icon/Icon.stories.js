import Icon from './Icon.vue'
import iconMap from '@/config/iconMap.js'

export default {
  title: 'Icons/_Icon',
  component: Icon,
  argTypes: {
    slug: {
      control: { type: 'select' },
      options: Object.keys(iconMap)
    },
    size: {
      control: { type: 'text' }
    },
    color: {
      control: { type: 'color' }
    },
    class: {
      control: { type: 'text' }
    }
  }
}

export const Default = {
  args: {
    slug: 'Home',
    size: '20',
    color: 'currentColor'
  }
}

export const Large = {
  args: {
    slug: 'Home',
    size: '32',
    color: '#3b82f6'
  }
}

export const AllIcons = {
  render: () => ({
    components: { Icon },
    setup() {
      return {
        iconSlugs: Object.keys(iconMap)
      }
    },
    template: `
      <div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(120px, 1fr)); gap: 16px; padding: 20px;">
        <div v-for="slug in iconSlugs" :key="slug" style="text-align: center; padding: 16px; border: 1px solid #e5e7eb; border-radius: 8px;">
          <Icon :slug="slug" size="24" />
          <div style="margin-top: 8px; font-size: 12px; color: #6b7280;">{{ slug }}</div>
        </div>
      </div>
    `
  })
}

export const CustomColors = {
  render: () => ({
    components: { Icon },
    template: `
      <div style="display: flex; gap: 16px; align-items: center; padding: 20px;">
        <Icon slug="Home" size="24" color="#ef4444" />
        <Icon slug="Jobs" size="24" color="#10b981" />
        <Icon slug="Profile" size="24" color="#3b82f6" />
        <Icon slug="Moon" size="24" color="#8b5cf6" />
        <Icon slug="Sun" size="24" color="#f59e0b" />
      </div>
    `
  })
}
