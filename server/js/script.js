import {clickGet} from "./get.js"

const url = "http://localhost:3007/";
const Asc = "asc";
const Desc = "desc";
const path = "";

function splitPath(path){
    const parts = path.split('/');
    parts.pop();
    path = parts.join('/');
    return path
}

// Кнопка сортировки
const sortButton = document.querySelector('.sort-button');
let sortAscending = false;
let AskDesk = Asc;   

//Обработка нажатия на кнопку сортировки
sortButton.addEventListener('click', async() => {
    sortAscending = !sortAscending;
    sortButton.textContent = sortAscending ? Asc : Desc ;
    if (sortAscending == true){
        sortButton.textContent =  Asc;
    }
    else {sortButton.textContent = Desc};
    path = document.querySelector('.path').innerHTML;

    const response = await fetch(url + "?root=" + path + "&sort=" + sortButton.textContent, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' 
        }
    });
    const data = await response.json(); // Получаем JSON-данные
    createTable(data);
    })

    //Обработка нажатия на кнопку Назад
const backButton = document.querySelector('.back-button');
backButton.addEventListener('click', async () => {
    const link = "Назад";
    clickGet(link);

    path = document.querySelector('.path').innerHTML;
    if (path === "/"){
        //
    }
    else{
        document.querySelector('.path').innerHTML = splitPath(path);
    }
})

// Функция для начальной загрузки корневой директории при старте приложения
async function loadInitialDirectory() {
    path = document.querySelector('.path').innerHTML; // Получаем значение корневой директории из тега div
    const response = await fetch(url + "?root=" + path + "&sort=" + AskDesk, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    });
    const data = await response.json(); // Получаем JSON-данные
    createTable(data); // Отображаем данные на странице
}

// Запускаем начальную загрузку при старте приложения
loadInitialDirectory();
