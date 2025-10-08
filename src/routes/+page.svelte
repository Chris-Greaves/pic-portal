<script lang="ts">
import Uppy from '@uppy/core'
import { type DashboardOptions } from '@uppy/dashboard'
import Dashboard from '@uppy/svelte/dashboard'
import Tus from '@uppy/tus'
import Compressor from '@uppy/compressor';

import '@uppy/svelte/css/style.css'
import '@uppy/dashboard/css/style.min.css';

const uppy = new Uppy()
  .use(Tus, {
    //endpoint: 'https://tusd.tusdemo.net/files/',
    endpoint: 'http://192.168.8.131:8080/files/',
  })
  .use(Compressor)

uppy.setOptions({
	restrictions: { allowedFileTypes: ["image/*", "video/*"] },
	autoProceed: true,
});

const dashboardOptions = {
  note: "Images won't be shared on Social Media without the permission of all seen in the photo",
  proudlyDisplayPoweredByUppy: false
} satisfies Partial<DashboardOptions<any, any>>

let nameValue: string = "";

function updateMetadata() {
  uppy.setMeta({uploader_name: nameValue})
}
</script>

<main class="p-5 max-w-xl mx-auto">
  <h1 class="text-4xl text-center font-bold my-4">Thank you for attending!</h1>

  <h3 class="text-l text-center my-4">If you wish to share your photos with the bride and groom, please provide your name, and upload the photos using the section below:</h3>

  <div>
      <!-- <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Name</label> -->
      <input placeholder="Your name..." bind:value={nameValue} on:change={updateMetadata} type="text" id="name" class="name-field text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 my-4" required />
  </div>
  <Dashboard uppy = {uppy} props = {dashboardOptions}/>
</main>

<style>
  body {
    font-family: ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  }
  .name-field {
    background-color: #f4f4f4;
    color: #000000;
    font-weight: 400;
  }
</style>