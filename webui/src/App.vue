<script >
import { RouterLink, RouterView } from 'vue-router'
import sidebar from './components/NavBar.vue'
export default {
	components : {
		sidebar,
	}
}
</script>


<template>
	<div>
		<header>
			 <sidebar> </sidebar>
		</header>
		<main>
			<routerView></routerView>
		</main>
	</div>

	
	
</template>

<style>
</style>
