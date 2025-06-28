<script setup>
const appName = import.meta.env.VITE_APP_NAME

import { XIcon, FacebookIcon, BlueskyIcon, MastodonIcon } from '@/components/icons';
import { Container } from '@/components/ui';

const socialLinks = [
  {
    name: 'Facebook',
    url: 'https://www.facebook.com/groups/spokanejobsnet',
    icon: FacebookIcon
  },
  {
    name: 'X (Twitter)',
    url: 'https://x.com/SpokaneJobsNet',
    icon: XIcon
  },
  {
    name: 'Bluesky',
    url: 'https://bsky.app/profile/spokanejobs.bsky.social',
    icon: BlueskyIcon
  },
  {
    name: 'Mastodon',
    url: 'https://mastodon.social/@SpokaneJobs',
    icon: MastodonIcon
  }
];
</script>

<template>
  <Container>
    <h1>Contact</h1>
    <hr class="divider" />

    <p class="contact-description">
      Connect with us on social media!
    </p>

    <div class="social-links">
      <a
        v-for="link in socialLinks"
        :key="link.name"
        :href="link.url"
        target="_blank"
        rel="noopener noreferrer"
        class="social-link"
      >
        <component :is="link.icon" class="social-icon" />
        <span>{{ link.name }}</span>
      </a>
    </div>
  </Container>
</template>

<style scoped>
.divider {
  margin: 1.5rem 0 2rem 0;
  border: none;
  border-top: 1px solid var(--color-border);
}

.contact-description {
  margin-bottom: 3rem;
  color: var(--color-text-secondary);
  line-height: 1.6;
}

.social-links {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1.5rem;
  max-width: 800px;
}

.social-link {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--border-radius-medium);
  box-shadow: 0 1px 3px var(--color-shadow);
  padding: 2rem 1.5rem;
  text-decoration: none;
  color: inherit;
  transition: all 0.2s;
  font-weight: 600;
  font-size: 1.1rem;
}

.social-link:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px var(--color-shadow-elevated);
  border-color: var(--color-border-hover);
}

.social-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--color-primary-600);
}

@media (max-width: 600px) {
  .social-links {
    grid-template-columns: 1fr;
    gap: 1rem;
  }

  .social-link {
    padding: 1.5rem 1rem;
  }
}
</style>
