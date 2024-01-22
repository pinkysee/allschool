<template>
    <nv />
    <div class="main">
        <leftmenu />
        <div class="content">

            <div class="blockcont">
                <div class="block">
                </div>

                <div class="block">
                
                </div>
                <div class="block">
                   
                </div>
                <div class="block">4</div>
                <div class="block"><iframe
                        src="https://yandex.ru/map-widget/v1/?um=constructor%3Abdeeae2c7732a99677f7113ad7733e3fe02a71b957e73cc416b40ba168cbef5f&amp;source=constructor"
                        width="100%" height="100%" id="id04" frameborder="0"></iframe></div>
                <div class="block"><img src="../assets/schoolphoto.jpg" width="100%" height="100%"></div>
            </div>
            <h1>Новости школы</h1>
            <div class="all-news-container">
                <div class="gl-news-container">
                    <h2>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Fugit, unde.</h2>
                    <div class="gl-news-img"></div>
                    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita quidem esse, atque corrupti
                        delectus maiores! Corrupti quam sapiente, eum consequuntur explicabo dolorem quas aliquid laborum
                        harum nisi voluptatum culpa voluptas quae expedita autem. Quam, quo! Voluptatibus, esse similique
                        dolor fuga ab eligendi sit, ipsam nostrum repudiandae, quasi nobis aliquid dignissimos.</p>
                </div>
                <div class="news-container">
                    <div class="news" v-for="k in News">
                        <div class="news-img"></div>
                        <div class="news-content">
                            <div class="calendar-container">
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                    xmlns="http://www.w3.org/2000/svg">
                                    <path
                                        d="M8.03135 14.0439C8.60095 14.0439 9.0627 13.5864 9.0627 13.022C9.0627 12.4575 8.60095 12 8.03135 12C7.46175 12 7 12.4575 7 13.022C7 13.5864 7.46175 14.0439 8.03135 14.0439Z"
                                        fill="#2C2C2C" />
                                    <path
                                        d="M12 14C12.5523 14 13 13.5523 13 13C13 12.4477 12.5523 12 12 12C11.4477 12 11 12.4477 11 13C11 13.5523 11.4477 14 12 14Z"
                                        fill="#2C2C2C" />
                                    <path
                                        d="M16.0314 14.0439C16.601 14.0439 17.0627 13.5864 17.0627 13.022C17.0627 12.4575 16.601 12 16.0314 12C15.4618 12 15 12.4575 15 13.022C15 13.5864 15.4618 14.0439 16.0314 14.0439Z"
                                        fill="#2C2C2C" />
                                    <path
                                        d="M18.013 5H5.98698C4.33732 5 3 6.29322 3 7.8885V18.1115C3 19.7068 4.33732 21 5.98698 21H18.013C19.6627 21 21 19.7068 21 18.1115V7.8885C21 6.29322 19.6627 5 18.013 5Z"
                                        stroke="#2C2C2C" stroke-width="2" stroke-miterlimit="10" />
                                    <path d="M8 3V6" stroke="#2C2C2C" stroke-width="2" stroke-miterlimit="10"
                                        stroke-linecap="round" />
                                    <path d="M16 3V6" stroke="#2C2C2C" stroke-width="2" stroke-miterlimit="10"
                                        stroke-linecap="round" />
                                    <path d="M3 9H21" stroke="#2C2C2C" stroke-width="1.5" stroke-miterlimit="10"
                                        stroke-linecap="round" />
                                    <path
                                        d="M8.03135 18.0439C8.60095 18.0439 9.0627 17.5864 9.0627 17.022C9.0627 16.4575 8.60095 16 8.03135 16C7.46175 16 7 16.4575 7 17.022C7 17.5864 7.46175 18.0439 8.03135 18.0439Z"
                                        fill="#2C2C2C" />
                                    <path
                                        d="M12.0314 18.0439C12.601 18.0439 13.0627 17.5864 13.0627 17.022C13.0627 16.4575 12.601 16 12.0314 16C11.4618 16 11 16.4575 11 17.022C11 17.5864 11.4618 18.0439 12.0314 18.0439Z"
                                        fill="#2C2C2C" />
                                </svg>
                                <p>{{ k.Date }}</p>
                            </div>
                            <h3>{{ k.Title }}</h3>
                        </div>
                    </div>
                </div>
            </div>
            <h1>Мероприятия школы</h1>
            <div class="mp"></div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Auth: () => {
                if (!localStorage.getItem("Bearer")) { return false }
            },
            News: [{ "Title": "Lorem ipsum dolor sit amet consectetur", "Text": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita quidem es", "Preview": "", "Date": "26.07.24" },
            { "Title": "Lorem ipsum dolor sit amet consectetur", "Text": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita quidem es", "Preview": "", "Date": "26.07.24" },
            { "Title": "Lorem ipsum dolor sit amet consectetur", "Text": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita quidem es", "Preview": "", "Date": "26.07.24" }
            ],
        }
    },
    methods: {
        checkuser() {
            if (!localStorage.getItem("Bearer")) { return false }
        }
    },
    mounted() {
    }
}
</script>
<style scoped>
.blockcont {
    display: grid;
    width: 80%;
    grid-template-columns: repeat(4, 300px);
    grid-gap: 10px;
    margin: 0 auto 0 auto;
    grid-template-rows: 300px 300px;
}

.block {
    background-color: var(--bgcolorhover);
    width: 100%;
    height: 100%;
    border-radius: 12px;

}

.block:nth-child(4) {
    grid-column: 4;
    grid-row: 1/3;
}

.block:nth-child(5) iframe {
    border-radius: 12px;
}

.block:nth-child(6) {
    grid-row: 2;
    grid-column: 2/4;
}

.block:nth-child(6) img {
    border-radius: 12px;
}



:root {
    --padding-bar: 10px;
    --apadding: 10px;
    --bgcolorhover: #2c2c2c;
    --bgcolordefault: #171717;
    --textcolor: #fff;
}


.main {
    display: grid;
    grid-template-columns: 13% auto;
}


.content {
    color: var(--textcolor);
}

.content h1 {
    text-align: center;
}

.all-news-container {
    width: 1350px;

    margin: auto;
    display: grid;
    height: 600px;
    grid-template-columns: 550px  800px;
}

.gl-news-container {
    width: 100%;
    height: 100%;
}

.gl-news-img {
    width: 550px;
    height: 300px;
    background-color: grey;
    border-radius: 12px;
}

.news-container {
    margin-left: 30px;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
}

.news {
    margin-bottom: 15px;
    display: grid;
    grid-template-columns: 400px 400px;
    width: 100%;
    height: 33%;
}

.news h3 {
    color: #747474;
    font-size: 16px;
}

.news-img {
    width: 100%;
    border-radius: 6px;
    background-color: gray;
    height: 100%;
}

.news-content {
    max-width: 70%;
    height: 100%;
    margin-left: 5px;
}

* {
    padding: 0;
}


a {
    text-decoration: none;
    color: #fff;
    padding: 10px;
    border-radius: 12px;
}
@media (max-width: 1607px) {
    .leftmenu  {
        display: none;
    }
    .main  {
        grid-template-columns: auto;
    }
    .blockcont {
        margin: 0 0 0 20px;
    }
    .all-news-container{margin: 0 0 0 20px;width: 1240px;overflow: hidden;}
}

</style>