import AboutIcon from './About.vue';

export default {
  title: 'Icons/About',
  component: AboutIcon,
};

const Template = (args) => ({
  components: { AboutIcon },
  setup() { return { args }; },
  template: '<AboutIcon v-bind="args"></AboutIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
