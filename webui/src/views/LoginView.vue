<script>
export default {
	props : ['test'],
	data () {
		return {
			errormsg: null,
            username : '',
            islogged : false
		}
	},
    methods: {
        redirectProfile(id){
            //sends data to parent
            this.$emit('login', id)
            this.$emit('getUsername', this.username)
            
            this.$router.push({path: '/'})
        },
        async doLogin () {
            try {
                const body = {username: this.username}

                let response = await this.$axios.post('/session', body)
                  console.log(response.data)
                this.redirectProfile(response.data.identifier)
                this.$axios.defaults.headers.common['Authorization'] = `${response.data.identifier}`;
            } catch(e) {
              //errors
            }

        }
    },
    mounted() {
     
    }
}
</script>

<template>
	
    <section class="vh-100 bg-dark">
  <div class=" py-5 h-100">
    <div class="row d-flex justify-content-center align-items-center h-100">
      <div class="col-12 col-md-8 col-lg-6 col-xl-5">
        <div class="card bg-secondary text-white" style="border-radius: 2rem;">
          <div class="card-body p-5 text-center">

            <div class="mb-md-3 mt-md-2 pb-1">

              <h2 class="fw-bold mb-2 text-uppercase">Login</h2>
              <p class="text-white-50 mb-3">Please enter your username {{test}}</p>

              <div class="form-outline form-white mb-4 ">
                <input  v-model="username" class="form-control text-center form-control-lg" />
                <label class="form-label mt-3" for="typeEmailX">Username</label>
              </div>

              <button class="btn btn-outline-light btn-lg px-5" type="submit" @click="doLogin">Login</button>

            </div>

           

          </div>
        </div>
      </div>
    </div>
  </div>
</section>
</template>