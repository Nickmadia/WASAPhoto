<script>
import Comment from './Comment.vue'

export default {
    props: ['post'],
    components:{
        Comment
    },
    data() {
        return {
            commenttest: [ 
                    {text: 'that is great',username: 'jack',date: '12 dec 2022'},
                    {text: 'that is great',username: 'jack',date: '12 dec 2022'},
                    {text: 'that is great',username: 'jack',date: '12 dec 2022'},
             
                  
                ],
                user :null,
                img: null

        }
    },
    methods : {
        async getImg() {
            try {
                let response = await this.$axios.get('/media/' + this.post.id)
                this.img = response.data
                console.log(response.data)
            } catch (e) {

            }
        },
        async getUser() {
            try {
                let response = await this.$axios.get('/users/' + this.post.owner_id)
                this.user = response.data
                console.log(this.user)
                
                
            } catch (e) {

            }
        },
        getReadableDate() {
            let isodate = this.post.time_stamp
            var d = new Date(isodate)
            return d.toLocaleDateString('en-GB')
        },
        getb64() {
            return 'data:image/png;base64,' + this.img
        }
    },
    async mounted() {
        console.log(this.post)
        await this.getImg()
        await this.getUser()
        
    },
    computed : {
        
}
}
</script>
<template>
<div v-if="post!=null && user!=null">
    <div class="card bg-dark border-white " style="width: 40rem; ">
        <div class="card-body text-white">
            <div class="d-flex">
                <h4 class="card-title text-primary "> {{this.user.Username}}</h4>
            </div>
            <div class="imagecontainer ">
                <img :src="'data:image/png;base64,' + this.img" class="card-img-top" alt="Card image cap" >
            </div>
            <div class="d-flex border-bottom pb-2">
                <div class="pb-2 me-2">
                    <button class="btn btn-primary btn-sm ">Like</button>
                </div>
                <a class="text-primary me-2 pt-1" href="">
                    {{this.post.likes_count}} likes
                </a>
                <div class="text-primary pt-1">
                    {{this.post.comments_count}} comments
                </div>
                <div class="small text-muted ms-auto pt-1">{{this.getReadableDate()}}</div>
            </div>
                <div>
                    <div class="row d-flex justify-content-center mt-100 mb-100">
                        <div >
                                <div class="">
                                    <Comment v-for="item in commenttest"
                                            :key="item.id"
                                            :username="item.username"
                                            :comment_text="item.text"
                                            :date="item.date">
                                            
                                    </Comment>
                                </div> 
                        </div>
                </div>
            </div>
        
            <div class="input-group mb-3 pt-3">
                <input type="text" class="form-control bg-dark border-white" placeholder="Comment" aria-label="Comment" aria-describedby="basic-addon2">
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary border-white " type="button">Comment</button>
                </div>
            </div>
        </div>
    </div>
    </div>
</template>