import Checkbox from './Checkbox.vue';

export default {
  title: 'UI/Checkbox',
  component: Checkbox,
  argTypes: {
    modelValue: { control: 'boolean' },
    label: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { Checkbox },
  setup() { return { args }; },
  template: '<Checkbox v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { label: 'Accept terms', modelValue: false };

export const Checked = Template.bind({});
Checked.args = { label: 'Accept terms', modelValue: true };

export const Disabled = Template.bind({});
Disabled.args = { label: 'Accept terms', modelValue: false, disabled: true };
