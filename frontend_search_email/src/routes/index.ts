import EmailsRecords from "@/components/pages/EmailsRecords.vue"
import EmailContent from "@/components/pages/EmailContent.vue"

export const routes = [
  { path: '/', component: EmailsRecords},
  { path: '/:pathMatch(.*)*', component:EmailsRecords },
  { path: '/content/:messageId', component: EmailContent},

]
