import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/pages/Home/HomePage.vue'
import WelcomePage from '@/pages/Welcome/WelcomePage.vue'

const routes = [
    { path: '/', component: WelcomePage },
    { path: '/code', component: HomePage },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})
    

export default router