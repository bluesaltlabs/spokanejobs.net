import ViewIcon from './View.vue';

export default {
  title: 'Icons/View',
  component: ViewIcon,
};

const Template = (args) => ({
  components: { ViewIcon },
  setup() { return { args }; },
  template: '<ViewIcon width="40" height="40" v-bind="args"></ViewIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
