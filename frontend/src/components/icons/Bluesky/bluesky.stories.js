import BlueskyIcon from './Bluesky.vue';

export default {
  title: 'Icons/Bluesky',
  component: BlueskyIcon,
};

const Template = (args) => ({
  components: { BlueskyIcon },
  setup() { return { args }; },
  template: '<BlueskyIcon v-bind="args"></BlueskyIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
