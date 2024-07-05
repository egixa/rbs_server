import clickGet from "./get.js"

const url = "http://localhost:3007/";
const Asc = "asc"
const Desc = "desc"

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
    // это мы не трогаем
    sortAscending = !sortAscending;
    sortButton.textContent = sortAscending ? Asc : Desc ;
    if (sortAscending == true){
        sortButton.textContent =  Asc;
    }
    else {sortButton.textContent = Desc};
    path = document.querySelector('.path').innerHTML

    const response = await fetch(url + "?root=" + path + "&sort=" + sortButton.textContent, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' 
        }
    });
    const data = await response.json(); // Получаем JSON-данные
    createTable(data);
    })

    //Обратоботка нажатия на кнопку Назад
const backButton = document.querySelector('.back-button');
backButton.addEventListener('click', async () => {
    const link = "Назад";
    clickGet(link);

    path = document.querySelector('.path').innerHTML
    if (path === "/"){
        //
    }
    else{
        document.querySelector('.path').innerHTML = splitPath(path);
    }
})