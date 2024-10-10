<script lang="ts">
import { type Command } from './CmdLines.vue';

export default {
    data() {
        var emptyData: Command[] = []
        return { githubRepos: emptyData }
    },
    beforeCreate() {
        var fetchedRepos: Command[] = []
        var blacklist: Array<String> = ["bugracdnc/bugracdnc", "bugracdnc/grabu.dev"]
        fetch('https://api.github.com/users/bugracdnc/repos')
            .then(response => {
                response.json().then(data => {
                    for (let repo of data) {
                        if (blacklist.includes(repo.full_name)) { continue; }
                        fetchedRepos.push({
                            comment: repo.description,
                            command: 'curl',
                            link: repo.html_url,
                            text: repo.full_name
                        });
                    }
                    this.githubRepos = fetchedRepos
                })
            })
    }
};
</script>
<template>
    <hr /><br />
    <span class="comment"># Github Repos</span><br>
    <div v-for="command in githubRepos" :key="command.command">
        <span class="comment"># {{ command.comment }}</span>
        <span class="PS1">$ ></span>
        {{ command.command }}
        <a v-bind:href="command.link">{{ command.text }}</a><br /><br />
    </div>
</template>