<template lang="">
    <nv />
    <div class="container">
        <a :href="this.$config.public.API_URL + '/docs/' + item.Path" v-for="item in Docs"><span></span>{{ item.Name }}</a>

    </div>
</template>
<script>
import axios from 'axios';

export default {
    data() {
        return {
            Docs: [{
                Name: "загрузка",
                Path: "vk.com"
            }],
        }
    },
    methods: {
        async GetDocs() {
            const jsonData = {
                Page: "Материально-техническое обеспечение и оснащенность образовательного процесса"
            }
            const a = await axios.post(this.$config.public.API_URL + "/getdocs", jsonData)
            .then(function(response) {
                return response.data
            })
            .catch(function(err) {})
            return a 
        }
    },
    async mounted() {
        this.Docs = await this.GetDocs()
    }
}
</script>
<style scoped>
    .container {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }
    a span {
        display: block;
        position: absolute;
        rotate: 45deg;
        left: -5px;
        top: 15px;
        height: 10px;
        width: 10px;
        background-color: blueviolet;
    }
    a {
        margin: 0px 0px 0px 10px;
        display: flex;
        position: relative;
        padding: 10px;
        overflow: hidden;
    }
</style>