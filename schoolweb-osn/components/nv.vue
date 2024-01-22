<template>
    <header class="header">
        <a :href="`/${basePath}`" class="logo"><span>Шк</span>ола 102</a>
        <div id="burger" class="burger" @click="burger()">
            <span class="burgerspan"></span>
        </div>
        <nav class="osnmenu">
            <ul>
                <li>
                    <div id="aboutschool" class="aboutschool" @click="aboutschoolactive">
                        <div>
                            <p>Сведения об образовательной организации</p>
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24
                        24" width="24" height="24">
                                <path fill="none" d="M0 0h24v24H0z"></path>
                                <path fill="currentColor"
                                    d="M12 13.172l4.95-4.95 1.414 1.414L12 16 5.636 9.636 7.05 8.222z">
                                </path>
                            </svg>
                        </div>
                        <span class="aboutschoolspan">
                            <a :href="`${basePath}/info-about-school/aboutschool`" class="infoaboutschoollink"
                                exact>Основные сведения</a>
                            <a :href="`${basePath}/info-about-school/managment`" class="infoaboutschoollink" exact>Структура
                                и органы
                                управления
                                образовательной организацией</a>
                            <a :href="`${basePath}/info-about-school/documents`" class="infoaboutschoollink">Документы</a>
                            <a :href="`${basePath}${basePath}/info-about-school/education`"
                                class="infoaboutschoollink">Образование</a>
                            <a :href="`${basePath}/info-about-school/stuff-list`" class="infoaboutschoollink">Руководство.
                                Педагогический
                                состав</a>
                            <a :href="`${basePath}/info-about-school/equipment`"
                                class="infoaboutschoollink">Материально-техническое
                                обеспечение
                                и оснащенность образовательного процесса</a>
                            <a :href="`${basePath}/info-about-school/paid-services`" class="infoaboutschoollink">Платные
                                образовательные
                                услуги</a>
                            <a :href="`${basePath}/info-about-school/economic-activites`"
                                class="infoaboutschoollink">Финансово-хозяйственная
                                деятельность</a>
                            <a :href="`${basePath}/info-about-school/transfer`" class="infoaboutschoollink">Вакантные места
                                для приема
                                (перевода)
                                обучающихся</a>
                            <a :href="`${basePath}/info-about-school/enviroment`" class="infoaboutschoollink">Доступная
                                среда</a>
                            <a :href="`${basePath}/info-about-school/international`"
                                class="infoaboutschoollink">Международное
                                сотрудничество</a>
                        </span>
                    </div>
                </li>
                <li>
                    <a href="#" class="navlink">Студенты</a>
                </li>
                <li>
                    <a href="#" class="navlink">Студенты</a>
                </li>
                <li>
                    <a href="#" class="navlink">Студенты</a>
                </li>
            </ul>
        </nav>
        <div class="user-menu-not-auth" v-if="Auth == false">
            <a :href="`${basePath}/login`" class="munelink">Войти</a>
        </div>
        <div id="user-menu-auth" class="user-menu-auth" v-else @click="usermenu()">
            <!--             <div class="message">
                <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 32 32" fill="none">
                    <path
                        d="M15.9999 32C14.5855 32 13.2289 31.4381 12.2287 30.4379C11.2285 29.4377 10.6666 28.0812 10.6666 26.6667C10.6666 26.3131 10.8071 25.9739 11.0571 25.7239C11.3072 25.4738 11.6463 25.3333 12 25.3333H19.9999C20.3535 25.3333 20.6927 25.4738 20.9427 25.7239C21.1928 25.9739 21.3332 26.3131 21.3332 26.6667C21.3332 28.0812 20.7713 29.4377 19.7711 30.4379C18.771 31.4381 17.4144 32 15.9999 32ZM13.6933 28C13.9277 28.4041 14.2642 28.7395 14.669 28.9727C15.0738 29.2059 15.5328 29.3286 15.9999 29.3286C16.4671 29.3286 16.9261 29.2059 17.3309 28.9727C17.7357 28.7395 18.0722 28.4041 18.3066 28H13.6933Z"
                        fill="white" />
                    <path
                        d="M30.6665 28H1.33333C0.979709 28 0.640573 27.8595 0.390525 27.6095C0.140478 27.3594 2.81474e-06 27.0203 2.81474e-06 26.6667C-0.00116651 25.1609 0.36201 23.6772 1.05854 22.3421C1.75506 21.0071 2.76426 19.8605 3.99998 19V12C3.99991 8.83121 5.2532 5.79103 7.48635 3.54288C9.71951 1.29473 12.7512 0.0211252 15.9199 0H16.0666C19.2492 0 22.3014 1.26428 24.5518 3.51472C26.8023 5.76516 28.0665 8.8174 28.0665 12V19C29.2899 19.8678 30.2862 21.0177 30.9709 22.3523C31.6556 23.6868 32.0085 25.1667 31.9998 26.6667C31.9998 27.0203 31.8594 27.3594 31.6093 27.6095C31.3593 27.8595 31.0201 28 30.6665 28ZM2.79999 25.3333H29.1999C29.011 24.409 28.6281 23.5352 28.0765 22.7698C27.5249 22.0044 26.8171 21.3648 25.9999 20.8933C25.7964 20.7759 25.6276 20.6067 25.5105 20.403C25.3934 20.1993 25.3323 19.9683 25.3332 19.7333V12C25.3332 9.52465 24.3499 7.15068 22.5996 5.40034C20.8492 3.65 18.4753 2.66667 15.9999 2.66667C13.5246 2.66667 11.1506 3.65 9.40029 5.40034C7.64996 7.15068 6.66664 9.52465 6.66664 12V19.7333C6.66757 19.9683 6.6064 20.1993 6.48933 20.403C6.37226 20.6067 6.20344 20.7759 5.99997 20.8933C5.18279 21.3648 4.47498 22.0044 3.92335 22.7698C3.37173 23.5352 2.98883 24.409 2.79999 25.3333Z"
                        fill="white" />
                </svg>
            </div> -->
            <svg class="strel" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                <path fill="none" d="M0 0h24v24H0z"></path>
                <path fill="currentColor" d="M12 13.172l4.95-4.95 1.414 1.414L12 16 5.636 9.636 7.05 8.222z"></path>
            </svg>
            <p>{{ User.Name }}</p>
            <div class="avatar"></div>
            <nav class="usnav">
                <a class="userautha" :href="`${basePath}/profile`">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                        <path fill="none" d="M0 0h24v24H0z"></path>
                        <path fill="currentColor"
                            d="M20 22h-2v-2a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v2H4v-2a5 5 0 0 1 5-5h6a5 5 0 0 1 5 5v2zm-8-9a6 6 0 1 1 0-12 6 6 0 0 1 0 12zm0-2a4 4 0 1 0 0-8 4 4 0 0 0 0 8z">
                        </path>
                    </svg>
                    <span>Профиль</span>
                </a>
                <a class="userautha">Не придумал</a>
                <a class="userautha"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"
                        fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
                        <path
                            d="M18.189 9a15 15 0 0 1 2.654 2.556c.105.13.157.287.157.444m-2.811 3a14.998 14.998 0 0 0 2.654-2.556A.704.704 0 0 0 21 12m0 0H8m5-7.472A6 6 0 0 0 3 9v6a6 6 0 0 0 10 4.472">
                        </path>
                    </svg>
                    <span>Выйти</span>
                </a>
            </nav>
        </div>
    </header>
