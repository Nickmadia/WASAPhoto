<script>
import Post from '../components/Post.vue'
export default {
  components: {Post},
  props:['userId', 'username'],
    data () {
        return {
          followerList: [],
          followingList: [],
          posts:[],
          currentUsername:'',
          usernameVar:''
        }
    },
    methods: {
      async changeUsername() {
        try{
          let body = {username:this.currentUsername}
          console.log(body)
          let response = await this.$axios.put('/users/' + this.userId + '/username', body)
          if (response.status == 204) {
            this.usernameVar = this.currentUsername
          }
          this.currentUsername =''

        } catch(e) {

        }
      },
      async getProfileInfo() {
        try {
          let response = await this.$axios.get("/users/" + this.userId + "/info")
          if (response.data.followers != null) {
            this.followerList = response.data.followers 
          }
          if (response.data.following != null) {

            this.followingList = response.data.following 
          }
          if (response.data.posts != null) {
            this.posts = response.data.posts
          }
          
        } catch(e) {

        }
      },async deletePost(post_id){
        try {
          let response = await this.$axios.delete('/media/'+ post_id)
          this.posts = this.posts.filter(x => x.id!= post_id) 
        } catch(e)
        {

        }
      }
    },
    async beforeMount() {
      await this.getProfileInfo()
      this.usernameVar = this.username
    }
}
</script>


<template>
    <section class="h-100 d-flex justify-content-center bg-black pt-5 ">
  <div class="container py-4 h-100">
    <div class="row d-flex justify-content-center  h-100">
      <div class="col col-lg-3 col-xl-8">
        <div class="border-bottom">
         
          <div class="p-2 my-2 d-flex text-white bg-black" style="background-color: #f8f9fa;">
            <div class="ms-4  " >
              <h2 class="text-primary fw-bold">{{this.usernameVar}}</h2>
              
              <div class="d-flex ">
                <div>
                  <button  type="button"  class="btn  btn-primary " data-mdb-ripple-color="dark"
                    style="z-index: 1;" data-bs-toggle="modal" data-bs-target="#changeUsernameWindow">
                    Edit Profile
                  </button>
                  
                </div>
              </div>
            </div>
            <div class="d-flex ms-auto text-center py-1">
              <div>
                <p class="mb-1 h5">{{this.posts.length}}</p>
                <p class="small text-muted mb-0">Photos</p>
              </div>
              <div class="px-3">
                <p class="mb-1 h5">{{this.followerList.length}}</p>
                <p class="small text-muted mb-0">Followers</p>
              </div>
              <div>
                <p class="mb-1 h5">{{this.followingList.length}}</p>
                <p class="small text-muted mb-0">Following</p>
              </div>
            </div>
            
          </div>
           
        </div>
            <div class="d-flex mt-2 justify-content-center card-body py-2 bg-black">
              <div class="">
                <div v-for="item in posts" :key="item" class="my-4">
                  <Post :post="item" :userId="this.userId" :userName="this.username" @delPost="deletePost"></Post>
                </div>
              </div>
            </div>
      </div>
    </div>
  </div>
  <div class="modal fade" id="changeUsernameWindow" tabindex="-1" aria-hidden="true">
				<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
					<div class="modal-content">
						<div class="modal-header">
							<h5 class="modal-title " >Edit Username</h5>
							<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
						</div>
						<div class="modal-body">
							<div class="row d-flex justify-content-center align-items-center ">
                                <div class="d-flex">
								    <input class="form-control" type="text" placeholder="Username" v-model="currentUsername">
                                    <button class="btn btn-primary ms-auto " data-bs-dismiss="modal" @click="this.changeUsername">Edit</button>
                                </div>
							</div>
						</div>
					
					</div>
				</div>
			</div>
</section>
</template>