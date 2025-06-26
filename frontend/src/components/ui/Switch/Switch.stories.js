import Switch from './Switch.vue';

export default {
  title: 'UI/Switch',
  component: Switch,
  argTypes: {
    modelValue: { control: 'boolean' },
    label: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { Switch },
  setup() { return { args }; },
  template: '<Switch v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { label: 'Enable feature', modelValue: false };

export const Checked = Template.bind({});
Checked.args = { label: 'Enable feature', modelValue: true };

export const Disabled = Template.bind({});
Disabled.args = { label: 'Enable feature', modelValue: false, disabled: true };
