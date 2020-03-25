<template>
    <div class="row">
        <div v-bind:key=" link.id
    " v-for="link in links.items">
            <LinkCard v-bind:link="link"
                      v-on:del-link="$emit('del-link', link.id)"/>
        </div>
    </div>
</template>

<script>
    import LinkCard from './LinkCard.vue';

    export default {
        name: "LinkPlato",
        components: {
            LinkCard
        },
        props: ["links"]
    }
</script>

<style scoped>
</style>