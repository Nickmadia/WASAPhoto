<script>

export default {
  props:['extUser','userId'],
    data () {
        return {
          isFollowing:null
        }
    },
    mounted() {
        
    },
    methods: {
		async followUser(){
			try {
				console.log('stuff')
        console.log(this.userId)
				let response = await this.$axios.put("/users/" + this.userId +"/follow/" + this.extUser.user_id)
				console.log(response.data)
			} 
			catch (e){

			}
		}
    }
}
</script>


<template>
    <section class="h-100 d-flex justify-content-center bg-dark">
  <div class="container py-4 h-100">
    <div class="row d-flex justify-content-center align-items-center h-100">
      <div class="col col-lg-9 col-xl-7">
        <div class="card">
         
          <div class="p-4 d-flex text-black" style="background-color: #f8f9fa;">
            <div class="ms-4  " >
              <h5>{{this.extUser.username}}</h5>
              <button type="button" @click="followUser" class="btn btn-outline-dark" data-mdb-ripple-color="dark"
                style="z-index: 1;">
                Follow
              </button>
            </div>
            <div class="d-flex ms-auto text-center py-1">
              <div>
                <p class="mb-1 h5">{{this.extUser.media_count}}</p>
                <p class="small text-muted mb-0">Photos</p>
              </div>
              <div class="px-3">
                <p class="mb-1 h5">{{this.extUser.followers_count}}</p>
                <p class="small text-muted mb-0">Followers</p>
              </div>
              <div>
                <p class="mb-1 h5">{{this.extUser.following_count}}</p>
                <p class="small text-muted mb-0">Following</p>
              </div>
            </div>
            
          </div>
          <div class="card-body p-4 text-black">
            
            <div class="d-flex justify-content-between align-items-center mb-4">
              <p class="lead fw-normal mb-0">Recent photos</p>
              <p class="mb-0"><a href="#!" class="text-muted">Show all</a></p>
            </div>
            <div class="row g-2">
              <div class="col mb-2">
                <img src="https://mdbcdn.b-cdn.net/img/Photos/Lightbox/Original/img%20(112).webp"
                  alt="image 1" class="w-100 rounded-3">
              </div>
              <div class="col mb-2">
                <img src="https://mdbcdn.b-cdn.net/img/Photos/Lightbox/Original/img%20(107).webp"
                  alt="image 1" class="w-100 rounded-3">
              </div>
            </div>
            <div class="row g-2">
              <div class="col">
                <img src="https://mdbcdn.b-cdn.net/img/Photos/Lightbox/Original/img%20(108).webp"
                  alt="image 1" class="w-100 rounded-3">
              </div>
              <div class="col">
                <img src="https://mdbcdn.b-cdn.net/img/Photos/Lightbox/Original/img%20(114).webp"
                  alt="image 1" class="w-100 rounded-3">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
</template>