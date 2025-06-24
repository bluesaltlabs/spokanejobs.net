import UiMultiSelect from './UiMultiSelect.vue';
import { ref } from 'vue';

export default {
  title: 'UI/UiMultiSelect',
  component: UiMultiSelect,
  argTypes: {
    modelValue: { control: 'object' },
    options: { control: 'object' },
    placeholder: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiMultiSelect },
  setup() {
    const value = ref(args.modelValue || []);
    return { args, value };
  },
  template: '<UiMultiSelect v-bind="args" v-model:modelValue="value" />',
});

export const Default = Template.bind({});
Default.args = {
  modelValue: [],
  options: [
    { value: 'apple', label: 'Apple' },
    { value: 'banana', label: 'Banana' },
    { value: 'cherry', label: 'Cherry' },
  ],
  placeholder: 'Select fruits',
};

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  disabled: true,
}; 