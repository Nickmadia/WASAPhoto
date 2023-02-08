<script>
import Post from '../components/Post.vue'


export default {
	props: ['username', 'userId'],
  components: { Post },
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
		async deletePost(post_id){
			try {
				let response = await this.$axios.delete('/media/'+ post_id)
				this.posts = this.posts.filter(x => x.id!= post_id) 
			} catch(e)
			{

			}
      },
		async refreshPosts() {
			try {
				let response = await this.$axios.get('/feed/' + this.userId)
				if(response.data != null) {
				this.posts = response.data
				}
				
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
			<div >
				<div v-for="post in posts"
					:key="post.id">
					<Post :post="post" :userId="this.userId" :userName="this.username" @delPost="deletePost" class="row py-2"></Post>
				</div>
			</div>
		</div>
		<div v-else class="py-5 text-white bg-black  text-center">nothing to show</div>
		
		<!-- Modal -->
			
	</div>
</template>

<style>
</style>
