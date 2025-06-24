import UiDateInput from './UiDateInput.vue';

export default {
  title: 'UI/UiDateInput',
  component: UiDateInput,
  argTypes: {
    modelValue: { control: 'text' },
    withTime: { control: 'boolean' },
    min: { control: 'text' },
    max: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiDateInput },
  setup() { return { args }; },
  template: '<UiDateInput v-bind="args" />',
});

export const DateOnly = Template.bind({});
DateOnly.args = { modelValue: '', withTime: false };

export const DateTime = Template.bind({});
DateTime.args = { modelValue: '', withTime: true };

export const Disabled = Template.bind({});
Disabled.args = { modelValue: '', disabled: true }; 