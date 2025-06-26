import CompaniesIcon from './Companies.vue';

export default {
  title: 'Icons/Companies',
  component: CompaniesIcon,
};

const Template = (args) => ({
  components: { CompaniesIcon },
  setup() { return { args }; },
  template: '<CompaniesIcon width="40" height="40" v-bind="args"></CompaniesIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
