import MoonIcon from './Moon.vue';

export default {
  title: 'Icons/Moon',
  component: MoonIcon,
};

const Template = (args) => ({
  components: { MoonIcon },
  setup() { return { args }; },
  template: '<MoonIcon  v-bind="args"></MoonIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
