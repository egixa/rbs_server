const url = "http://localhost:3007/";
const Asc = "asc"
const Desc = "desc"

function splitPath(path){
    const parts = path.split('/');
    parts.pop();
    path = parts.join('/');
    return path
}
//GET-запрос на получение содержимого директории
async function getTableFolder(link){
    path = document.querySelector('.path').innerHTML
    var linkContent;
    if (link === "Назад") {
        path = splitPath(path);
        linkContent = "";
    }
    else {
        linkContent = link.textContent;
        // if (path === "/") {
        //     path += link.textContent;
        // } else {
        //     path += "/" + link.textContent;
        // }
    }

    response = await fetch(url + "?root=" + path + "/" + linkContent + "&sort=" + sortButton.textContent, { // Запрос на сервер
        method: 'GET',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' 
        }
    }).then(response => response.json()).catch();
    return response
}

// Создание таблицы с данными из GET-запроса
function createTable(data){
    path = document.querySelector('.path').innerHTML

    // Очищаем таблицу
    fileTableBody.innerHTML = '';
     // Заполняем таблицу данными
     data.forEach(item => {
        const row = fileTableBody.insertRow();
        const typeCell = row.insertCell();
        const nameCell = row.insertCell();
        const sizeCell = row.insertCell();

        // Добавляем ссылку для папок
        if (item.type === "Директория") {
            const link = document.createElement('a');
            link.href = "#"; // Замените на реальный URL, если необходимо
            link.textContent = item.name;
            link.addEventListener('click', async (event) => {
                clickGet(link)
                
                if (path === "/"){
                    document.querySelector('.path').innerHTML = path + link.textContent;
                }
                else{
                    document.querySelector('.path').innerHTML = path + "/" + link.textContent;
                }
                });
            nameCell.appendChild(link);
        } else {
            nameCell.textContent = item.name;
          }
          typeCell.textContent = item.type;
          sizeCell.textContent = item.formatingSize;
      });
}

// Нажатие на ссылку в директорию
async function clickGet(link){
    data = await getTableFolder(link);
    createTable(data);
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