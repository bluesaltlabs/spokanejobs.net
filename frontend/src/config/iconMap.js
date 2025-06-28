import { defineAsyncComponent } from 'vue'

// Map icon names to components
const iconMap = {
  About:      defineAsyncComponent(() => import('@/components/icons/About/About.vue')),
  Bluesky:    defineAsyncComponent(() => import('@/components/icons/Bluesky/Bluesky.vue')),
  Companies:  defineAsyncComponent(() => import('@/components/icons/Companies/Companies.vue')),
  Contact:    defineAsyncComponent(() => import('@/components/icons/Contact/Contact.vue')),
  Export:     defineAsyncComponent(() => import('@/components/icons/Export/Export.vue')),
  Facebook:   defineAsyncComponent(() => import('@/components/icons/Facebook/Facebook.vue')),
  Github:     defineAsyncComponent(() => import('@/components/icons/Github/Github.vue')),
  Home:       defineAsyncComponent(() => import('@/components/icons/Home/Home.vue')),
  Import:     defineAsyncComponent(() => import('@/components/icons/Import/Import.vue')),
  Jobs:       defineAsyncComponent(() => import('@/components/icons/Jobs/Jobs.vue')),
  Location:   defineAsyncComponent(() => import('@/components/icons/Location/Location.vue')),
  Mastodon:   defineAsyncComponent(() => import('@/components/icons/Mastodon/Mastodon.vue')),
  Moon:       defineAsyncComponent(() => import('@/components/icons/Moon/Moon.vue')),
  Profile:    defineAsyncComponent(() => import('@/components/icons/Profile/Profile.vue')),
  Resources:  defineAsyncComponent(() => import('@/components/icons/Resources/Resources.vue')),
  Sun:        defineAsyncComponent(() => import('@/components/icons/Sun/Sun.vue')),
  View:       defineAsyncComponent(() => import('@/components/icons/View/View.vue')),
  X:          defineAsyncComponent(() => import('@/components/icons/X/X.vue')),
}

export default iconMap
