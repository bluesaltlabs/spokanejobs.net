import CompaniesIcon from './Companies.vue';

export default {
  title: 'Icons/Companies',
  component: CompaniesIcon,
};

const Template = (args) => ({
  components: { CompaniesIcon },
  setup() { return { args }; },
  template: '<CompaniesIcon v-bind="args"></CompaniesIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
