import ProfileIcon from './Profile.vue';

export default {
  title: 'Icons/Profile',
  component: ProfileIcon,
};

const Template = (args) => ({
  components: { ProfileIcon },
  setup() { return { args }; },
  template: '<ProfileIcon  v-bind="args"></ProfileIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
