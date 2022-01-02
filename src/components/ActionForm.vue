<template>
    <form>
        <button class="close" @click="onClose">
            <span class="material-icons-round">close</span>
        </button>
        <input 
            :type="field.type" 
            :name="field.name" 
            v-model="data[field['name']]" 
            v-for="(field, index) in fields" 
            :key="index" 
            :placeholder="field.placeholder"
            :required="field.required"
        >
        <button type="button" class="submit" @click="onSubmit">Сохранить</button>
    </form>
</template>
<script>
    export default {
        props: {
            fields: {
                type: Array,
                required: true,
            }
        },
        data() {
            return {
                data: {},
            };
        },
        mounted() {
            this.createModels();
        },
        methods: {
            createModels() {
                this.fields.forEach(el => {
                    this.data[el['name']] = {};
                });
            },
            onSubmit() {
                this.$emit('submit', this.data);
            },
            onClose() {
                this.$emit('close');
            }
        },
    }
</script>
<style scoped>
    form {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 315rem;
        padding: 25rem;
        background-color: var(--background-primary);
        box-shadow: 0 0 5rem 1rem rgba(0, 0, 0, 0.1);
        border-radius: 10rem;
    }

    button {
        border: none;
    }

    .close {
        width: 25rem;
        height: 25rem;
        margin: 0 0 15rem auto;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: var(--emphasis-color);
        border-radius: 50rem;
    }

    .close span {
        display: flex;
        font-size: 15rem;
        line-height: 0.8;
        color: #fff;
    }

    input {
        width: 100%;
        height: 35rem;
        margin-top: 10rem;
        padding: 0 8rem;
        font-size: 15rem;
        color: #B2B2B2;
        background-color: #fff;
        border: none;
        border-radius: 5rem;
        box-shadow: 0 0 5rem 1rem rgba(0, 0, 0, 0.1);
    }

    input:first-of-type {
        margin-top: 0;
    }

    input::placeholder {
        font-family: inherit;
        font-size: inherit;
        color: inherit;
    }

    .submit {
        width: 100rem;
        height: 30rem;
        margin-top: 10rem;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 15rem;
        color: #fff;
        background-color: var(--background-tertiary);
        border-radius: 5rem;
        box-shadow: 0 0 5rem 1rem rgba(0, 0, 0, 0.1);
    }
</style>