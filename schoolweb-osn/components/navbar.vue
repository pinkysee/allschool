<template>
    <header class="header">
        <a href="#" class="homelink"><span>ШК</span>ола 102</a>
        <div id="menuspan" @click="menuactive()">
            <span class="menuspanspan"></span>
            <span class="menuspanspan"></span>
            <span class="menuspanspan"></span>
            <div class="menu">
                
                <div id="aboutschool" @click="aboutschoolactive()" class="infoaboutschool">Сведения об образовательной
                    организации
                 <div>   <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                <path fill="none" d="M0 0h24v24H0z"></path>
                <path fill="currentColor" d="M12 13.172l4.95-4.95 1.414 1.414L12 16 5.636 9.636 7.05 8.222z"></path>
            </svg></div>
                    <span class="iaslist">
                        <a href="./info-about-school/aboutschool" class="infoaboutschoollink">Основные сведения</a>
                        <a href="./info-about-school/managment" class="infoaboutschoollink">Структура и органы управления
                            образовательной организацией</a>
                        <a href="./info-about-school/documents" class="infoaboutschoollink">Документы</a>
                        <a href="./info-about-school/education" class="infoaboutschoollink">Образование</a>
                        <a href="./info-about-school/stuff-list" class="infoaboutschoollink">Руководство. Педагогический
                            состав</a>
                        <a href="./info-about-school/equipment" class="infoaboutschoollink">Материально-техническое
                            обеспечение
                            и оснащенность образовательного процесса</a>
                        <a href="./info-about-school/paid-services" class="infoaboutschoollink">Платные образовательные
                            услуги</a>
                        <a href="./info-about-school/economic-activites" class="infoaboutschoollink">Финансово-хозяйственная
                            деятельность</a>
                        <a href="./info-about-school/transfer" class="infoaboutschoollink">Вакантные места для приема
                            (перевода)
                            обучающихся</a>
                        <a href="./info-about-school/enviroment" class="infoaboutschoollink">Доступная среда</a>
                        <a href="./info-about-school/international" class="infoaboutschoollink">Международное
                            сотрудничество</a>
                    </span>
                </div>
                <a href="#" class="menulink">Приём в школу</a>
                <a href="/disteducation" class="menulink">Дистанционное обучение</a>
                <a href="#" class="menulink">Наставничество</a>
            </div>
        </div>
        <div class="user-menu-not-auth" v-if="Auth == false">
            <a href="#" class="munelink">Войти</a>
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
                <a class="userautha">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                        <path fill="none" d="M0 0h24v24H0z"></path>
                        <path fill="currentColor"
                            d="M20 22h-2v-2a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v2H4v-2a5 5 0 0 1 5-5h6a5 5 0 0 1 5 5v2zm-8-9a6 6 0 1 1 0-12 6 6 0 0 1 0 12zm0-2a4 4 0 1 0 0-8 4 4 0 0 0 0 8z">
                        </path>
                    </svg>
                    <span>Профиль</span>
                </a>
                <a class="userautha">Не придумал</a>
                <a class="userautha">
                    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M21 18.5C21 19.881 19.881 21 18.5 21H5.5C4.1195 21 3 19.881 3 18.5V5.5C3 4.119 4.1195 3 5.5 3H18.5C19.881 3 21 4.119 21 5.5V18.5Z"
                            fill="#1976D2" />
                        <path
                            d="M17.9685 9.0205C17.9915 8.945 18.0025 8.875 17.9995 8.8125C17.992 8.6315 17.8675 8.5 17.5745 8.5H16.2655C15.935 8.5 15.7825 8.7 15.6935 8.9005C15.6935 8.9005 14.8775 10.58 13.937 11.6875C13.632 12.008 13.477 12 13.312 12C13.2235 12 13 11.893 13 11.5995V9.007C13 8.66 12.9135 8.5 12.634 8.5H10.3095C10.106 8.5 10 8.66 10 8.8205C10 9.154 10.449 9.234 10.5 10.1685V11.98C10.5 12.42 10.4235 12.5 10.2585 12.5C9.8135 12.5 8.9375 11 8.351 9.034C8.224 8.647 8.097 8.5 7.7665 8.5H6.445C6.0635 8.5 6 8.687 6 8.887C6 9.2475 6.3 11.1965 7.9375 13.4375C9.125 15.0625 10.6895 16 12.0745 16C12.9135 16 12.9995 15.7865 12.9995 15.453V13.967C13 13.5665 13.0915 13.5 13.3585 13.5C13.549 13.5 13.9375 13.625 14.6875 14.5C15.5525 15.509 15.7095 16 16.2055 16H17.5145C17.8185 16 17.993 15.8725 18 15.625C18.0015 15.562 17.9925 15.4915 17.972 15.413C17.875 15.125 17.43 14.421 16.875 13.75C16.5675 13.3785 16.264 13.0105 16.1245 12.8105C16.031 12.68 15.9955 12.588 16 12.5C16.0045 12.4075 16.0525 12.3195 16.1245 12.1965C16.1115 12.1965 17.8035 9.821 17.9685 9.0205Z"
                            fill="white" />
                    </svg>
                    <span>Интеграция с Вконтакте</span>
                </a>
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
export default {
    data() {
        return {
            Auth: true,
            User: { Name: "Артур К." },
        }
    },
    methods: {
        checkuser() {
            if (!localStorage.getItem("Bearer")) { return false }
        },
        menuactive() {
            document.getElementById("menuspan").classList.toggle("active");
        },
        aboutschoolactive() {
            document.getElementById("aboutschool").classList.toggle("activeaboutschool")
        },
        usermenu() {
            document.getElementById("user-menu-auth").classList.toggle("usnavactive");
            document.getElementsByTagName("body").classList.toggle("close")
            console.log(document.getElementsByTagName("body"), "ffff")
        }
    },

    mounted() {
    },
}
</script>
<style scoped>

