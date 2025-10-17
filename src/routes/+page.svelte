<script lang="ts">
import Uppy from '@uppy/core'
import { type DashboardOptions } from '@uppy/dashboard'
import Dashboard from '@uppy/svelte/dashboard'
import DashboardModal from '@uppy/svelte/dashboard-modal'
import Tus from '@uppy/tus'
import Compressor from '@uppy/compressor';

import '@uppy/svelte/css/style.css'
import '@uppy/dashboard/css/style.min.css';

import { env } from '$env/dynamic/public';

const tusPort = env.PUBLIC_TUS_PORT != undefined ? `:${env.PUBLIC_TUS_PORT}` : ""

const uppy = new Uppy()
  .use(Tus, {
    endpoint: `${env.PUBLIC_TUS_SCHEMA}://${env.PUBLIC_TUS_HOST}${tusPort}/files/`,
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
let allowUpload: boolean = false;

function nameFieldChanged() {
  uppy.setMeta({uploaded_by: nameValue})
  allowUpload = nameValue !== ""
}
</script>

<main class="p-5 max-w-xl mx-auto">
  <h1 class="text-4xl text-center font-bold my-4">Thank you for attending!</h1>

  <h3 class="text-l text-center my-4">If you wish to share your photos with the bride and groom, please provide your name below, and a modal should pop up allowing you to upload as many photos as you like</h3>

  <div>
      <!-- <label for="name" class="block mb-2 text-sm font-medium text-gray-900">Name</label> -->
      <input placeholder="Your name..." bind:value={nameValue} on:change={nameFieldChanged} type="text" id="name" class="name-field text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 my-4" required />
  </div>
  <!-- <Dashboard uppy = {uppy} props = {dashboardOptions}/> -->
  <DashboardModal uppy = {uppy} props = {dashboardOptions} open = {allowUpload}/>
</main>

<style>
  main {
    font-family: ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  }
  .name-field {
    background-color: #f4f4f4;
    color: #000000;
    font-weight: 400;
  }
</style>