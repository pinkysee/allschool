<template>
    <nv />
    <div class="cont">
        <h1>{{ $route.query.class }} класс</h1>
        <div class="pluscont" @click="backcontvisible('Zav')"></div>

        <div class="usercont">
            <div>Имя</div>
            <div>Логин</div>
           
        </div>
        <div class="usercont" v-for="(item, index) in Userlist">
            <div>{{ item.Name }}</div>
            <div>{{ item.Login }}</div>
            <div class="delete" @click="DeleteUser(item.ID); Userlist.splice(index, 1)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M10 2L9 3H4V5H20V3H15L14 2H10ZM5 7V22H19V7H5ZM8 9H10V20H8V9ZM14 9H16V20H14V9Z"
                        fill="#7E2CBE" />
                </svg>
            </div>
        </div>
    </div>
    <div class="adduserback adduserbackclose" id="adduserback">
        <div class="adduser">
            <div @click="backcontvisible()" class="otcontclose"></div>
            <h1>Добавить учащиегося</h1>
            <div class="adduch" v-if="BackContType == 'reg'">
                <input v-model="AddUserInfo.Name" placeholder="Имя ученика">
                <input v-model="AddUserInfo.Login" placeholder="Логин">
                <button class="btn-add-docs" @click="RegisterUser('user')"><span></span>Добавить учителя</button>
            </div>
            <div class="reg_responde" v-if="BackContType == 'resp'">
                <h3>Логин: {{ RegInfo.Login }}</h3>
                <h3>Пароль: {{ RegInfo.Password }}</h3>
            </div>
        </div>
    </div>
</template>
<script>
import axios from "axios"
export default {
    data() {
        return {
            Userlist: [],
            AddUserInfo: {
                Name: "",
                Login: "",
                Role: "",
                ClassName: this.$route.query.class,
            },
            BackContType: "reg",
            RegInfo: "",
        }
    },
    methods: {
        async GetUser() {
            const a = await axios.post(this.$config.public.API_URL + "/getuserbyclass", {
                ClassName: this.$route.query.class
            }, {
                headers: {
                    "Authorization": localStorage.getItem("Bearer")
                }
            })
                .then(function (response) {
                    return response.data
                })
                .catch(function (err) { })
            return a
        },

        RegisterUser(role) {
            this.AddUserInfo.Role = role
            // Установите URL, куда вы хотите отправить запрос
            // Отправьте POST-запрос с данными в формате JSON
            axios.post(this.$config.public.API_URL + "/register", this.AddUserInfo, {
                headers: {
                    'Authorization': localStorage.getItem("Bearer"),
                    'Content-Type': 'application/json',
                },
            })
                .then(response => {
                    this.BackContType = "resp"
                    this.RegInfo = response.data

                })
                .catch(error => {
                    console.error('Ошибка запроса:', error);
                });
            this.Userlist.push(JSON.parse(JSON.stringify(this.AddUserInfo)))
        },
        async DeleteUser(id) {
            const jsonData = {
                ID: id,
            }
            let a = await axios.post(this.$config.public.API_URL + "/deluser", jsonData,{
                headers: {
                    'Authorization': `${localStorage.getItem("Bearer")}`,
                    'Content-Type': 'application/json',
                }
            }) .then(response => {
                })
                .catch(error => {
                    console.error('Ошибка запроса:', error);
                });
        },
        backcontvisible(type) {
            let bb = document.getElementById("adduserback")
            this.BackContType = "reg"
            bb.classList.toggle("adduserbackclose")
            this.AddUserInfo = {
                Name: "",
                Login: "",
                Role: "",

                ClassName: this.$route.query.class,
            }
        },
    },
    async mounted() {
        this.Userlist = await this.GetUser()
    }
}
</script>
<style scoped>

.btn-add-docs {
    background-color: var(--dopbgcolor);
    color: white;
    font-family: montserrat;
    border: none;
    width: 210px;
    position: relative;
    margin: auto;
    overflow: hidden;
    height: 50px;
    border-radius: 25px;
}

.btn-add-docs span {
    display: block;
    height: 100px;
    width: 100px;
    background-color: #7E2CBE;
    border-radius: 50%;
    position: absolute;
    top: -200px;
    left: -200px;
    transition: all 0.3s;
}

.btn-add-docs:hover span {
    top: -60px;
    left: -60px;
}

.otcontclose {
    margin: 15px;
    position: relative;
    height: 32px;
    width: 32px;
}

.adduch {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.adduch input {
    margin: auto;
    outline-style: none;
    width: 210px;
    background-color: var(--dopbgcolor);
    border: none;
    color: white;
    font-family: montserrat;
    height: 40px;
    border-bottom: 2px solid #7E2CBE;
}

.otcontclose::after,
.otcontclose::before {
    position: absolute;
    display: block;
    content: "";
    width: 100%;
    height: 2px;
    background-color: white;
    top: 12px;
    transform: rotate(-45deg);
    transition: all 0.3s;
}

.otcontclose::after {
    transform: rotate(45deg);
}

.adduserback {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    position: fixed;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.692);
}

.adduser {
    text-align: center;
    color: white;
    background-color: #2C2C2C;
    width: 500px;
    height: 500px;
    border-radius: 12px;
    margin: auto;
}

* {
    color: white;
}

.pluscont {
    height: 32px;
    width: 32px;
    position: relative;
    /*     right: 15px;
    top: 5px; */
    margin-bottom: 25px;
}

.pluscont::after {
    content: "";
    display: block;
    background-color: white;
    position: absolute;
    width: 100%;
    height: 2px;
    top: 15px;
}

.pluscont::before {
    content: "";
    background-color: white;
    display: block;
    position: absolute;
    width: 2px;
    height: 100%;
    left: 15px;
}

.cont {
    width: 100%;
}

h1 {
    text-align: center;
}

.usercont {
    /*     border-bottom: 1px solid gray;
 */
    height: 35px;
    display: flex;
}

.usercont div:nth-child(1) {
    width: 750px;
    min-width: 300px;
}
.usercont div:nth-child(2) {
    width: 300px;
}
.adduserbackclose {
    display: none;
}</style>