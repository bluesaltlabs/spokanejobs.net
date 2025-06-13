import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import CompaniesListView from '@/views/CompaniesListView.vue'
import JobsListView from '@/views/JobsListView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
    {
      path: '/companies',
      name: 'companies',
      component: CompaniesListView,
    },
    {
      path: '/jobs',
      name: 'jobs',
      component: JobsListView,
    },
  ],
})

export default router
