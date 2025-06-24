import UiSelect from './UiSelect.vue';
import { ref } from 'vue';

export default {
  title: 'UI/UiSelect',
  component: UiSelect,
  argTypes: {
    modelValue: { control: 'text' },
    options: { control: 'object' },
    placeholder: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiSelect },
  setup() {
    const value = ref(args.modelValue || '');
    return { args, value };
  },
  template: '<UiSelect v-bind="args" v-model:modelValue="value" />',
});

export const Default = Template.bind({});
Default.args = {
  modelValue: '',
  options: [
    { value: 'one', label: 'One' },
    { value: 'two', label: 'Two' },
    { value: 'three', label: 'Three' },
  ],
  placeholder: 'Select an option',
};

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  disabled: true,
}; 