import UiFormRow from './UiFormRow.vue';
import UiTextInput from './UiTextInput.vue';

export default {
  title: 'UI/UiFormRow',
  component: UiFormRow,
};

const Template = (args) => ({
  components: { UiFormRow, UiTextInput },
  setup() { return { args }; },
  template: `<UiFormRow><UiTextInput placeholder='First' /><UiTextInput placeholder='Second' /></UiFormRow>`,
});

export const Default = Template.bind({}); 