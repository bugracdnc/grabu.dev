<script setup lang="ts">
interface Command {
  comment: string
  command: string
  link: string
  text: string
}

var commands: Command[] = []
commands.push({
  comment: 'source code',
  command: 'curl',
  link: 'https://github.com/bugracdnc/grabu.dev',
  text: 'github.com/bugracdnc/grabu.dev'
})
commands.push({
  comment: 'main website',
  command: 'curl',
  link: 'https://grabu.dev',
  text: 'grabu.dev'
})
commands.push({
  comment: 'my cv in english',
  command: 'curl',
  link: 'cv/en',
  text: 'grabu.dev/cv/en'
})
commands.push({
  comment: 'türkçe özgeçmişim',
  command: 'curl',
  link: 'cv/tr',
  text: 'grabu.dev/cv/tr'
})
</script>

<template>
  <div v-for="command in commands" :key="command.command">
    <span class="comment"># {{ command.comment }}</span>
    <span class="PS1">$</span>
    {{ command.command }}
    <a v-bind:href="command.link">{{ command.text }}</a
    ><br /><br />
  </div>
</template>

<style lang="css" scoped>
a {
  margin-bottom: 1em;
}

.comment {
  display: block;
}

.PS1 {
  color: #1a1;
}

.comment {
  color: #444;
}
</style>
