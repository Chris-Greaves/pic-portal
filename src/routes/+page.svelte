<script lang="ts">
import Uppy from '@uppy/core'
import {
  Dropzone,
  FilesList,
  FilesGrid,
  UploadButton,
  UppyContextProvider,
} from '@uppy/svelte'
import Tus from '@uppy/tus'
import Compressor from '@uppy/compressor';

import '@uppy/svelte/css/style.css'

const uppy = new Uppy()
  .use(Tus, {
    //endpoint: 'https://tusd.tusdemo.net/files/',
    endpoint: 'http://192.168.8.131:8080/files/',
  })
  .use(Compressor)

  let dialogRef: HTMLDialogElement
</script>

<UppyContextProvider {uppy}>
  <main class="p-5 max-w-xl mx-auto">
    <h1 class="text-4xl font-bold my-4">Welcome to the pic-portal!</h1>

    <p>Select your files and click upload</p>

    <Dropzone />
    <UploadButton />

    <dialog
      bind:this={dialogRef}
      class="backdrop:bg-gray-500/50 rounded-lg shadow-xl p-0 fixed inset-0 m-auto"
    >
    </dialog>

    <article>
      <h2 class="text-2xl my-4">Files</h2>
      <FilesList />
    </article>

    <article>
      <h2 class="text-2xl my-4">Thumbnails</h2>
      <FilesGrid columns={3} />
    </article>
  </main>
</UppyContextProvider>
