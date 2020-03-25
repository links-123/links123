<template>
    <LinkPlato v-bind:links="links"></LinkPlato>
</template>

<script>
    import LinkPlato from '../components/LinkPlato.vue';
    import axios from 'axios';

    const linksAPIAddress = 'http://127.0.0.1/api/v1';

    export default {
        components: {
            LinkPlato,
        },
        data() {
            return {
                links: []
            }
        },
        methods: {},
        created() {
            axios.get(linksAPIAddress + '/links')
                .then(res => this.links = res.data)
                .catch(err => console.log(err));
        }
    }
</script>

<style>
</style>