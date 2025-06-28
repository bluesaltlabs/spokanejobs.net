import ItemEmptyState from './ItemEmptyState.vue';

export default {
  title: 'Common/ItemEmptyState',
  component: ItemEmptyState,
  parameters: {
    docs: {
      description: {
        component: 'A component for displaying empty states when no items are found or available.'
      }
    }
  }
};

const Template = (args) => ({
  components: { ItemEmptyState },
  setup() { return { args }; },
  template: `
    <ItemEmptyState v-bind="args">
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
        </svg>
      </template>
      <template #title>No items found</template>
      <template #desc>Try adding some items to see them here.</template>
    </ItemEmptyState>
  `,
});

export const Default = Template.bind({});
Default.args = {};

export const CompaniesEmpty = (args) => ({
  components: { ItemEmptyState },
  setup() { return { args }; },
  template: `
    <ItemEmptyState v-bind="args">
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
          <polyline points="9,22 9,12 15,12 15,22"></polyline>
        </svg>
      </template>
      <template #title>No companies found</template>
      <template #desc>Companies will appear here once they are added to the system.</template>
    </ItemEmptyState>
  `,
});

export const JobsEmpty = (args) => ({
  components: { ItemEmptyState },
  setup() { return { args }; },
  template: `
    <ItemEmptyState v-bind="args">
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"/>
        </svg>
      </template>
      <template #title>No jobs found</template>
      <template #desc>Check back later for new opportunities in the Spokane area.</template>
    </ItemEmptyState>
  `,
});

export const SearchEmpty = (args) => ({
  components: { ItemEmptyState },
  setup() { return { args }; },
  template: `
    <ItemEmptyState v-bind="args">
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="11" cy="11" r="8"></circle>
          <path d="m21 21-4.35-4.35"></path>
        </svg>
      </template>
      <template #title>No results found</template>
      <template #desc>Try adjusting your search criteria or browse all items.</template>
    </ItemEmptyState>
  `,
});

export const ErrorState = (args) => ({
  components: { ItemEmptyState },
  setup() { return { args }; },
  template: `
    <ItemEmptyState v-bind="args">
      <template #icon>
        <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="15" y1="9" x2="9" y2="15"></line>
          <line x1="9" y1="9" x2="15" y2="15"></line>
        </svg>
      </template>
      <template #title>Something went wrong</template>
      <template #desc>We encountered an error while loading the data. Please try again later.</template>
    </ItemEmptyState>
  `,
});
