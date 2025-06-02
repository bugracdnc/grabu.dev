<script lang="ts">
import { type Command } from './CmdLines.vue'

export default {
    data() {
        const emptyData: Command[] = []
        return { githubRepos: emptyData }
    },
    beforeCreate() {
        const fetchedRepos: Command[] = []
        const blacklist: Array<String> = ['bugracdnc/bugracdnc', 'bugracdnc/grabu.dev']
        fetch('https://api.github.com/users/bugracdnc/repos').then((response) => {
            response.json().then((data) => {
                for (let repo of data) {
                    if (blacklist.includes(repo.full_name)) {
                        continue
                    }
                    fetchedRepos.push({
                        comment: repo.description,
                        command: 'curl',
                        updated_at: new Date(Date.parse(repo.updated_at)),
                        link: repo.html_url,
                        text: repo.full_name
                    })
                }
                this.githubRepos = fetchedRepos.sort((a, b) =>
                    a.updated_at! > b.updated_at! ? -1 : b.updated_at! > a.updated_at! ? 1 : 0
                )
                console.log(this.githubRepos)
            })
        })
    }
}
</script>
<template>
    <hr />
    <br />
    <span class="comment"># Github Repos</span><br />
    <div v-for="command in githubRepos" :key="command.command">
        <span class="comment"># {{ command.comment }}</span>
        <span class="PS1">$ ></span>
        {{ command.command }}
        <a v-bind:href="command.link">{{ command.text }}</a>
        <span class="date">{{
            ` (${command.updated_at!.getDate().toString().padStart(2, '0')}-${(command.updated_at!.getMonth() + 1).toString().padStart(2, '0')}-${command.updated_at!.getFullYear()} ${command.updated_at!.getHours().toString().padStart(2, '0')}:${command.updated_at!.getMinutes().toString().padStart(2, '0')})`
        }}</span>
        <br /><br />
    </div>
</template>
