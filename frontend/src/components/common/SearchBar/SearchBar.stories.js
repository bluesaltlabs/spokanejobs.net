import SearchBar from './SearchBar.vue';

export default {
  title: 'Common/SearchBar',
  component: SearchBar,
  argTypes: {
    modelValue: { control: 'text' },
    placeholder: { control: 'text' },
  },
};

const Template = (args) => ({
  components: { SearchBar },
  setup() {
    return { args };
  },
  template: '<SearchBar v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {
  modelValue: '',
  placeholder: 'Search...'
};

export const WithValue = Template.bind({});
WithValue.args = {
  modelValue: 'Acme',
  placeholder: 'Search companies...'
};
