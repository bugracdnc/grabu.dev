<script setup lang="ts">
import TextIcon from './components/TextIcon.vue'
import Info from './components/PersonalInfo.vue'
import Commands from './components/CmdLines.vue'
</script>

<template>
  <span class="PS1">$ > </span>echo "$grabu_dev"
  <br /><br />
  <TextIcon />
  <br />
  <Info />
  <br />
  <Commands />
</template>
