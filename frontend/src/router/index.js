import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import CompaniesListView from '@/views/CompaniesListView.vue'
import CompanyDetailView from '@/views/CompanyDetailView.vue'
import JobsListView from '@/views/JobsListView.vue'
import JobDetailView from '@/views/JobDetailView.vue'

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
      name: 'companies-list',
      component: CompaniesListView,
    },
    {
      path: '/companies/:slug',
      name: 'company-detail',
      component: CompanyDetailView,
    },
    {
      path: '/jobs',
      name: 'jobs-list',
      component: JobsListView,
    },
    {
      path: '/jobs/:id',
      name: 'job-detail',
      component: JobDetailView,
    },
  ],
})

export default router
