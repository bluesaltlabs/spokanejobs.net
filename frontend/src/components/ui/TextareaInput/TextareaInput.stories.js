import TextareaInput from './TextareaInput.vue';

export default {
  title: 'UI/TextareaInput',
  component: TextareaInput,
  argTypes: {
    modelValue: { control: 'text' },
    placeholder: { control: 'text' },
    rows: { control: 'number' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { TextareaInput },
  setup() { return { args }; },
  template: '<TextareaInput v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { placeholder: 'Enter text...', rows: 3 };

export const Disabled = Template.bind({});
Disabled.args = { disabled: true, placeholder: 'Disabled textarea' };