</template>
<script>
import axios from 'axios';
export default {
    data() {
        return {
            Auth: false,
            User: { Name: "Артур К." },
        }
    },
    computed: {
        basePath() {
            // Получите базовый путь из конфигурации Nuxt.js
            return this.$config.basePath || '';
        }
    },
    methods: {
        async checkuser() {
            if (!localStorage.getItem("Bearer")) { return false }
            let a = await axios.get(this.$config.public.API_URL + "/checkuser", {
                headers: {
                    'Authorization': `${localStorage.getItem("Bearer")}`
                }
            })
                .then(response => {
                    if (response.status != 202) {return false }
                    this.User = response.data
                    this.User.Name = response.data.Name.split(" ")[0] + " " + response.data.Name.split(" ")[1][0] + "."
                    return true
                })
                .catch(error => {
                    localStorage.removeItem("Bearer");
                });
            return a
            },
        menuactive() {
            document.getElementById("menuspan").classList.toggle("active");
        },
        aboutschoolactive() {
            document.getElementById("aboutschool").classList.toggle("activeaboutschool")
        },
        usermenu() {
            document.getElementById("user-menu-auth").classList.toggle("usnavactive");
            document.getElementsByTagName("body")[0].classList.toggle("close")
            setTimeout(() => {
                document.getElementsByTagName("body")[0].classList.toggle("close")
            }, 300);
        },
        burger() {
            document.getElementById("burger").classList.toggle("burgeractive");
            document.getElementsByTagName("body")[0].classList.toggle("close")
        }
    },

    async mounted() {
        this.Auth = await this.checkuser()
    },
}
</script>
<style>
.aboutschool div {
    display: flex;
}

