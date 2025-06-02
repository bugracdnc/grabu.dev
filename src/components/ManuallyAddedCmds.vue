<script lang="ts" setup>
import { type Command } from './CmdLines.vue'

const manualCommands: Command[] = []

manualCommands.push({
    comment: 'source code',
    command: 'curl',
    updated_at: undefined,
    link: 'https://github.com/bugracdnc/grabu.dev',
    text: 'bugracdnc/grabu.dev'
})
manualCommands.push({
    comment: 'main website',
    command: 'curl',
    updated_at: undefined,
    link: 'https://grabu.dev',
    text: 'grabu.dev'
})
manualCommands.push({
    comment: 'my cv in english',
    command: 'curl',
    updated_at: undefined,
    link: 'cv/en',
    text: 'grabu.dev/cv/en'
})
manualCommands.push({
    comment: 'türkçe özgeçmişim',
    command: 'curl',
    updated_at: undefined,
    link: 'cv/tr',
    text: 'grabu.dev/cv/tr'
})
</script>
<template>
    <div v-for="command in manualCommands" :key="command.command">
        <span class="comment"># {{ command.comment }}</span>
        <span class="PS1">$ ></span>
        {{ command.command }}
        <a v-bind:href="command.link">{{ command.text }}</a
        ><br /><br />
    </div>
</template>
