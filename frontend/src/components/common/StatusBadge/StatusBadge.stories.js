import StatusBadge from './StatusBadge.vue';

export default {
  title: 'Common/StatusBadge',
  component: StatusBadge,
  argTypes: {
    status: {
      control: { type: 'select' },
      options: ['built', 'in progress', 'error', 'not built', 'other'],
    },
    label: { control: 'text' },
    icon: { control: 'text' },
  },
};

const Template = (args) => ({
  components: { StatusBadge },
  setup() { return { args }; },
  template: '<StatusBadge v-bind="args" />',
});

export const Built = Template.bind({});
Built.args = {
  status: 'built',
  label: 'Scraper Built',
  icon: 'mdi:check-circle-outline',
};

export const InProgress = Template.bind({});
InProgress.args = {
  status: 'in progress',
  label: 'Scraping In Progress',
  icon: 'mdi:progress-clock',
};

export const Error = Template.bind({});
Error.args = {
  status: 'error',
  label: 'Scraper Error',
  icon: 'mdi:alert-circle-outline',
};

export const NotBuilt = Template.bind({});
NotBuilt.args = {
  status: 'not built',
  label: 'No Scraper',
  icon: 'mdi:close-circle-outline',
};

export const Other = Template.bind({});
Other.args = {
  status: 'other',
  label: 'Unknown',
  icon: '',
};
