import SkeletonTableRow from './TableRow.vue';

export default {
  title: 'Skeleton/TableRow',
  component: SkeletonTableRow,
  argTypes: {
    columns:      { control: 'number' },
    columnWidths: { control: 'array' },
  },
};

const Template = (args) => ({
  components: { SkeletonTableRow },
  setup() { return { args }; },
  template: '<table style="width:100%"><tbody><SkeletonTableRow v-bind="args" /></tbody></table>',
});

export const Default = Template.bind({});
Default.args = {
  columns: 3,
  columnWidths: ['40%', '30%', '30%'],
};
