import UiRichTextInput from './UiRichTextInput.vue';

export default {
  title: 'UI/UiRichTextInput',
  component: UiRichTextInput,
  argTypes: {
    modelValue: { control: 'text' },
    mode: { control: 'select', options: ['markdown', 'json', 'html'] },
    placeholder: { control: 'text' },
    disabled: { control: 'boolean' },
  },
};

const Template = (args) => ({
  components: { UiRichTextInput },
  setup() { return { args }; },
  template: '<UiRichTextInput v-bind="args" />',
});

export const Markdown = Template.bind({});
Markdown.args = { mode: 'markdown', placeholder: 'Write markdown...' };

export const JSON = Template.bind({});
JSON.args = { mode: 'json', placeholder: 'Write JSON...' };

export const HTML = Template.bind({});
HTML.args = { mode: 'html', placeholder: 'Write HTML...' };

export const Disabled = Template.bind({});
Disabled.args = { mode: 'markdown', disabled: true, placeholder: 'Disabled rich text' }; 