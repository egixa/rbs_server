export {}
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
    }

    response = await fetch(url + "?root=" + path + "/" + linkContent + "&sort=" + sortButton.textContent, { 
        method: 'GET',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' 
        }
    }).then((response) => {
        if (!response.ok) { 
            throw new Error('Error occurred!')
        } 
        return response.json()
    }).catch((err)=>{
            console.log(err)
    })
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
    let data = await getTableFolder(link);
    createTable(data);
}
