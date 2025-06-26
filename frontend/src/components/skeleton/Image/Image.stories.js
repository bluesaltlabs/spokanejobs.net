import SkeletonImage from './Image.vue';

export default {
  title: 'Skeleton/Image',
  component: SkeletonImage,
  argTypes: {
    width:        {},
    height:       {},
    borderRadius: {},
    animated:     { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { SkeletonImage },
  setup() { return { args }; },
  template: '<SkeletonImage v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {
  width: "100%",
  height: "200px",
  borderRadius: "var(--border-radius)",
  animated: true,
};
