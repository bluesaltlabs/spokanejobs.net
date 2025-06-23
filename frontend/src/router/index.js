import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/static/Home.vue'
import AboutView from '@/views/static/About.vue'
import ContactView from '@/views/static/Contact.vue'
import SitemapView from '@/views/static/Sitemap.vue'
import CompaniesGridView from '@/views/companies/Grid.vue'
//import CompaniesListView from '@/views/companies/List.vue'
import CompanyDetailView from '@/views/companies/Detail.vue'
import JobsListView from '@/views/jobs/List.vue'
import JobDetailView from '@/views/jobs/Detail.vue'
import ProfileView from '@/views/profile/Detail.vue'
import ProfileEditView from '@/views/profile/Edit.vue'

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
      path: '/contact',
      name: 'contact',
      component: ContactView,
    },
    {
      path: '/sitemap',
      name: 'sitemap',
      component: SitemapView,
    },
    {
      path: '/companies',
      name: 'companies-grid',
      //component: CompaniesListView,
      component: CompaniesGridView,
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
