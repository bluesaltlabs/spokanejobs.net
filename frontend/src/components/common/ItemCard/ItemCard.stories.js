import ItemCard from './ItemCard.vue';

export default {
  title: 'Common/ItemCard',
  component: ItemCard,
  argTypes: {
    customClass: { control: 'text' },
    onClick: { action: 'clicked' },
  },
  parameters: {
    docs: {
      description: {
        component: 'A versatile card component for displaying items in grids or lists with header, meta, and actions sections.'
      }
    }
  }
};

const Template = (args) => ({
  components: { ItemCard },
  setup() { return { args }; },
  template: `
    <ItemCard v-bind="args">
      <template #header>
        <h3 style="margin:0; color: var(--color-heading);">Header Content</h3>
      </template>
      <template #meta>
        <div style="color: var(--color-text-muted); font-size: 0.875rem;">Meta information here</div>
      </template>
      <template #actions>
        <button style="background: var(--color-primary-600); color: white; border: none; padding: 0.5rem 1rem; border-radius: 4px; cursor: pointer;">Action</button>
      </template>
    </ItemCard>
  `,
});

export const Default = Template.bind({});
Default.args = {};

export const WithCustomClass = Template.bind({});
WithCustomClass.args = {
  customClass: 'custom-card'
};

export const CompanyCard = (args) => ({
  components: { ItemCard },
  setup() { return { args }; },
  template: `
    <ItemCard v-bind="args">
      <template #header>
        <div style="display: flex; align-items: flex-start; gap: 0.75rem;">
          <div style="width: 40px; height: 48px; background: var(--color-primary-50); border-radius: 8px; display: flex; align-items: center; justify-content: center; color: var(--color-primary-600);">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
              <polyline points="9,22 9,12 15,12 15,22"></polyline>
            </svg>
          </div>
          <div style="flex: 1;">
            <h3 style="margin: 0 0 0.25rem 0; font-size: 1rem; font-weight: 600; color: var(--color-heading);">Example Company</h3>
          </div>
        </div>
      </template>
      <template #meta>
        <div style="color: var(--color-text-muted); font-size: 0.875rem;">Spokane, WA</div>
      </template>
      <template #actions>
        <div style="display: flex; justify-content: flex-end; align-items: center; padding-top: 0.75rem; border-top: 1px solid var(--color-border);">
          <span style="font-size: 0.875rem; color: var(--color-primary-600); font-weight: 500;">View Details →</span>
        </div>
      </template>
    </ItemCard>
  `,
});

export const JobCard = (args) => ({
  components: { ItemCard },
  setup() { return { args }; },
  template: `
    <ItemCard v-bind="args">
      <template #header>
        <div style="margin-bottom: 0.75rem;">
          <h3 style="margin: 0 0 0.25rem 0; font-size: 1rem; font-weight: 600; color: var(--color-heading);">Senior Developer</h3>
          <div style="font-weight: 500; color: var(--color-text);">Tech Company Inc.</div>
        </div>
      </template>
      <template #meta>
        <div style="display: flex; align-items: center; gap: 0.5rem; color: var(--color-text-muted); font-size: 0.875rem;">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
            <circle cx="12" cy="10" r="3"></circle>
          </svg>
          Spokane, WA
        </div>
      </template>
      <template #actions>
        <div style="display: flex; justify-content: flex-end; align-items: center; padding-top: 0.75rem; border-top: 1px solid var(--color-border);">
          <span style="font-size: 0.875rem; color: var(--color-primary-600); font-weight: 500;">View Details →</span>
        </div>
      </template>
    </ItemCard>
  `,
});
