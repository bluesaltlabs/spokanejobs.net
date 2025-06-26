import SkeletonText from './Text.vue';

export default {
  title: 'Skeleton/Text',
  component: SkeletonText,
  argTypes: {
    lines:          { control: 'number', min: 0, step: 1 },
    lineHeight:     { control: 'number', min: 0, step: 1 },
    variant:        { control: 'select', options: ['default', 'title', 'subtitle'] },
    lastLineWidth:  {},
    animated:       { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { SkeletonText },
  setup() { return { args }; },
  template: '<SkeletonText v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {
  lines: 1,
  lineHeight: 16,
  variant: 'default',
  lastLineWidth: '',
  animated: true,
};
