<template>
    <section class="home">
        <div class="container-fluid py-5">
            <h1 class="display-5 fw-bold">Choose a QUIZ</h1>
            <div class="form-check" v-for="(item) in quizzes" :key="item">
                <input class="form-check-input"
                       :value="item.name"
                       v-model="selectedOption"
                       type="radio"
                       :name="item.name"
                       :id="item.name">
                <label class="form-check-label" for={{item.name}}>
                    {{ item.title }}
                </label>
            </div>
            <button class="btn btn-primary btn-lg my-3"
                    type="button"
                    @click="passEvent()"
                    :disabled="selectedOption === undefined">
                Start!
            </button>
        </div>
    </section>
</template>

<script>
import axios from "@/services/axios";

export default {
    name: "Home",
    data() {
        return {
            quizzes: undefined,
            selectedOption: undefined
        }
    },
    created() {
        this.makeRequest()
    },
    methods: {
        async makeRequest() {
            try {
                const {data} = await axios.get(
                    `/quizzes/all`
                );
                this.quizzes = data;
                console.log(data);
            } catch (error) {
                console.error(error);
            }
        },
        passEvent() {
            this.$emit("startTheGame", {
                quizName: this.selectedOption,
            });
        },
    },
    watch: {
        selectedOption(value) {
            this.selectedOption = value
        },
    },
};
</script>