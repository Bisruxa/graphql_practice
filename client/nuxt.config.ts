import { defineNuxtConfig } from "nuxt/config";

export default defineNuxtConfig({
  modules: ["@nuxtjs/apollo"],

  apollo: {
    clients: {
      default: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
        httpLinkOptions: {
          headers: {
            "x-hasura-admin-secret": "mysecret", // replace with your actual admin secret
          },
        },
      },
    },
  },

  compatibilityDate: "2025-03-21",
});
