<script>
import Post from '../components/Post.vue'
export default {
  components: {Post},
  props:['extUser','userId', 'username'],
    data () {
        return {
          isFollowing:false,
          followerList: [],
          followingList: [],
          posts:[],
          isBanned: false

        }
    },
    mounted() {
        
    },
    methods: {
      async ban() {
      
      if (this.isBanned)
      {
        try {
        let response = await this.$axios.delete("/users/" + this.userId +"/bans/" + this.extUser.user_id)
				//reactive follow button, check if already folloews from followers list, update followers
			} 
			catch (e){
        
        }
      } else {
          try {
        let response = await this.$axios.put("/users/" + this.userId +"/bans/" + this.extUser.user_id)
				//reactive follow button, check if already folloews from followers list, update followers
			} 
			catch (e){
        
        }
      }
			
      this.isBanned = !this.isBanned
    },
		async followUser(){
      if (this.isFollowing)
      {
        try {
        let response = await this.$axios.delete("/users/" + this.userId +"/follows/" + this.extUser.user_id)
				//reactive follow button, check if already folloews from followers list, update followers
			} 
			catch (e){
        
        }
      } else {
          try {
        let response = await this.$axios.put("/users/" + this.userId +"/follows/" + this.extUser.user_id)
				//reactive follow button, check if already folloews from followers list, update followers
			} 
			catch (e){
        
        }
      }
			
      this.isFollowing = !this.isFollowing
		},
    async getProfileInfo() {
      try {
        let response = await this.$axios.get("/users/" + this.extUser.user_id + "/info")
        if (response.data.followers != null) {
          this.followerList = response.data.followers 
        }
        if (response.data.following != null) {

          this.followingList = response.data.following 
        }
        if (response.data.posts != null) {
          this.posts = response.data.posts
        }
        this.isBanned = response.data.is_banned
        
      } catch(e) {

      }
    },
    checkFollow() {
      for (let user of this.followerList) {
        if(user.user_id == this.userId) {

          return true
        }
      }
        console.log(this.followerList)
        return false
    }
    },
    async beforeMount() {
      await this.getProfileInfo()
      this.isFollowing = this.checkFollow()
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
              <h2 class="text-primary fw-bold">{{this.extUser.username}}</h2>
              
              <div class="d-flex ">
                <div>
                  <button v-if="!this.isFollowing" type="button" @click="followUser" class="btn  btn-primary " data-mdb-ripple-color="dark"
                    style="z-index: 1;">
                    Follow
                  </button>
                  <button v-else type="button" @click="followUser" class="btn btn-outline-secondary btn-black text-white" data-mdb-ripple-color="dark"
                    style="z-index: 1;">
                    Unfollow
                  </button>
                </div>
                <div class="">
                  <button v-if="!this.isBanned" type="button" @click="ban" class=" btn btn-outline-secondary text-white btn-danger " 
                    style="z-index: 1;">
                    Ban
                  </button>
                  <button v-else type="button" @click="ban" class="btn btn-outline-secondary text-white  " 
                    style="z-index: 1;">
                    Unban
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
                  <Post :post="item" :userId="this.userId" :userName="this.username"></Post>
                </div>
              </div>
            </div>
      </div>
    </div>
  </div>
</section>
</template>