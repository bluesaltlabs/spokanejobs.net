import UiFormGroup from './UiFormGroup.vue';
import UiTextInput from './UiTextInput.vue';

export default {
  title: 'UI/UiFormGroup',
  component: UiFormGroup,
  argTypes: {
    label: { control: 'text' },
    error: { control: 'text' },
    required: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiFormGroup, UiTextInput },
  setup() { return { args }; },
  template: `<UiFormGroup v-bind="args"><UiTextInput placeholder='Inside group' /></UiFormGroup>`,
});

export const Default = Template.bind({});
Default.args = { label: 'Label', required: false };

export const WithError = Template.bind({});
WithError.args = { label: 'Label', error: 'This field is required', required: true }; 