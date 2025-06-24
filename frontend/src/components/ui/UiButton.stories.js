import UiButton from './UiButton.vue';

export default {
  title: 'UI/UiButton',
  component: UiButton,
  argTypes: {
    variant: { control: 'select', options: ['primary', 'secondary'] },
    loading: { control: 'boolean' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiButton },
  setup() { return { args }; },
  template: '<UiButton v-bind="args">Button</UiButton>',
});

export const Primary = Template.bind({});
Primary.args = { variant: 'primary' };

export const Secondary = Template.bind({});
Secondary.args = { variant: 'secondary' };

export const Loading = Template.bind({});
Loading.args = { loading: true };

export const Disabled = Template.bind({});
Disabled.args = { disabled: true }; 