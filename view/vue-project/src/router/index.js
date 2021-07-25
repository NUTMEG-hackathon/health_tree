import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Select from '../views/Select.vue'
import Training from '../views/Training.vue'
import Training2 from '../views/Training2.vue'
import Training3 from '../views/Training3.vue'
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
    path: '/select',
    name: 'Select',
    component: Select
  },
  {
    path: '/training',
    name: 'Training',
    component: Training 
  },
  {
    path: '/training2',
    name: 'Training2',
    component: Training2 
  },
  {
    path: '/training3',
    name: 'Training3',
    component: Training3 
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
