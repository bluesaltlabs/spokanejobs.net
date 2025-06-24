import UiTextInput from './UiTextInput.vue';

export default {
  title: 'UI/UiTextInput',
  component: UiTextInput,
  argTypes: {
    modelValue: { control: 'text' },
    placeholder: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiTextInput },
  setup() { return { args }; },
  template: '<UiTextInput v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = { placeholder: 'Enter text...' };

export const Disabled = Template.bind({});
Disabled.args = { disabled: true, placeholder: 'Disabled input' }; 