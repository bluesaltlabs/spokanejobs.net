import HomeIcon from './Home.vue';

export default {
  title: 'Icons/Home',
  component: HomeIcon,
};

const Template = (args) => ({
  components: { HomeIcon },
  setup() { return { args }; },
  template: '<HomeIcon v-bind="args"></HomeIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
