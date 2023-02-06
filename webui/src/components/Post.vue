<script>
import Comment from './Comment.vue'

export default {
    props: ['post', 'userId'],
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
                img: null,
                liked: false,
                likes_count: 0,
                comments_count:0,
                current_comment: ""

        }
    },
    methods : {
        async getImg() {
            try {
                let response = await this.$axios.get('/media/' + this.post.id)
                this.img = response.data
            } catch (e) {

            }
        },
        async getUser() {
            try {
                let response = await this.$axios.get('/users/' + this.post.owner_id)
                this.user = response.data
                
                
                
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
        },
        async like(){
            try {
            if(!this.liked){
                let res = await this.$axios.put('/media/' + this.post.id +'/likes/' + this.userId)
                this.likes_count += 1
                
            } else {
                let res = await this.$axios.delete('/media/' + this.post.id +'/likes/' + this.userId)
                this.likes_count -= 1
            }
            this.liked = !this.liked
            }
            catch(e) {

            }
        },
        hasLike() {
            if (this.post.likes != null) {
            for( let el of this.post.likes ){
                if(el.user_id == this.userId)
                {
                    this.liked = true
                }
            }
            }
        },
        getLikesCount() {
            if (this.post.likes != null) {
                this.likes_count = this.post.likes.length
            }
        
        },
        getCommentsCount() {
            if (this.post.comments != null) {
                this.comments_count = this.post.comments.length
            }
            
        },
        async commentPost() {
            if(this.current_comment!= '') {
                let body = { 'comment_text': this.current_comment}
                try {
                    let res = await this.$axios.post('/posts/'+ this.post.id+ '/comments/comment/'+ this.userId, body)
                    if (res.status == 204) {
                        this.post.comments.push({"comment_id":1,"owener_id":3,"owner_username":"nic","comment_text":this.current_comment,"time_stamp":"2023-02-05T21:38:07Z"})
                        this.current_comment = ""
                        //check owner in backend
                    }
                    
                } catch (e) {

                }
            }
        }
    },
    async beforeMount() {
        console.log(this.post)
        await this.getImg()
        await this.getUser()
        this.getLikesCount()
        this.getCommentsCount()
        this.hasLike()
        
    },
    computed : {
        
}
}
</script>
<template>
<div v-if="post!=null && user!=null">
    <div class="card bg-black border-white " style="width: 50rem; ">
        <div class="card-body text-white">
            <div class="d-flex">
                <h4 class="card-title text-primary fw-bold"> {{this.user.username}}</h4>
            </div>
            <div class="imagecontainer ">
                <img :src="'data:image/png;base64,' + this.img" class="card-img-top" alt="Card image cap" >
            </div>
            <div class="d-flex border-bottom pb-2">
                <div class="pb-2 me-2">
                    <button v-if="!this.liked" class="btn btn-primary btn-sm " @click="like">Like</button>
                    <button v-else class="btn btn-secondary btn-sm " @click="like">Like</button>
                </div>
                <a class="text-primary me-2 pt-1" href="">
                    {{this.likes_count}} likes
                </a>
                <div class="text-primary pt-1">
                    {{this.comments_count}} comments
                </div>
                <div class="small text-muted ms-auto pt-1">{{this.getReadableDate()}}</div>
            </div>
                <div>
                    <div class="row d-flex justify-content-center mt-100 mb-100">
                        <div >
                                <div class="">
                                    <Comment v-for="item in post.comments"
                                            :key="item.comment_id"
                                            :username="item.owner_username"
                                            :comment_text="item.comment_text"
                                            :date="getReadableDate(item.time_stamp)">
                                    </Comment>
                                </div> 
                        </div>
                </div>
            </div>
        
            <div class="input-group mb-3 pt-3">
                <input v-model="this.current_comment" type="text" class="form-control bg-dark border-white text-white" placeholder="Comment" aria-label="Comment" aria-describedby="basic-addon2">
                <div class="input-group-append">
                    <button @click="this.commentPost" class="btn btn-outline-secondary border-white " type="button">Comment</button>
                </div>
            </div>
        </div>
    </div>
    </div>
</template>