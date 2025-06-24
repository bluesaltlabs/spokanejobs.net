import UiRadio from './UiRadio.vue';

export default {
  title: 'UI/UiRadio',
  component: UiRadio,
  argTypes: {
    modelValue: { control: 'boolean' },
    label: { control: 'text' },
    name: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiRadio },
  setup() { return { args }; },
  template: '<UiRadio v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { label: 'Option A', name: 'group1', modelValue: false };

export const Checked = Template.bind({});
Checked.args = { label: 'Option A', name: 'group1', modelValue: true };

export const Disabled = Template.bind({});
Disabled.args = { label: 'Option A', name: 'group1', modelValue: false, disabled: true }; 