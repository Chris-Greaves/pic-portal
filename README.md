# pic-portal

A simple web-app using [Uppy](https://uppy.io/) and connecting to a [Tus](https://tus.io) server to provide fast, reliable, bulk uploads of images and video.

## Why does this exist

I'm getting married, and with all modern weddings, there are lots of photos that the attendees will have taken on their phone. A simple portal to privately share these images to the bride and groom, without sharing it with anyone else, is exactly what I need! Add in [Uppy's](https://uppy.io) for upload queuing and [Tus's](https://tus.io) ability to make resumable uploads, and you have something robust enough that even the in-laws can use it.

## Features

- [x] Ability to upload multiple files at once
- [x] Handle reuploading files more than once, without taking up more space
- [x] Can remove files from the queue
- [ ] Can limit the uploads to specific file extensions
- [ ] Works on IOS, Android and Desktop OSs
- [ ] Is securable using a password
- [ ] A tool for taking the Tus data and turn it back into the files on the server

## Developing

installed dependencies with `npm install` (or `pnpm install` or `yarn`), 

Then you can start a development server:

```sh
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```sh
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.

#### Note

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).