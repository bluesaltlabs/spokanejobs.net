import '../src/assets/main.css';

//import { app } from '@storybook/vue3';
//import { createPinia } from 'pinia';

// Add Pinia to Storybook globally for Vue 3 (todo: figure out why this breaks storybook)
//app.use(createPinia());

/** @type { import('@storybook/vue3-vite').Preview } */
const preview = {
  parameters: {
    controls: {
      matchers: {
       color: /(background|color)$/i,
       date: /Date$/i,
      },
    },
  },
};

export default preview;
