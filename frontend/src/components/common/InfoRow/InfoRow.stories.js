import InfoRow from './InfoRow.vue';

export default {
  title: 'Common/InfoRow',
  component: InfoRow,
  argTypes: {
    label: { control: 'text' },
    value: { control: 'text' },
    icon: { control: 'text' },
  },
};

const Template = (args) => ({
  components: { InfoRow },
  setup() { return { args }; },
  template: '<InfoRow v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {
  label: 'Last Scraped',
  value: '2025-06-30 18:00',
  icon: 'mdi:clock-outline',
};

export const NoIcon = Template.bind({});
NoIcon.args = {
  label: 'Active Jobs',
  value: '12',
  icon: '',
};
