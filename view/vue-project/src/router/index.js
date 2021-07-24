import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Training from '../views/Training.vue'
import Battle from '../views/Battle.vue'
import Dashboard from '../views/Dashboard.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/training',
    name: 'Training',
    component: Training 
  },
  {
    path: '/battle',
    name: 'Battle',
    component: Battle 
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard 
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
