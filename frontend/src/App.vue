<script setup lang="ts">
const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'

const socket = new WebSocket(`${protocol}://${window.location.host}/ws`)

socket.addEventListener('open', () => {
  console.log('WebSocket connected')

  socket.send(
    JSON.stringify({
      type: 'hello',
      message: 'Hello from Vue',
    }),
  )
})

socket.addEventListener('message', (event) => {
  console.log('Message from server:', event.data)
})

socket.addEventListener('close', (event) => {
  console.log('WebSocket closed:', event.code, event.reason)
})

socket.addEventListener('error', (event) => {
  console.error('WebSocket error:', event)
})
</script>

<template>
  <h1>You did it!</h1>
  <p>
    Visit <a href="https://vuejs.org/" target="_blank" rel="noopener">vuejs.org</a> to read the
    documentation
  </p>
  <p>print stuff:</p>
</template>

<style scoped></style>
