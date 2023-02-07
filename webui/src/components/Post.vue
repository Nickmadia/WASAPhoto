<script>
import Comment from './Comment.vue'

export default {
    props: ['post', 'userId', 'userName'],
    components:{
        Comment
    },
    data() {
        return {
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
        getReadableDate(date) {
            var d = new Date(date)
            return d.toDateString() + ' ' + d.toLocaleTimeString()
        },
        getCurrentDate() {
            var d = new Date()
            
            return d.toDateString() + ' ' + d.toLocaleTimeString()
        },
        getb64() {
            return 'data:image/png;base64,' + this.img
        },
        async like(){
            try {
            if(!this.liked){
                let res = await this.$axios.put('/posts/' + this.post.id +'/likes/' + this.userId)
                this.likes_count += 1
                
            } else {
                let res = await this.$axios.delete('/posts/' + this.post.id +'/likes/' + this.userId)
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
                    let res = await this.$axios.post('/posts/'+ this.post.id+ '/comments/'+ this.userId, body)
                    if (res.status == 201) {
                        if (this.post.comments != null) {
                            this.post.comments.push({"comment_id":res.data,"owner_id":this.userId,"owner_username":this.userName,"comment_text":this.current_comment,"time_stamp":this.getCurrentDate()})
                        } else {
                            this.post.comments = [{"comment_id":res.data,"owner_id":this.userId,"owner_username":this.userName,"comment_text":this.current_comment,"time_stamp":this.getCurrentDate()}]
                        }
                        this.comments_count++
                        this.current_comment = ""

                        //check owner in backend
                    }
                    
                } catch (e) {

                }
            }
        }, async uncommentPost(id) {  
            try {
                let res = await this.$axios.delete('/posts/'+ this.post.id+ '/comments/comment/'+ id)
                if (res.status == 204) {
                    
                    this.post.comments = this.post.comments.filter(x => x.comment_id != id)
                }
                this.comments_count--
            } catch (e) {

            }
            
        },
        isOwner(id) {
            return id == this.userId
        }
    },
    async beforeMount() {
        
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
            <div class="d-flex mb-1">
                <h4 class="card-title text-primary fw-bold "> {{this.user.username}}</h4>
                <button v-if="this.post.owner_id == this.userId" @click="$emit('delPost', this.post.id)" class="ms-auto btn btn-balck btn-sm btn-outline-danger ">delete</button>
            </div>
            <div class="imagecontainer ">
                <img :src="'data:image/png;base64,' + this.img" class="card-img-top" alt="Card image cap" >
            </div>
            <div class="d-flex border-bottom pb-2">
                <div class="pb-2 me-2">
                    <button v-if="!this.liked" class="btn btn-primary btn-sm " @click="like">Like</button>
                    <button v-else class="btn btn-black btn-outline-secondary text-white btn-sm " @click="like">Like</button>
                </div>
                <a class="text-primary me-2 pt-1" href="">
                    {{this.likes_count}} likes
                </a>
                <div class="text-primary pt-1">
                    {{this.comments_count}} comments
                </div>
                <div class="small text-muted ms-auto pt-1">{{this.getReadableDate(this.post.time_stamp)}}</div>
            </div>
                <div>
                    <div class="row d-flex justify-content-center mt-100 mb-100">
                        <div >
                                <div class="">
                                    <Comment v-for="item in post.comments"
                                            :key="item.comment_id"
                                            :comment_id="item.comment_id"
                                            :username="item.owner_username"
                                            :comment_text="item.comment_text"
                                            :owner="this.isOwner(item.owner_id)"
                                            :date="this.getReadableDate(item.time_stamp)"
                                            @removeComment="uncommentPost">
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