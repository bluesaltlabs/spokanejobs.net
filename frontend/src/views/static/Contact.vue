<script setup>
const appName = import.meta.env.VITE_APP_NAME

import { XIcon, FacebookIcon, BlueskyIcon, MastodonIcon } from '@/components/icons';
import { Container } from '@/components/ui';
import { ItemContainer, ItemCard } from '@/components/common';

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

const handleSocialClick = (url) => {
  window.open(url, '_blank', 'noopener,noreferrer');
};
</script>

<template>
  <Container>
    <h1>Contact</h1>
    <hr class="divider" />

    <p class="contact-description">
      Connect with us on social media!
    </p>

    <ItemContainer grid customClass="social-links">
      <ItemCard
        v-for="link in socialLinks"
        :key="link.name"
        :onClick="() => handleSocialClick(link.url)"
        customClass="social-link"
      >
        <template #header>
          <div class="social-link-content">
            <component :is="link.icon" class="social-icon" />
            <span>{{ link.name }}</span>
          </div>
        </template>
      </ItemCard>
    </ItemContainer>
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
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)) !important;
  gap: 1.5rem !important;
  max-width: 800px;
}

.social-link {
  padding: 2rem 1.5rem !important;
  font-weight: 600;
  font-size: 1.1rem;
  text-align: center;
}

.social-link-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.social-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--color-primary-600);
}

@media (max-width: 600px) {
  .social-links {
    grid-template-columns: 1fr !important;
    gap: 1rem !important;
  }

  .social-link {
    padding: 1.5rem 1rem !important;
  }
}
</style>
