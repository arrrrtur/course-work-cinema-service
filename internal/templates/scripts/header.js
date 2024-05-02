// Функция для создания и добавления футера в DOM
function createHeader() {
    const header = document.createElement('header');
    header.innerHTML = `
        <div class="px-3 py-2 text-bg-dark border-bottom">
            <div class="container">
                <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
                    <a href="/" class="d-flex align-items-center my-2 my-lg-0 me-lg-auto text-white text-decoration-none">
                        <svg class="bi me-2" width="40" height="32" role="img" aria-label="Bootstrap"><use xlink:href="#bootstrap"></use></svg>
                    </a>

                    <ul class="nav col-12 col-lg-auto my-2 justify-content-center my-md-0 text-small">
                        <li>
                            <a href="index.html" class="nav-link text-white">
                                <svg class="bi d-block mx-auto mb-1" width="24" height="24"><use xlink:href="#index"></use></svg>
                                Главная страница
                            </a>
                        </li>
                        <li>
                            <a href="cinemas.html" class="nav-link text-white">
                                <svg class="bi d-block mx-auto mb-1" width="24" height="24"><use xlink:href="#cinemas"></use></svg>
                                Кинотеатры
                            </a>
                        </li>
                        <li>
                            <a href="sessions.html" class="nav-link text-white">
                                <svg class="bi d-block mx-auto mb-1" width="24" height="24"><use xlink:href="#sessions"></use></svg>
                                Сеансы
                            </a>
                        </li>
                        <li>
                            <a href="movies.html" class="nav-link text-white">
                                <svg class="bi d-block mx-auto mb-1" width="24" height="24"><use xlink:href="#movies"></use></svg>
                                Фильмы
                            </a>
                        </li>
                        <li>
                            <a href="admin.html" class="nav-link text-white">
                                <svg class="bi d-block mx-auto mb-1" width="24" height="24"><use xlink:href="#user"></use></svg>
                                Админка
                            </a>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
       
    `;
    document.body.appendChild(header);
}

// Вызываем функцию создания футера
createHeader();
