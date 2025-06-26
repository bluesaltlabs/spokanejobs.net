import ExportIcon from './Export.vue';

export default {
  title: 'Icons/Export',
  component: ExportIcon,
};

const Template = (args) => ({
  components: { ExportIcon },
  setup() { return { args }; },
  template: '<ExportIcon width="40" height="40" v-bind="args"></ExportIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
