import GithubIcon from './github.vue';

export default {
  title: 'Icons/github',
  component: GithubIcon,
};

const Template = (args) => ({
  components: { GithubIcon },
  setup() { return { args }; },
  template: '<GithubIcon v-bind="args"></GithubIcon>',
});

export const Default = Template.bind({});
Default.args = { width: 50, height: 50, variant: 'default' };
