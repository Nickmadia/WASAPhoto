<script >
import { RouterLink, RouterView } from 'vue-router'
import navbar from './components/NavBar.vue'

export default {
	components : {
		navbar
	},
	data () {
		return {
			username: '',
			userId: '',
			isLogged: false,
			currentExtUser: null

		}
	},
	methods: {
		Login (id) {
			this.isLogged = true
			this.userId = id
		},
		doLogout() {
			this.isLogged = false
			console.log('logging out .....')
			this.userId = ''
		},
		getUsername(username) {
			this.username = username
		},
		redirectToProfile(user){
			this.currentExtUser = user
			this.$router.push({name:'extProfile', params:{ username: user.username}})
		}
	}
}
</script>


<template>
	<div>
		<header>
			 <navbar :userId="this.userId" v-if="isLogged" @logout="doLogout" @redirect="this.redirectToProfile"> </navbar>
		</header>

		<main>
			<routerView :username="this.username" :extUser="currentExtUser" :userId="this.userId" @login="Login"  
						@getUsername="getUsername" ></routerView>
		</main>
	</div>

	
	
</template>

<style>
</style>
