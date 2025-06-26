import FormColumn from './FormColumn.vue';
import TextInput from '../TextInput/TextInput.vue';

export default {
  title: 'UI/FormColumn',
  component: FormColumn,
};

const Template = (args) => ({
  components: { FormColumn, TextInput },
  setup() { return { args }; },
  template: `<FormColumn><TextInput placeholder='First' /><TextInput placeholder='Second' /></FormColumn>`,
});

export const Default = Template.bind({});
