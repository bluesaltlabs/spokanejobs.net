import Select from './Select.vue';
import { ref } from 'vue';

export default {
  title: 'UI/Select',
  component: Select,
  argTypes: {
    modelValue: { control: 'text' },
    options: { control: 'object' },
    placeholder: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { Select },
  setup() {
    const value = ref(args.modelValue || '');
    return { args, value };
  },
  template: '<Select v-bind="args" v-model:modelValue="value" />',
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
