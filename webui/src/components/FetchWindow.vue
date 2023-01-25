<script>
export default {
  props: ['userList', 'userId'],
  data() {
    return {

    }
  },
  methods: {
		
  },
  computed: {
	removeMyId(){
		console.log(this.userId)
		return this.userList.filter(n => n.user_id != this.userId)
	}
  }
}
</script>

<template>
     <div class="modal fade" id="fetchWindow" tabindex="-1" aria-hidden="true">
				<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
					<div class="modal-content">
						<div class="modal-header">
							<h5 class="modal-title " >Search</h5>
							<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
						</div>
						<div class="modal-body">
							<div class="row d-flex justify-content-center align-items-center ">
								<div >
									<div v-if="userList.length>0" class="card bg-dark text-white ">
										<ul class="list-group list-group text-center ">
											<a v-for="user in removeMyId" 
                          :key="user.id"
                          @click="this.$emit('goOnProfile',user)" data-bs-dismiss="modal" data-bs-target="fetchWindow" class="list-group-item  ">
						  					{{user.username}}
											
											</a>
										</ul>
									</div>
                  <div v-else class="text-center">
                    no result
                  </div>
								</div>
							</div>
						</div>
					
					</div>
				</div>
			</div>


</template>