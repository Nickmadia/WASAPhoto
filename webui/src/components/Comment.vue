<script>
    export default {
    props: ['comment_id','comment_text', 'username', 'date', 'owner'],

}
</script>
<template>
    <div class="border-bottom py-2">
        <div class="d-flex flex-row align-items-center commented-user">
            <h5 class="mr-2 text-primary">{{this.username}}</h5><span class="dot mb-1"></span><span class="mb-1 ms-2 text-muted sm ">{{date}}</span>
            <button v-if="this.owner" @click="$emit('removeComment', this.comment_id)" class="btn btn-outline-danger  ms-auto ">del</button>
        </div>
        <div class="comment-text-sm"><span>{{comment_text}}</span></div>
    </div>
  </template>