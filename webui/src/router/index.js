import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Profile from '../views/ProfileView.vue'
import Login from '../views/LoginView.vue'
import osProfile from '../views/UserProfileView.vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: Login}, 
		{path: '/', name:'home', component: HomeView},
		{path: '/:username' , name:'extProfile' ,component: osProfile},
		{path:'/profile', name:'profile', component: Profile},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
