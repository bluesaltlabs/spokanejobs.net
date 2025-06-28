import SunIcon from './Sun.vue';

export default {
  title: 'Icons/Sun',
  component: SunIcon,
};

const Template = (args) => ({
  components: { SunIcon },
  setup() { return { args }; },
  template: '<SunIcon  v-bind="args"></SunIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
