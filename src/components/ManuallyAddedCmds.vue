<script setup lang="ts">

import { type Command } from './CmdLines.vue';

var manualCommands: Command[] = []

manualCommands.push({
    comment: 'source code',
    command: 'curl',
    link: 'https://github.com/bugracdnc/grabu.dev',
    text: 'bugracdnc/grabu.dev'
})
manualCommands.push({
    comment: 'main website',
    command: 'curl',
    link: 'https://grabu.dev',
    text: 'grabu.dev'
})
manualCommands.push({
    comment: 'my cv in english',
    command: 'curl',
    link: 'cv/en.html',
    text: 'grabu.dev/cv/en'
})
manualCommands.push({
    comment: 'türkçe özgeçmişim',
    command: 'curl',
    link: 'cv/tr.html',
    text: 'grabu.dev/cv/tr'
})

</script>
<template>
    <div v-for="command in manualCommands" :key="command.command">
        <span class="comment"># {{ command.comment }}</span>
        <span class="PS1">$ ></span>
        {{ command.command }}
        <a v-bind:href="command.link">{{ command.text }}</a><br /><br />
    </div>
</template>