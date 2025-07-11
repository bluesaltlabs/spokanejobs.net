import ImportIcon from './Import.vue';

export default {
  title: 'Icons/Import',
  component: ImportIcon,
};

const Template = (args) => ({
  components: { ImportIcon },
  setup() { return { args }; },
  template: '<ImportIcon  v-bind="args"></ImportIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
