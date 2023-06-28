<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Upload a certificate</h1>

        <form v-on:submit.prevent="makeWebsiteThumbnail">
          <div>
            <input type="file" ref="fileInput" />
            <button @click="uploadFile">Upload</button>
            <p v-if="uploadStatus">{{ uploadStatus }}</p>
          </div>
        </form>
        
        <TableComponent />
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import TableComponent from './components/TableComponent.vue';

export default {
  name: 'App',

  components: {
    TableComponent, // Register the table component
  },

  data() {
    return {
      uploadStatus: '',
    };
  },

  methods: {
    uploadFile() {
      const fileInput = this.$refs.fileInput;
      const file = fileInput.files[0]; // Get the selected file
      
      // Create a FormData object to send the file
      const formData = new FormData();
      formData.append('myFile', file);

      // Make the POST request using Axios
      axios
        .post('http://localhost:8081/cert', formData)
        .then(response => {
          // Handle the response
          console.log(response);
          this.uploadStatus = 'File uploaded successfully';
        })
        .catch(error => {
          // Handle any errors
          if (error.response) {
            // The request was made and the server responded with a status code
            // that falls out of the range of 2xx
            console.error('Server responded with an error:', error.response);
            this.uploadStatus = 'Error: Server responded with an error';
          } else if (error.request) {
            // The request was made but no response was received
            // (e.g., network error or server is not reachable)
            this.uploadStatus = 'Error: Unable to reach the server';
          } else {
            // Something happened in setting up the request
            console.error('Error:', error.message);
            this.uploadStatus = 'Error: Something went wrong';
          }
        });
    },
  },
};
</script>
