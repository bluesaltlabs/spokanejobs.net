import JobsIcon from './Jobs.vue';

export default {
  title: 'Icons/Jobs',
  component: JobsIcon,
};

const Template = (args) => ({
  components: { JobsIcon },
  setup() { return { args }; },
  template: '<JobsIcon width="40" height="40" v-bind="args"></JobsIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
