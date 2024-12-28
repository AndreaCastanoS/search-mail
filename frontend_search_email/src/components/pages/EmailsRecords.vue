<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { RouterLink } from 'vue-router'

interface Email {
  content: string
  from: string
  subject: string
  to: string
  message_id: string
}

const emails = ref<Email[]>([])
const searchQuery = ref<string>('')

const filteredEmails = computed(() => {
  if (!searchQuery.value.trim()) return emails.value
  return emails.value.filter((email) =>
    [email.to, email.from, email.subject, email.content, email.message_id]
      .join(' ')
      .toLowerCase()
      .includes(searchQuery.value.toLowerCase())
  )
})

function truncateContent(content: string): string {
  return content.length > 70 ? `${content.slice(0, 70)}...` : content
}

async function loadEmails() {
  try {
      const response = await fetch('http://localhost:3333/search', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      }
    })

    if (response.ok) {
      const data = await response.json()
      emails.value = data.results
    } else if (response.status === 401) {
      console.log('Credenciales inválidas')
    }
  } catch (err) {
    console.error('Error en la solicitud:', err)
  }
}

onMounted(() => {
  loadEmails()
})
</script>

<template>
  <div class="max-w-4xl mx-auto my-8 font-sans text-gray-800">
    <div class="flex justify-between items-center mb-5">
      <h1 class="text-2xl text-blue-600">Email Search</h1>
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search emails"
        class="p-3 border border-gray-300 rounded-md text-lg"
      />
    </div>
    <div class="bg-white p-4 rounded-lg shadow-md">
      <div class="overflow-x-auto">
        <table class="w-full table-auto border-collapse">
          <thead>
            <tr class="bg-gray-100">
              <th class="py-3 px-4 text-left text-gray-600">To</th>
              <th class="py-3 px-4 text-left text-gray-600">From</th>
              <th class="py-3 px-4 text-left text-gray-600">Subject</th>
              <th class="py-3 px-4 text-left text-gray-600">Content</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="email in filteredEmails" :key="email.message_id" class="even:bg-gray-50 hover:bg-gray-100">
              <td class="py-3 px-4">
                <RouterLink class="text-black hover:text-blue-800" :to="`/content/${encodeURIComponent(email.message_id)}`">{{ truncateContent(email.to) }}</RouterLink>
              </td>
              <td class="py-3 px-4">
                <RouterLink class="text-black hover:text-blue-800" :to="`/content/${encodeURIComponent(email.message_id)}`">{{ truncateContent(email.from) }}</RouterLink>
              </td>
              <td class="py-3 px-4">
                <RouterLink class="text-black hover:text-blue-800" :to="`/content/${encodeURIComponent(email.message_id)}`">{{ truncateContent(email.subject) }}</RouterLink>
              </td>
              <td class="py-3 px-4">
                <RouterLink class="text-black hover:text-blue-800" :to="`/content/${encodeURIComponent(email.message_id)}`">{{ truncateContent(email.content) }}</RouterLink>
              </td>
            </tr>
            
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
