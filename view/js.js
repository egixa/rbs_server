const url = "http://localhost:3007/";

function getTableFolder(){
    const response = fetch(url + "?root=&sort=" + sortButton.textContent, { // Запрос на сервер
        method: 'GET',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' 
        }
    });
    const data = response.json(); // Получаем JSON-данные
    fileTableBody.innerHTML = '';
    return data

};

function createTable(data){
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
                });
            nameCell.appendChild(link);
        } else {
            nameCell.textContent = item.name;
          }
          typeCell.textContent = item.type;
          sizeCell.textContent = item.formatingSize;
      });
}

function clickGet(){
    data = getTableFolder()
    createTable(data)
}

const sortButton = document.querySelector('.sort-button');
    let sortAscending = false;
    let AskDesk = "asc";

    sortButton.addEventListener('click', async() => {
    sortAscending = !sortAscending;
    sortButton.textContent = sortAscending ? 'asc' : 'desc' ;
    if (sortAscending == true){
        sortButton.textContent =  "asc";
    }
    else {sortButton.textContent = "desc"};
    
    const response = await fetch(url + "?root=&sort=" + sortButton.textContent, { // Запрос на сервер
        method: 'GET',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' 
        }
    });
        
    const data = await response.json(); // Получаем JSON-данные
    
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
                clickGet()
            });
            nameCell.appendChild(link);
        }
        else {
            nameCell.textContent = item.name;
          }
          typeCell.textContent = item.type;
          sizeCell.textContent = item.formatingSize;
      })
    })