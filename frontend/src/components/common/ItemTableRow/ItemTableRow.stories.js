import ItemTableRow from './ItemTableRow.vue';

export default {
  title: 'Common/ItemTableRow',
  component: ItemTableRow,
  parameters: {
    docs: {
      description: {
        component: 'A table row component for displaying items in a structured table format with hover effects.'
      }
    }
  }
};

const Template = (args) => ({
  components: { ItemTableRow },
  setup() { return { args }; },
  template: `
    <table style="width:100%; border-collapse: collapse; background: var(--color-surface); border-radius: 8px; overflow: hidden; box-shadow: 0 1px 3px var(--color-shadow);">
      <thead>
        <tr style="background: var(--color-background-soft);">
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Name</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Type</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Status</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Actions</th>
        </tr>
      </thead>
      <tbody>
        <ItemTableRow v-bind="args">
          <template #cells>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">Example Item</td>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">Type A</td>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">Active</td>
          </template>
          <template #actions>
            <td style="padding: 1rem 1.5rem;">
              <button style="background: var(--color-primary-600); color: white; border: none; padding: 0.5rem 1rem; border-radius: 4px; cursor: pointer; font-size: 0.875rem;">View</button>
            </td>
          </template>
        </ItemTableRow>
      </tbody>
    </table>
  `,
});

export const Default = Template.bind({});
Default.args = {};

export const CompaniesTable = (args) => ({
  components: { ItemTableRow },
  setup() { return { args }; },
  template: `
    <table style="width:100%; border-collapse: collapse; background: var(--color-surface); border-radius: 8px; overflow: hidden; box-shadow: 0 1px 3px var(--color-shadow);">
      <thead>
        <tr style="background: var(--color-background-soft);">
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Name</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Location</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Website</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Actions</th>
        </tr>
      </thead>
      <tbody>
        <ItemTableRow v-for="company in companies" :key="company.id" v-bind="args">
          <template #cells>
            <td style="padding: 1rem 1.5rem; color: var(--color-text); font-weight: 600;">{{ company.name }}</td>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">{{ company.location }}</td>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">
              <a href="#" style="color: var(--color-link); text-decoration: none;">{{ company.website }}</a>
            </td>
          </template>
          <template #actions>
            <td style="padding: 1rem 1.5rem;">
              <button style="background: var(--color-primary-600); color: white; border: none; padding: 0.5rem 1rem; border-radius: 4px; cursor: pointer; font-size: 0.875rem;">View Details</button>
            </td>
          </template>
        </ItemTableRow>
      </tbody>
    </table>
  `,
  data() {
    return {
      companies: [
        { id: 1, name: 'TechCorp', location: 'Spokane, WA', website: 'techcorp.com' },
        { id: 2, name: 'InnovateLab', location: 'Spokane Valley, WA', website: 'innovatelab.com' },
        { id: 3, name: 'Digital Solutions', location: 'Coeur d\'Alene, ID', website: 'digitalsolutions.com' }
      ]
    };
  }
});

export const JobsTable = (args) => ({
  components: { ItemTableRow },
  setup() { return { args }; },
  template: `
    <table style="width:100%; border-collapse: collapse; background: var(--color-surface); border-radius: 8px; overflow: hidden; box-shadow: 0 1px 3px var(--color-shadow);">
      <thead>
        <tr style="background: var(--color-background-soft);">
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Title</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Company</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Location</th>
          <th style="padding: 1rem 1.5rem; text-align: left; font-weight: 600; color: var(--color-heading); font-size: 0.9rem; text-transform: uppercase; letter-spacing: 0.5px;">Actions</th>
        </tr>
      </thead>
      <tbody>
        <ItemTableRow v-for="job in jobs" :key="job.id" v-bind="args">
          <template #cells>
            <td style="padding: 1rem 1.5rem; color: var(--color-text); font-weight: 600;">{{ job.title }}</td>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">{{ job.company }}</td>
            <td style="padding: 1rem 1.5rem; color: var(--color-text);">{{ job.location }}</td>
          </template>
          <template #actions>
            <td style="padding: 1rem 1.5rem;">
              <button style="background: var(--color-primary-600); color: white; border: none; padding: 0.5rem 1rem; border-radius: 4px; cursor: pointer; font-size: 0.875rem;">Apply</button>
            </td>
          </template>
        </ItemTableRow>
      </tbody>
    </table>
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
