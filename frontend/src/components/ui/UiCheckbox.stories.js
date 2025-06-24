import UiCheckbox from './UiCheckbox.vue';

export default {
  title: 'UI/UiCheckbox',
  component: UiCheckbox,
  argTypes: {
    modelValue: { control: 'boolean' },
    label: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiCheckbox },
  setup() { return { args }; },
  template: '<UiCheckbox v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { label: 'Accept terms', modelValue: false };

export const Checked = Template.bind({});
Checked.args = { label: 'Accept terms', modelValue: true };

export const Disabled = Template.bind({});
Disabled.args = { label: 'Accept terms', modelValue: false, disabled: true }; 