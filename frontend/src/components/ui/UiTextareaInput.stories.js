import UiTextareaInput from './UiTextareaInput.vue';

export default {
  title: 'UI/UiTextareaInput',
  component: UiTextareaInput,
  argTypes: {
    modelValue: { control: 'text' },
    placeholder: { control: 'text' },
    rows: { control: 'number' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiTextareaInput },
  setup() { return { args }; },
  template: '<UiTextareaInput v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { placeholder: 'Enter text...', rows: 3 };

export const Disabled = Template.bind({});
Disabled.args = { disabled: true, placeholder: 'Disabled textarea' }; 