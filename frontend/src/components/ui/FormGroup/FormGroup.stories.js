import FormGroup from './FormGroup.vue';
import TextInput from '../TextInput/TextInput.vue';

export default {
  title: 'UI/FormGroup',
  component: FormGroup,
  argTypes: {
    label: { control: 'text' },
    error: { control: 'text' },
    required: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { FormGroup, TextInput },
  setup() { return { args }; },
  template: `<FormGroup v-bind="args"><TextInput placeholder='Inside group' /></FormGroup>`,
});

export const Default = Template.bind({});
Default.args = { label: 'Label', required: false };

export const WithError = Template.bind({});
WithError.args = { label: 'Label', error: 'This field is required', required: true };
