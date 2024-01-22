<template lang="">
    <nv />
    <div class="container">
        <div class="leftcontainer">
            <h1><span class="shspan">Шк</span>ола102</h1>
            <h2>Рады видеть вас снова!</h2>
            <div class="formcontainer">
                <div class="inp">
                <input v-model="Login" class="log" placeholder="Логин"><br>
                <input v-model="Password" class="pas" placeholder="Пароль" type="password">
                  </div>
                <div class="checkboxcont">
                    <input type="checkbox" id="cbx" style="display: none;">
                    <label for="cbx" class="check">
                        <svg width="18px" height="18px" viewBox="0 0 18 18">
                            <path d="M1,9 L1,3.5 C1,2 2,1 3.5,1 L14.5,1 C16,1 17,2 17,3.5 L17,14.5 C17,16 16,17 14.5,17 L3.5,17 C2,17 1,16 1,14.5 L1,9 Z"></path>
                            <polyline points="1 9 7 14 15 4"></polyline>
                        </svg>
                    </label>
                <span class="zpm">Запомнить меня</span>
                </div>
                <button @click="authuser()">Войти</button>
            </div>
        </div>
        <div class="rightcontainer">
            <div class="slider">
            <div class="slideractive"></div>
            <div class="slidernotactive centerslider"></div>
            <div class="slidernotactive"></div>
            </div>
        </div>
    </div>
    <footers /> 
</template>
<script>
import axios from 'axios';
export default {
    data() {
        return {
            Login: "",
            Password: "",
        }
    },
    methods: {
        authuser() {
            const jsonData = {
                Login: this.Login,
                Password: this.Password,
            }

            // Установите URL, куда вы хотите отправить запрос
            console.log(this.$config.public.API_URL)
            // Отправьте POST-запрос с данными в формате JSON
            axios.post(this.$config.public.API_URL+"/auth", jsonData, {
                headers: {
                    'Content-Type': 'application/json',
                },
            })
                .then(response => {
                    localStorage.setItem('Bearer', response.data.Token)
                    this.$router.replace('../profile')
                })
                .catch(error => {
                    console.error('Ошибка запроса:', error);
                });
        }
    },

    mounted() {
    },
}
</script>
<style scoped>
.slider {
    display: flex;
    position: absolute;
    bottom: 5%;
    left: 48%;
}

.centerslider {
    margin: 0 10px 0 10px;
}

.slideractive {
    background-color: #D9D9D9;
    width: 50px;
    height: 10px;
    border-radius: 12px;
}

.slidernotactive {
    background-color: #D9D9D9;
    width: 12px;
    height: 12px;
    border-radius: 50%;
}

.container {
    width: 100%;
    height: 93.1vh;
    display: grid;
    grid-template-columns: 50% 50%;
}

.checkboxcont {
    margin-top: 10px;
    position: relative;
}

.checkboxcont span {
    color: white;
    position: absolute;
    left: 30px;
}

button {
    width: 100%;
    height: 50px;
    transition: all 0.3s;
    margin-top: 48px;
    border: none;
    outline: none;
    background-color: blueviolet;
    border-radius: 12px;
    color: #fff;
    font-size: 20px;
}

button:hover {
    background-color: #00000000;
    box-shadow: 0 0 5px #ae00ff, 0 0 12px #001eff, 0 0 25px #ae00ff;
}

body {
    margin: 0;
}

.rightcontainer {
    height: 100%;
    position: relative;
    background-color: var(--bgcolorhover);
}

.leftcontainer {
    width: 100%;
}

h1,
h2 {
    text-align: center;
    color: white;
}

h1 {
    font-size: 48px;
}

h2 {
    font-size: 36px;
}

.shspan {
    color: blueviolet;
}

.formcontainer {
    width: 50%;
    margin-left: 25%;
    margin-top: 10%;
    height: 100%;

    height: auto;
    display: flex;
    flex-direction: column;
}

input {
    border: none;
    border-bottom: 1px solid #fff;
    color: #666161;
    height: 40px;
    width: 100%;
    background-color: #66616100;
    outline: none;
    font-family: montserrat;
}

.pas {
    margin-top: 28px;
}

.check {
    cursor: pointer;
    position: relative;
    margin: auto;
    width: 88px;
    height: 88px;
    -webkit-tap-highlight-color: transparent;
    transform: translate3d(0, 0, 0);
}

.check:before {
    content: "";
    position: absolute;
    top: -15px;
    left: -15px;
    width: 48px;
    height: 48px;
    border-radius: 50%;
    background: rgba(34, 50, 84, 0.03);
    opacity: 0;
    transition: opacity 0.2s ease;
}

.check svg {
    position: relative;
    z-index: 1;
    fill: none;
    stroke-linecap: round;
    stroke-linejoin: round;
    stroke: #c8ccd4;
    stroke-width: 1.5;
    transform: translate3d(0, 0, 0);
    transition: all 0.2s ease;
}

.check svg path {
    stroke-dasharray: 60;
    stroke-dashoffset: 0;
}

.check svg polyline {
    stroke-dasharray: 22;
    stroke-dashoffset: 66;
}

.check:hover:before {
    opacity: 1;
}

.check:hover svg {
    stroke: blueviolet;
}

#cbx:checked+.check svg {
    stroke: blueviolet;
}

#cbx:checked+.check svg path {
    stroke-dashoffset: 60;
    transition: all 0.3s linear;
}

#cbx:checked+.check svg polyline {
    stroke-dashoffset: 42;
    transition: all 0.2s linear;
    transition-delay: 0.15s;
}

@media (max-width: 1034px) {
    .leftcontainer h1 {
        display: none;
    }

    .rightcontainer {
        display: none;
    }

    .container {
        grid-template-columns: 100%;
    }
}</style>