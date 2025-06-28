import ItemContainer from './ItemContainer.vue';
import ItemCard from '../ItemCard/ItemCard.vue';

export default {
  title: 'Common/ItemContainer',
  component: ItemContainer,
  argTypes: {
    grid: { control: 'boolean' },
    customClass: { control: 'text' },
  },
  parameters: {
    docs: {
      description: {
        component: 'A container component that can display items in either a grid or list layout.'
      }
    }
  }
};

const Template = (args) => ({
  components: { ItemContainer, ItemCard },
  setup() { return { args }; },
  template: `
    <ItemContainer v-bind="args">
      <ItemCard v-for="i in 3" :key="i">
        <template #header>
          <h3 style="margin:0; color: var(--color-heading);">Item {{ i }}</h3>
        </template>
        <template #meta>
          <div style="color: var(--color-text-muted); font-size: 0.875rem;">Description for item {{ i }}</div>
        </template>
        <template #actions>
          <button style="background: var(--color-primary-600); color: white; border: none; padding: 0.5rem 1rem; border-radius: 4px; cursor: pointer;">Action</button>
        </template>
      </ItemCard>
    </ItemContainer>
  `,
});

export const Grid = Template.bind({});
Grid.args = { grid: true };

export const List = Template.bind({});
List.args = { grid: false };

export const CompaniesGrid = (args) => ({
  components: { ItemContainer, ItemCard },
  setup() { return { args }; },
  template: `
    <ItemContainer v-bind="args">
      <ItemCard v-for="company in companies" :key="company.id">
        <template #header>
          <div style="display: flex; align-items: flex-start; gap: 0.75rem;">
            <div style="width: 40px; height: 48px; background: var(--color-primary-50); border-radius: 8px; display: flex; align-items: center; justify-content: center; color: var(--color-primary-600);">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                <polyline points="9,22 9,12 15,12 15,22"></polyline>
              </svg>
            </div>
            <div style="flex: 1;">
              <h3 style="margin: 0 0 0.25rem 0; font-size: 1rem; font-weight: 600; color: var(--color-heading);">{{ company.name }}</h3>
            </div>
          </div>
        </template>
        <template #meta>
          <div style="color: var(--color-text-muted); font-size: 0.875rem;">{{ company.location }}</div>
        </template>
        <template #actions>
          <div style="display: flex; justify-content: flex-end; align-items: center; padding-top: 0.75rem; border-top: 1px solid var(--color-border);">
            <span style="font-size: 0.875rem; color: var(--color-primary-600); font-weight: 500;">View Details →</span>
          </div>
        </template>
      </ItemCard>
    </ItemContainer>
  `,
  data() {
    return {
      companies: [
        { id: 1, name: 'TechCorp', location: 'Spokane, WA' },
        { id: 2, name: 'InnovateLab', location: 'Spokane Valley, WA' },
        { id: 3, name: 'Digital Solutions', location: 'Coeur d\'Alene, ID' },
        { id: 4, name: 'StartupHub', location: 'Spokane, WA' },
        { id: 5, name: 'CodeWorks', location: 'Post Falls, ID' },
        { id: 6, name: 'TechStart', location: 'Spokane, WA' }
      ]
    };
  }
});

CompaniesGrid.args = { grid: true };

export const JobsList = (args) => ({
  components: { ItemContainer, ItemCard },
  setup() { return { args }; },
  template: `
    <ItemContainer v-bind="args">
      <ItemCard v-for="job in jobs" :key="job.id">
        <template #header>
          <div style="margin-bottom: 0.75rem;">
            <h3 style="margin: 0 0 0.25rem 0; font-size: 1rem; font-weight: 600; color: var(--color-heading);">{{ job.title }}</h3>
            <div style="font-weight: 500; color: var(--color-text);">{{ job.company }}</div>
          </div>
        </template>
        <template #meta>
          <div style="display: flex; align-items: center; gap: 0.5rem; color: var(--color-text-muted); font-size: 0.875rem;">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
              <circle cx="12" cy="10" r="3"></circle>
            </svg>
            {{ job.location }}
          </div>
        </template>
        <template #actions>
          <div style="display: flex; justify-content: flex-end; align-items: center; padding-top: 0.75rem; border-top: 1px solid var(--color-border);">
            <span style="font-size: 0.875rem; color: var(--color-primary-600); font-weight: 500;">View Details →</span>
          </div>
        </template>
      </ItemCard>
    </ItemContainer>
  `,
  data() {
    return {
      jobs: [
        { id: 1, title: 'Senior Developer', company: 'TechCorp', location: 'Spokane, WA' },
        { id: 2, title: 'Frontend Engineer', company: 'InnovateLab', location: 'Spokane Valley, WA' },
        { id: 3, title: 'Full Stack Developer', company: 'Digital Solutions', location: 'Coeur d\'Alene, ID' }
      ]
    };
  }
});

JobsList.args = { grid: false };