.logo {
    font-size: 24px;
    font-weight: 900;
    margin-top: 10px;
}

.logo span {
    color: blueviolet;
}

.userautha {
    padding: 10px;
    text-wrap: nowrap;
    display: flex;
    border-radius: 6px;
    margin-top: 5px;
}

.userautha:nth-child(3) {
    background-color: #1976d2;

}

.userautha span {
    margin: 4px 0 0 5px;
}

.userautha:hover {
    background-color: var(--bgcolordefault);
}

header {
    display: flex;
}

a {
    text-decoration: none;
    color: white;
}

ul {
    display: flex;
}

li {
    list-style: none;
    margin-left: 5px;
}

.aboutschool {
    position: relative;
    display: flex;
    margin-top: -10px;
    margin-left: 10px;
    color: white;
    background-color: var(--bgcolorhover);
    padding: var(--apadding);
    border-radius: 12px;
}

.aboutschool:hover span {
    display: grid;
}

.aboutschool p {
    margin: 0px;
}

.aboutschoolspan {
    padding: var(--apadding);
    position: absolute;
    left: 0;
    z-index: 100;
    top: 45px;
    width: 200%;
    display: none;
    grid-template-columns: 50% 48%;
    border-radius: 12px;
    gap: 5px;
    background-color: var(--otbgcolor);
}

.infoaboutschoollink {
    text-align: center;
    height: 40px;
    background-color: var(--bgcolorhover);
    padding: 10px;
    border-radius: 12px;
}

.munelink {
    position: absolute;
    right: 15px;
    top: 8px;
    border-radius: 12px;
    padding: 10px;
}

.munelink:hover {
    background-color: var(--bgcolorhover);
}

.infoaboutschoollink:last-child {
    grid-column: 1/3;
}

.infoaboutschoollink:hover {
    background-color: var(--dopbgcolor);
}

.navlink {
    padding: var(--apadding);
    border-radius: 12px;
}

.navlink:hover {
    background-color: var(--bgcolorhover);
}

.avatar {
    width: 32px;
    height: 32px;
    border-radius: 6px;
    background-color: var(--bgcolorhover);
}

.user-menu-auth {
    position: absolute;
    right: 10px;
    display: flex;
    color: white;
    padding: var(--apadding);
    border-radius: 12px;
}

.user-menu-auth:hover {
    cursor: pointer;
}

.user-menu-auth svg {
    margin: 5px 5px 0 0;
}

.user-menu-auth p {
    margin: 8px 8px 0 0;
}

