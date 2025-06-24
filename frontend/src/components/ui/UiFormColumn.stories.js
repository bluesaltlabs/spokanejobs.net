import UiFormColumn from './UiFormColumn.vue';
import UiTextInput from './UiTextInput.vue';

export default {
  title: 'UI/UiFormColumn',
  component: UiFormColumn,
};

const Template = (args) => ({
  components: { UiFormColumn, UiTextInput },
  setup() { return { args }; },
  template: `<UiFormColumn><UiTextInput placeholder='First' /><UiTextInput placeholder='Second' /></UiFormColumn>`,
});

export const Default = Template.bind({}); 