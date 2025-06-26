import SkeletonCard from './Card.vue';

export default {
  title: 'Skeleton/Card',
  component: SkeletonCard,
  argTypes: {
    showImage:        { control: 'boolean' },
    imageWidth:       { },
    imageHeight:      { },
    showTitle:        { control: 'boolean' },
    showDescription:  { control: 'boolean' },
    descriptionLines: { control: 'number' },
    lastLineWidth:    { },
    showButtons:      { control: 'boolean' },
    buttonCount:      { control: 'number' },
    buttonWidths:     { control: 'select' },
    animated:         { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { SkeletonCard },
  setup() { return { args }; },
  template: '<SkeletonCard v-bind="args" />',
});

export const Default = Template.bind({});
Default.args = {
  showImage: true,
  imageWidth: "100%",
  imageHeight: "200px",
  showTitle: true,
  showDescription: true,
  descriptionLines: 3,
  lastLineWidth: "60%",
  showButtons: true,
  buttonCount: 2,
  buttonWidths: "80px",
  animated: true,
};
