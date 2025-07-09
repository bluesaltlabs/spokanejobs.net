import Calendar from './Calendar.vue';

export default {
  title: 'Icons/Calendar',
  component: Calendar,
};

const Template = (args) => ({
  components: { Calendar },
  setup() { return { args }; },
  template: '<Calendar v-bind="args"></Calendar>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
