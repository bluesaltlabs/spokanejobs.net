import ContactIcon from './Contact.vue';

export default {
  title: 'Icons/Contact',
  component: ContactIcon,
};

const Template = (args) => ({
  components: { ContactIcon },
  setup() { return { args }; },
  template: '<ContactIcon v-bind="args"></ContactIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
