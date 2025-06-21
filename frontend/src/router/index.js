import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
//import CompaniesListView from '@/views/CompaniesListView.vue'
import CompaniesGridView from '@/views/CompaniesGridView.vue'
import CompanyDetailView from '@/views/CompanyDetailView.vue'
import CompanyCreateView from '@/views/CompanyCreateView.vue'
import CompanyEditView from '@/views/CompanyEditView.vue'
import JobsListView from '@/views/JobsListView.vue'
import JobDetailView from '@/views/JobDetailView.vue'
import ProfileView from '@/views/ProfileView.vue'
import ProfileEditView from '@/views/ProfileEditView.vue'

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
      name: 'companies-grid',
      component: CompaniesGridView,
    },
    {
      path: '/companies/create',
      name: 'company-create',
      component: CompanyCreateView,
    },
    {
      path: '/companies/:slug',
      name: 'company-detail',
      component: CompanyDetailView,
    },
    {
      path: '/companies/:slug/edit',
      name: 'company-edit',
      component: CompanyEditView,
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
    {
      path: '/profile',
      name: 'profile',
      component: ProfileView,
    },
    {
      path: '/profile/edit',
      name: 'profile-edit',
      component: ProfileEditView,
    }
  ],
})

export default router
