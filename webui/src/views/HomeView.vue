<script>
import Post from '../components/Post.vue'
import FetchWindow from '../components/FetchWindow.vue'
import UserItemVue from '../components/UserItem.vue'
export default {
	props: ['username', 'userId'],
  components: { Post , FetchWindow},
	data () {
		return {
			posts: [],
			isLogged: false,
			show: false
		}
	},
	mounted() {
		this.isLogged = true
		this.refreshPosts()

	},
	methods: {
		async refreshPosts() {
			try {
				let response = await this.$axios.get('/feed/' + this.userId)
				this.posts = response.data
				console.log(response.data)
			}
			catch(e){

			}
		}
	}
}
</script>

<template>
	<div class="py-5 section">
		<div v-if="this.posts.length>0" class="d-flex justify-content-center bg-black text-white py-3" >
			<div >{{this.username}} {{this.userId}}
				<div v-for="post in posts"
					:key="post.id">
					<Post :post="post" :userId="this.userId" class="row py-2"></Post>
				</div>
			</div>
		</div>
		<div v-else class="py-5 text-white bg-black  text-center">nothing to show</div>
		
		<!-- Modal -->
			
	</div>
</template>

<style>
</style>
