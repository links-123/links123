<template>
    <div class="col-sm-6 col-md-4 col-lg-4">
        <div class="media card p-30">
            <div class="media-body media-text-left">
                <a v-bind:href=link.address target="_blank">
                    <h2>{{ link.name }}</h2>
                </a>
                <a v-bind:href=link.address target="_blank"
                   class="m-b-0">
                    {{ link.address }}
                </a>
            </div>

        </div>
    </div>
</template>

<script>
    export default {
        name: "LinkCard",
        props: [
            "link"
        ]
    }
</script>

<style scoped>
</style>