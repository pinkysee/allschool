<template>
    <div class="leftmenu">
        <a class="leftmenulink" href="#">Новости</a>
        <a class="leftmenulink" href="#">Адрес</a>
        <a class="leftmenulink" href="#">Учителя</a>
        <a class="leftmenulink" href="#">Контакты</a>
        <a class="leftmenulink" href="#">Мероприятия</a>
        <a class="leftmenulink" href="#">Пед.состав</a>
        <a class="leftmenulink" href="#">Для учащихся</a>
        <a class="leftmenulink" href="#">lorem</a>
        <a class="leftmenulink" href="#">lorem</a>
        <a class="leftmenulink" href="#">Cтарая весия сайта</a>
    </div>
</template>
<script>
export default {

}
</script>
<style scoped>
.leftmenu {
    display: flex;
    flex-direction: column;
    padding: var(--padding-bar);
}

.leftmenulink {
    margin-top: 5px;
    padding: var(--apadding);
}

.leftmenulink:hover {
    background-color: var(--bgcolorhover);
}
a {
    text-decoration: none;
    color: #fff;
    padding: 10px;
    border-radius: 12px;
}
</style>