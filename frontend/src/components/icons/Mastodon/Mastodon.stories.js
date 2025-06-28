import MastodonIcon from './Mastodon.vue';

export default {
  title: 'Icons/Mastodon',
  component: MastodonIcon,
};

const Template = (args) => ({
  components: { MastodonIcon },
  setup() { return { args }; },
  template: '<MastodonIcon  v-bind="args"></MastodonIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
