import DarkModeToggle from './DarkModeToggle.vue';

export default {
  title: 'Common/DarkModeToggle',
  component: DarkModeToggle,
  argTypes: {},
};

const Template = (args) => ({
  components: { DarkModeToggle },
  setup() {
    return { args };
  },
  template: '<DarkModeToggle v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {};
