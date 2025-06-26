import ResourcesIcon from './Resources.vue';

export default {
  title: 'Icons/Resources',
  component: ResourcesIcon,
};

const Template = (args) => ({
  components: { ResourcesIcon },
  setup() { return { args }; },
  template: '<ResourcesIcon width="40" height="40" v-bind="args"></ResourcesIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
