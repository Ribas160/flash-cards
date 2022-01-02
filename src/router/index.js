import Vue from 'vue'
import VueRouter from 'vue-router'
import Lists from '../views/Lists.vue'
import List from '../views/List.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'lists',
    component: Lists
  },
  {
    path: '/list/:id',
    name: 'list',
    component: List
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
