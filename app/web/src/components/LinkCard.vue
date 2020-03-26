<template>
    <div class="card">

        <!-- Named link -->
        <div v-if="link.name !== ''" class="card-body">
            <h5 class="card-title">
                <a v-bind:href=link.address target="_blank">
                    {{ link.name | truncate(64, '...')}}
                </a>
            </h5>
            <a v-bind:href=link.address class="card-link" target="_blank">
                {{ link.address | truncate(64, '...')}}
            </a>
            <i class="icon-paper-clip card-link"
               v-on:click="showCopiedMessage()"
               v-clipboard="link.address">
            </i>
            <i class="icon-trash card-link"
               @click="$emit('del-link', link.id)"></i>
        </div>

        <!-- Unnamed link -->
        <div v-else class="card-body">
            <h5 class="card-title">
                <a v-bind:href=link.address target="_blank">
                    {{link.address | truncate(64, '...')}}
                </a>
            </h5>
            <i class="icon-paper-clip card-link"
               v-on:click="showCopiedMessage()"
               v-clipboard="link.address">
            </i>
            <i class="icon-trash card-link"
               @click="$emit('del-link', link.id)">
            </i>
        </div>
    </div>
</template>

<!--https://codesandbox.io/s/9oqml4vyyr-->

<script>
    export default {
        name: "LinkCard",
        props: [
            "link"
        ],
        methods: {
            showCopiedMessage() {
                let toast = this.$toasted.show("Copied!", {
                    theme: "outline",
                    position: "top-center",
                    duration: 1000
                });
            }
        }
    }
</script>

<style scoped>
</style>