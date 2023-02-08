<script>
import fetchwin from './FetchWindow.vue'
import uploadWindow from './UploadWindow.vue'
export default {
    props:['userId'],
    components: {
        fetchwin,
        uploadWindow
    },
    data () {
        return {
            currentFetch: '',
            userList: []

        }
    },
    methods: {
        async fetchUsername(){
            
            try {
                
                console.log(this.currentFetch)
                let response = await this.$axios.get('/result', {
                    params: {
                        username : this.currentFetch
                    }
                })
                this.userList = response.data
                
            } catch(e) {
              //errors
            }

        
        },
        logout(){
            this.$emit('logout')
            
        },
        redirect(user){
            this.$emit('redirect',user)
        }
    }
}
</script>


<template>
    <div>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
        <a class="navbar-brand" href="#" >WASAPHOTO</a>
        <button 
            class="navbar-toggler "
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#toggleMobileMenu"

            aria-controls="toggleMobileMenu"
            aria-expanded="false"
            aria-label="Toggle navigation"

        >
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="align-middle ">
            <input class="text-center form-control-sm " type="text" placeholder="search" v-model="currentFetch"> 
            <button class="btn btn-outline-light btn-sm px-1 m-1" type="submit" data-bs-toggle="modal" data-bs-target="#fetchWindow" @click="fetchUsername">search</button>
            <button class="btn btn-outline-light btn-sm px-1 m-1" type="file" data-bs-toggle="modal" data-bs-target="#uploadWindow" @click="upload">upload </button>
        </div>
        <div class="collapse navbar-collapse justify-content-end icons " id="toggleMobileMenu" >
            <ul class="navbar-nav text-center ">
                <li class="nav-item active">
                    <a class="nav-link" href="#">Home</a>
                </li>
                <li class="nav-item active">
                    <a class="nav-link" href="#profile">Profile</a>
                </li>
                <li class="nav-item active">
                    <a class="nav-link" href="#login" @click="logout">Logout</a>
                </li>
                
            </ul>
        </div>
    </nav>
        <fetchwin :userId="this.userId" :userList="this.userList" @goOnProfile="redirect"></fetchwin>
        <uploadWindow @reload="this.$emit('reloadHome')" > </uploadWindow>
    </div>
</template>

<style>

</style>

