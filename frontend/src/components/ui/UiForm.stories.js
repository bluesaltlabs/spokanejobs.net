import UiForm from './UiForm.vue';
import UiFormGroup from './UiFormGroup.vue';
import UiTextInput from './UiTextInput.vue';
import UiButton from './UiButton.vue';

export default {
  title: 'UI/UiForm',
  component: UiForm,
};

const Template = (args) => ({
  components: { UiForm, UiFormGroup, UiTextInput, UiButton },
  setup() { return { args }; },
  template: `<UiForm @submit="alert('Submitted!')"><UiFormGroup label='Name'><UiTextInput /></UiFormGroup><UiButton type='submit'>Submit</UiButton></UiForm>`,
});

export const Default = Template.bind({}); 