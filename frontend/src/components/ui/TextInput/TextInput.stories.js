import TextInput from './TextInput.vue';

export default {
  title: 'UI/TextInput',
  component: TextInput,
  argTypes: {
    modelValue: { control: 'text' },
    placeholder: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { TextInput },
  setup() { return { args }; },
  template: '<TextInput v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { placeholder: 'Enter text...' };

export const Disabled = Template.bind({});
Disabled.args = { disabled: true, placeholder: 'Disabled input' };