.usnav {
    overflow: hidden;
    padding: var(--apadding);
    width: 270px;
    display: none;
    flex-direction: column;
    border-radius: 12px;
    position: absolute;
    right: -400px;
    top: 60px;
    animation: usnavun 0.3s;
    background-color: var(--otbgcolor);
}

.usnavactive {
    background-color: var(--bgcolorhover);
}

.usnavactive .usnav {
    animation: usnav2 0.3s;
    animation-fill-mode: forwards;
    display: flex;
}

.usnavactive .usnav a {
    display: block;
}

* {
    padding: 0;
}

@keyframes usnav {
    0% {
        left: -410px;
    }

    100% {
        left: 0;
    }
}

@keyframes usnav2 {
    from {
        right: -470px;
    }

    to {
        right: 0px;
    }
}

@media (max-width: 830px) {
    .header {
        display: grid;

        position: relative;
        grid-template-columns: 10% 80% 10%;
    }

    .logo {
        grid-column: 2/3;
        grid-row: 1;
        text-align: center;
    }

    .user-menu-auth {
        display: grid;
        grid-template-columns: auto auto auto;
        grid-column: 1/2;
        left: 5px;
        background-color: #ffffff00;
    }

    .user-menu-auth p {
        grid-row: 1/2;
    }

    .usnav {
        display: none;
    }

    .usnavactive .usnav {
        display: flex;
        animation: usnav 0.3s ease-in-out;
        animation-fill-mode: forwards;
    }

    .user-menu-auth div {
        grid-row: 1/2;
        grid-column: 1/2;
    }

    .osnmenu {
        display: block;
        position: absolute;
        top: -10px;
        z-index: 2;
        left: -10px;
        padding: 0;
        width: 100vw;
        height: 0px;
        overflow: hidden;
        transition: all 0.3s;
        background-color: var(--bgcolorhover);
    }

    .osnmenu ul {
        display: none;
        margin-top: 50px;
    }

    .burgeractive~.osnmenu {
        height: 100vh;
        overflow: scroll;
    }

    .burgeractive~.osnmenu ul {
        display: block;
        flex-direction: column;
        padding: 0;

    }

    .burgeractive~.osnmenu li {
        width: 98%;
        overflow: hidden;
        margin: 30px 0 0 5px;
    }

    .navlink {
        display: inline-block;
        width: 100%;
    }

    .aboutschool {
        margin-top: 0px;
        width: 100%;
        background-color: #1976d200;

    }

    .burger::before,
    .burger::after,
    .burgerspan {
        content: "";
        position: absolute;
        background-color: #fff;
        width: 100%;
        height: 2px;
        left: 0;
        transition: all 0.3s ease-in-out 0s;
    }

    .burger::after {
        bottom: 0;
    }

    .burgeractive::after {
        position: absolute;
        transform: rotate(-45deg);
        top: 50%;

    }

    .burgeractive::before {
        top: 50%;
        position: absolute;
        transform: rotate(45deg);
    }

    .burgeractive .burgerspan {
        transform: scale(0);
    }

    .burgerspan {
        display: block;
        top: 9px;
    }


    .burger {
        position: relative;
        width: 25px;
        height: 20px;
        z-index: 10;
        margin-top: 30%;
        grid-column: 3/4;
    }

    .activeaboutschool {
        position: relative;
        flex-direction: column;
    }

    .aboutschool svg {
        position: absolute;
        right: 20px;
    }

    .activeaboutschool .aboutschoolspan {
        display: flex;
        flex-direction: column;
        position: static;
        width: 100vw;
        padding: 0;
        margin-left: -20px;
        gap: 0;
        background-color: var(--dopbgcolor);
        border-radius: 0px;
        height: auto;
    }

    .activeaboutschool .infoaboutschoollink {
        border-radius: 0px;
        width: 96%;
        background-color: #ffffff00;
    }

    .activeaboutschool {
        padding: 0;
    }
}

body.close {
    overflow: hidden;
    height: 100%;
}

@media (max-width: 510px) {
    .user-menu-auth p {
        display: none;
    }

}</style>