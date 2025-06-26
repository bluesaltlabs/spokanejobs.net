import FormRow from './FormRow.vue';
import TextInput from '../TextInput/TextInput.vue';

export default {
  title: 'UI/FormRow',
  component: FormRow,
};

const Template = (args) => ({
  components: { FormRow, TextInput },
  setup() { return { args }; },
  template: `<FormRow><TextInput placeholder='First' /><TextInput placeholder='Second' /></FormRow>`,
});

export const Default = Template.bind({});
