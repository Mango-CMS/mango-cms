import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { apolloClient } from './graphql/client'
import { DefaultApolloClient } from '@vue/apollo-composable'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// 初始化Apollo Client
app.provide(DefaultApolloClient, apolloClient)

app.mount('#app')
