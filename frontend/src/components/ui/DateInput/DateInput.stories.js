import DateInput from './DateInput.vue';

export default {
  title: 'UI/DateInput',
  component: DateInput,
  argTypes: {
    modelValue: { control: 'text' },
    withTime: { control: 'boolean' },
    min: { control: 'text' },
    max: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { DateInput },
  setup() { return { args }; },
  template: '<DateInput v-bind="args" />',
});

export const DateOnly = Template.bind({});
DateOnly.args = { modelValue: '', withTime: false };

export const DateTime = Template.bind({});
DateTime.args = { modelValue: '', withTime: true };

export const Disabled = Template.bind({});
Disabled.args = { modelValue: '', disabled: true };
