<script>

export default {
    data() {
        return {
            likes: 0,
            

        }
    }
}
</script>
<template>
    <div class="card bg" style="width: 20rem;">
        <div class="card-body">
            <h4 class="card-title"> username</h4>
            <div class="imagecontainer">
                <img src="./imgdb.jpeg" class="card-img-top" alt="Card image cap" >
            </div>
                <h6 class="card-subtitle text-muted"></h6>
                <div class="d-flex justify-content-between">
                    <a class="btn btn-outline-primary " href="javascript:">Like</a>
                    <span class="border border-primary ">
                        <p>10 people like this shit</p>
                    </span>
                </div>
            <div>
                <div class="row d-flex justify-content-center mt-100 mb-100">
    <div >
        <div class="card bg-light">
            <div class="card-body">
                <!-- Comment Row -->
                <div class="d-flex flex-row comment-row m-t-0">
                    
                    <div class="comment-text w-100">
                        <h6 class="font-medium">James Thomas</h6> <span class="m-b-15 d-block">This is awesome website. I would love to comeback again. </span>
                        <div class="comment-footer d-flex justify-content-between"> 
                            <span class="text-muted float-right">April 14, 2019</span>  <button type="button" class="btn btn-danger btn-sm">Delete</button>
                        </div>
                    </div>
                </div> 
                
                
            </div> 
        </div>
    </div>
</div>
            </div>
        
            <div class="input-group mb-3">
                <input type="text" class="form-control" placeholder="Comment" aria-label="Comment" aria-describedby="basic-addon2">
                <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button">Comment</button>
                </div>
            </div>
        </div>
    </div>
</template>