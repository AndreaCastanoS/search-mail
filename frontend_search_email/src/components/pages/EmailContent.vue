<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

interface Email {
  content: string
  from: string
  subject: string
  to: string
  message_id: string
}

// Obtiene el parámetro de la ruta
const route = useRoute()
const message_id = route.params.messageId as string
console.log(message_id)

const email = ref<Email | null>(null)

async function loadEmailById(id: string) {
  try {
       const response = await fetch(`http://localhost:3333/email?message_id=${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      }
    })

    if (response.ok) {
      const data = await response.json()
      console.log(data)
      email.value = data
    } else {
      console.error(`Error al cargar el email con ID ${id}`)
    }
  } catch (err) {
    console.error('Error en la solicitud:', err)
  }
}

onMounted(() => {
  if (message_id) {
    loadEmailById(message_id)
  }
})
</script>

<template>
  <div class="max-w-3xl mx-auto my-6 p-6 bg-white rounded-lg shadow-md" v-if="email">
    <h1 class="text-2xl font-semibold text-blue-600 border-b pb-2 mb-4">{{ email.subject }}</h1>
    <div class="text-sm text-gray-600 space-y-2">
      <p><strong>From:</strong> {{ email.from }}</p>
      <p><strong>To:</strong> {{ email.to }}</p>
    </div>

    <div class="mt-6 text-base text-gray-800">
      <p>{{ email.content }}</p>
    </div>
  </div>
  <div v-else class="flex justify-center items-center h-64">
    <p class="text-xl text-gray-500">Cargando contenido del email...</p>
  </div>
</template>
