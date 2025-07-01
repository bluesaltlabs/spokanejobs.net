import ScrapingStatus from './ScrapingStatus.vue';

export default {
  title: 'Common/ScrapingStatus',
  component: ScrapingStatus,
  argTypes: {
    status: {
      control: { type: 'select' },
      options: ['built', 'in progress', 'error', 'not built'],
    },
    lastScrapedAt: { control: 'text' },
    scraperCreatedAt: { control: 'text' },
    activeJobsCount: { control: 'number' },
    scraperError: { control: 'text' },
  },
};

const Template = (args) => ({
  components: { ScrapingStatus },
  setup() { return { args }; },
  template: '<ScrapingStatus v-bind="args" />',
});

export const Built = Template.bind({});
Built.args = {
  status: 'built',
  lastScrapedAt: '2025-06-30T18:00:00Z',
  scraperCreatedAt: '2024-12-01T10:00:00Z',
  activeJobsCount: 12,
  scraperError: '',
};

export const InProgress = Template.bind({});
InProgress.args = {
  status: 'in progress',
  lastScrapedAt: '2025-06-30T18:00:00Z',
  scraperCreatedAt: '2024-12-01T10:00:00Z',
  activeJobsCount: 8,
  scraperError: '',
};

export const Error = Template.bind({});
Error.args = {
  status: 'error',
  lastScrapedAt: '2025-06-30T18:00:00Z',
  scraperCreatedAt: '2024-12-01T10:00:00Z',
  activeJobsCount: 0,
  scraperError: 'Failed to connect to job board',
};

export const NotBuilt = Template.bind({});
NotBuilt.args = {
  status: 'not built',
  lastScrapedAt: '',
  scraperCreatedAt: '',
  activeJobsCount: undefined,
  scraperError: '',
};
