import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'
import AppelliView from '../views/AppelliView.vue'
import PrenotazioniView from '../views/PrenotazioniView.vue'
import LibrettoView from '../views/LibrettoView.vue'
import ProfiloView from '../views/ProfiloView.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', name: 'login', component: LoginView },
		{ path: '/dashboard', name: 'dashboard', component: DashboardView },
		{ path: '/dashboard/appelli', name: 'appelli', component: AppelliView },
		{ path: '/dashboard/prenotazioni', name: 'prenotazioni', component: PrenotazioniView },
		{ path: '/dashboard/libretto', name: 'libretto', component: LibrettoView },
		{ path: '/dashboard/profilo', name: 'profilo', component: ProfiloView }
	]
})

export default router