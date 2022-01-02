<template>
    <div class="button">
        <button @click="onClick">
            <span class="material-icons-round">{{ role }}</span>
        </button>
    </div>
</template>
<script>
    export default {
        props: {
            role: {
                validator: function (value) {
                    return ['add', 'delete'].indexOf(value) !== -1;
                }
            },
        },
        methods: {
            onClick() {
                this.$emit(this.role);
            }
        }
    }
</script>
<style scoped>
    button {
        position: absolute;
        right: 28rem;
        bottom: 24rem;
        width: 55rem;
        height: 55rem;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--emphasis-color);
        background-color: var(--background-primary);
        border: none;
        border-radius: 50%;
        box-shadow: 0 0 5rem 1rem rgba(0, 0, 0, 0.1);
    }

    button span {
        font-size: 38rem;
    }

    button:hover {
        cursor: pointer;
    }
</style>