.usnavactive .usnav {
    display: flex;
}

.usnavactive {
    background-color: var(--bgcolorhover);
}

.usnav {
    position: absolute;
    width: 270px;
    top: 50px;
    left: -140px;
    display: none;
    flex-direction: column;
    background-color: var(--otbgcolor);
    border-radius: 12px;
    padding: 7px;
}

.userautha {
    height: 30px;
    margin-top: 5px;
    border-radius: 12px;
}

.strel {
    margin-top: 9px;
}

.userautha:hover {
    background-color: var(--bgcolordefault);
}

.userautha:nth-child(3) {
    background-color: #1976d2;
}

.userautha:nth-child(3) svg {
    margin: 0 0 0 0;
}

.userautha:nth-child(3) span {
    position: relative;
    top: -10px;
}

.userautha span {
    position: relative;
    top: -5px;
    margin: 0 0 0 5px;
}

.user-menu-auth {
    padding: 2px;
    margin-top: 5px;
    display: flex;
    color: #fff;
    border-radius: 6px;
}

.user-menu-auth:hover {
    background-color: var(--bgcolorhover);
    cursor: pointer;
}

.user-menu-auth p {
    margin: 10px 5px 0 5px;
    font-weight: 700;
}

.avatar {
    width: 32px;
    height: 32px;
    overflow: hidden;
    border-radius: 6px;
    margin: 5px 0 0 0;
    background-color: var(--bgcolorhover);
}

.menu {
    display: flex;
    padding: var(--padding-bar);
}

.homelink {

    font-size: 24px;
    font-weight: 700;
    padding: var(--apadding);
}

.homelink span {
    color: blueviolet;
}

.header {
    position: relative;
    display: grid;
    grid-template-columns: auto 90% 10%;
}

.menulink {
    margin: 0px 0px 0px 5px;
    background-color: var(--bgcolordefault);
}

.menulink:hover {
    background-color: var(--bgcolorhover);
}


.message {
    padding: var(--padding-bar);
    border-radius: 12px;
}

.message svg {
    height: 24px;
    width: 24px;
}

.message:hover {
    background-color: var(--bgcolorhover);
}

.auth {
    padding: var(--padding-bar);
}

