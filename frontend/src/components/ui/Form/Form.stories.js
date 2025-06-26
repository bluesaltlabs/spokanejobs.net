import Form from './Form.vue';
import FormGroup from '../FormGroup/FormGroup.vue';
import TextInput from '../TextInput/TextInput.vue';
import Button from '../Button/Button.vue';

export default {
  title: 'UI/Form',
  component: Form,
};

const Template = (args) => ({
  components: { Form, FormGroup, TextInput, Button },
  setup() { return { args }; },
  template: `<Form @submit="alert('Submitted!')"><FormGroup label='Name'><TextInput /></FormGroup><Button type='submit'>Submit</Button></Form>`,
});

export const Default = Template.bind({});
