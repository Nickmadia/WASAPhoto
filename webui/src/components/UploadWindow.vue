<script>
export default {
    methods: {
        async upload() {
            try {
                const file = this.$refs.file.files[0]
                const reader = new FileReader()
                reader.addEventListener('load', async (event) => {
                    this.img = event.target.result
                    
                    var body = this.convertToBase64(this.img);   
              
                    const response = await this.$axios.post('/media',body)
                  
                })
                reader.readAsArrayBuffer(file)
                
                
                    
                   
            }
            catch(e) {

            }
        },
        convertToBase64(array){
            var u8 = new Uint8Array(array)
            var bin = ''
            for (var i =0; i< u8.length ; i++){
                bin += String.fromCharCode(u8[i])
            }
                console.log('wtf')
            return window.btoa(bin)
        }
    },
    data () {
        return {
            img: null
        }
    }
}
</script>


<template>
    <div class="modal fade" id="uploadWindow" tabindex="-1" aria-hidden="true">
				<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
					<div class="modal-content">
						<div class="modal-header">
							<h5 class="modal-title " >Upload</h5>
							<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
						</div>
						<div class="modal-body">
							<div class="row d-flex justify-content-center align-items-center ">
                                <div class="d-flex">
								    <input class="from-control" accepts=".png .jpg .jpeg" type="file" ref="file" >
                                    <button class="btn btn-primary ms-auto " @click="upload">post</button>
                                </div>
							</div>
						</div>
					
					</div>
				</div>
			</div>
</template>