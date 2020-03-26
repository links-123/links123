<template>
    <div>
        <div class="row">
            <div class="col-md-4 col-sm-12"></div>
            <div class="col-md-4 col-sm-12">
                <LinkPlato v-bind:links="links"
                           v-on:del-link="deleteLink"></LinkPlato>
            </div>
            <div class="col-md-4 col-sm-12"></div>
        </div>
    </div>
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
        methods: {
            deleteLink(id) {
                axios.delete(linksAPIAddress + '/links', {data: {id: id}})
                    .then(res => this.links.items = this.links.items.filter(link => link.id !== id))
                    .catch(err => console.log(err));
            }
        },
        created() {
            axios.get(linksAPIAddress + '/links')
                .then(res => this.links = res.data)
                .catch(err => console.log(err));
        }
    }
</script>

<style>
</style>