.user-menu-auth {

    top: 10px;
    display: flex;
    right: 0;
    position: absolute;
}


a {
    text-decoration: none;
    color: #fff;
    padding: 10px;
    border-radius: 12px;
}

.infoaboutschool {
    color: white;
    display: flex;
    padding: var(--padding-bar);
    border-radius: 12px;
    position: relative;
    background-color: var(--bgcolorhover);
}

.iaslist {
    display: none;
}

.infoaboutschool:hover .iaslist {
    display: grid;
    grid-template-columns: 50% 50%;
    gap: 5px;
    position: absolute;
    bottom: -350px;
    left: 0;
    width: 1000px;
    background-color: var(--otbgcolor);
    padding: var(--padding-bar);
    border-radius: 12px;
}

.infoaboutschool .iaslist a {
    background-color: var(--bgcolorhover);
    padding: 10px;
    text-align: center;
    height: 30px;
}
.infoaboutschool .iaslist a:hover {
   background-color: var(--dopbgcolor);
}

.infoaboutschool span a:last-child {
    grid-column: 1/3;
}

@media (max-width: 1270px) {
    #menuspan {
        position: relative;
    }
    body.close {
        overflow: hidden;
    }
    .menuspanspan {
        display: block;
        margin-top: 10px;
        background-color: #fff;
        width: 30px;
        height: 2px;
        transition: all 0.3s ease 0s;
    }

    .menuspanspan:nth-child(2) {
        transition: all 0.1s ease 0s;
    }

    .active .menuspanspan {
        position: absolute;
    }

    .active .menuspanspan:nth-child(1) {
        transform: rotate(45deg);
        top: 12px;
    }

    .active .menuspanspan:nth-child(2) {
        transition: all 0s linear 0s;
        transform: scale(0);
    }

    .active .menuspanspan:nth-child(3) {
        top: 12px;
        transform: rotate(-45deg);
    }

    .header {
        grid-template-columns: 20% 60% 20%;
    }

    .user-menu-auth {
        display: grid;
        grid-template-columns: auto auto auto;
    }

    .user-menu-auth svg {
        grid-row: 1/2;
        grid-column: 3/4;
    }

    .user-menu-auth p {
        grid-row: 1/2;
    }

    .user-menu-auth div {
        grid-row: 1/2;
        grid-column: 1/2;
    }

    .homelink {
        text-align: center;
        grid-column: 2/3;
    }

    .menuspan {
        grid-column: 3/4;
    }

    .user-menu-not-auth {
        top: 5px;
        grid-column: 1/2;
    }

    .usnav {
        left: 0;
    }

    .user-menu-auth {
        top: 5px;
        grid-column: 1/2;
    }
}

@media (max-width: 1270px) and (min-width: 567px) {
    .menu {
        display: none;
    }

    .active .menu {
        width: 100vw;
        position: absolute;
        height: 100vh;
        top: 45px;
        padding: 0;
        left: -644px;
        display: flex;
        flex-direction: column;
        background-color: #1976d2;
    }

    .menulink {
        border-radius: 6px;
        background-color: #202020;
    }

}

@media (max-width: 566px) {
    .menu {
        display: none;
    }

    .active .menu {
        display: flex;
        flex-direction: column;
        padding: 0;
        position: absolute;
        width: 100vw;
        right: -5px;
        top: 50px;
    }

    .menulink {
        background-color: var(--bgcolorhover);
        margin: 5px 0 0 0;
        text-align: center;
    }

    .activeaboutschool {
        font-size: 16px;
        padding: 0;
        background-color: var(--bgcolorhover);
    }

    .activeaboutschool span {
        margin-top: 15px;
        display: flex;
        flex-direction: column;
        background-color: var(--dopbgcolor);
    }

    .infoaboutschool .iaslist a {
        height: auto;
        display: flex;
        margin: 5px 0 0 0;
        text-align: center;
        justify-content: center;
        align-items: center;
    }
}

@media (max-width: 737px) {
    .user-menu-auth p {
        display: none;
    }

    .user-menu-auth {
        margin: 0 0 0 0;
    }
}</style>