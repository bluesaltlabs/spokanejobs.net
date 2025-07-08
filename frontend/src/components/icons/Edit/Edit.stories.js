import EditIcon from './Edit.vue';

export default {
  title: 'Icons/Edit',
  component: EditIcon,
};

const Template = (args) => ({
  components: { EditIcon },
  setup() { return { args }; },
  template: '<EditIcon v-bind="args"></EditIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
