import Button from './Button.vue';

export default {
  title: 'UI/Button',
  component: Button,
  argTypes: {
    variant: { control: 'select', options: ['primary', 'secondary'] },
    loading: { control: 'boolean' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { Button },
  setup() { return { args }; },
  template: '<Button v-bind="args">Button</Button>',
});

export const Primary = Template.bind({});
Primary.args = { variant: 'primary' };

export const Secondary = Template.bind({});
Secondary.args = { variant: 'secondary' };

export const Loading = Template.bind({});
Loading.args = { loading: true };

export const Disabled = Template.bind({});
Disabled.args = { disabled: true };
