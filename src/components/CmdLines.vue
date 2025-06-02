<script setup lang="ts">

import FetchGithubRepos from './FetchGithubRepos.vue';
import ManuallyAddedCmds from './ManuallyAddedCmds.vue';

export interface Command {
  comment: string
  command: string
  updated_at: Date
  link: string
  text: string
}
</script>

<template>
  <ManuallyAddedCmds />
  <FetchGithubRepos />
</template>