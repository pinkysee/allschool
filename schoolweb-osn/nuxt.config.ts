// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ['nuxt-swiper'],
  devtools: { enabled: false },
  runtimeConfig: {
    public: {
      API_URL: process.env.API_URL || "http://194.87.232.129:8080",
 
    },
  },
  app: {
    head: {
      link: [
        // <link rel="stylesheet" href="https://myawesome-lib.css">
        { rel: 'stylesheet', href: `${''}/allasset/root.css` }
      ],
      // please note that this is an area that is likely to change
    }
  },
})
