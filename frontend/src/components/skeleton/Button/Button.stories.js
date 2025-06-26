import SkeletonButton from './Button.vue';

export default {
  title: 'Skeleton/Button',
  component: SkeletonButton,
  argTypes: {
    width:        { },
    height:       { },
    borderRadius: { },
    animated:     { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { SkeletonButton },
  setup() { return { args }; },
  template: '<SkeletonButton v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {
  width: "120px",
  height: "40px",
  borderRadius: "var(--border-radius-small)",
  animated: true,
};
