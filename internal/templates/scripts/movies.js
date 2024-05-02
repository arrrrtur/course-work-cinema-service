function loadCinemas() {
    console.log('Loading cinemas...');

    fetch('http://127.0.0.2:30001/api/cinemas')
        .then(response => response.json())
        .then(cinemas => {
            const cinemaDropdown = document.getElementById('cinemaDropdown');
            cinemaDropdown.innerHTML = '';
            console.log(cinemas)

            cinemas.forEach(cinema => {
                const listItem = document.createElement('li');
                const link = document.createElement('a');
                link.classList.add('dropdown-item');
                link.href = '#';
                link.textContent = cinema.name;
                link.addEventListener('click', () => loadMoviesInCinema(cinema));
                listItem.appendChild(link);
                cinemaDropdown.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error('Ошибка при загрузке кинотеатров:', error);
        });
}

function loadMoviesInCinema(cinema) {
    console.log("Get movies in " + cinema.title);

    console.log("Get cinema halls");
    fetch('http://127.0.0.2:30001/api/cinema-halls/cinema/' + cinema.id)
        .then(response => response.json())
        .then(cinemaHalls => {
            console.log(cinemaHalls);
            const cinemaHallsList = document.getElementById('cinemaHallsList');
            cinemaHallsList.innerHTML = '';

            cinemaHalls.forEach(cinemaHall => {
                const listItem = document.createElement('li');
                const link = document.createElement('p');
                link.classList.add('list-group-item');
                link.textContent = cinemaHall.name + " " + cinemaHall.capacity + " " + cinemaHall.class;
                listItem.appendChild(link);
                cinemaHallsList.appendChild(listItem);

                loadSessionsInCinemaHall(cinemaHall)
                    .then(sessions => {
                        const sessionsList = document.createElement('ul');
                        sessionsList.classList.add('list-group', 'list-group-flush');

                        sessions.forEach(session => {
                            const sessionItem = document.createElement('li');
                            sessionItem.classList.add('list-group-item', 'mx-4');
                            sessionItem.textContent = "Дата - " + session.date +
                                " Осталось билетов - " + session.ticket_left;
                            loadMovieBySession(session)
                                .then(movie => {
                                    console.log(movie)
                                    let aboutMovie = document.createElement('p')
                                    aboutMovie.classList.add('h6', 'px-4')
                                    if (typeof movie.rating === 'object' && movie.rating !== null) {
                                        let ratingString = '';

                                        // Проход по элементам map
                                        for (const [key, value] of Object.entries(movie.rating)) {
                                            ratingString += `${key}: ${value}, `;
                                        }

                                        // Удаление лишней запятой в конце строки
                                        ratingString = ratingString.replace(/, $/, '');

                                        aboutMovie.textContent = movie.title + " " + "(" + movie.release_year + ") " +
                                            movie.duration + " минут " +
                                            "Оценки - " + ratingString;
                                    } else {
                                        aboutMovie.textContent = movie.title +
                                            "(" + movie.release_year + ")" +
                                            "\n Длительность - " + movie.duration +
                                            "\n Оценки - Недоступно";
                                    }

                                    const buyButton = document.createElement('button');
                                    buyButton.classList.add('btn', 'btn-primary', 'mt-2');
                                    buyButton.textContent = 'Купить';

                                    buyButton.setAttribute('data-bs-toggle', 'modal');
                                    buyButton.setAttribute('data-bs-target', '#buyModal');

                                    buyButton.addEventListener('click', () => showBuyForm(session, buyButton));
                                    sessionItem.appendChild(aboutMovie)
                                    sessionItem.appendChild(buyButton);

                                }).catch(error => {
                                    console.error('Failed to load movie: ', error)
                            })

                            sessionsList.appendChild(sessionItem);
                        });

                        listItem.appendChild(sessionsList);
                    })
                    .catch(error => {
                        console.error('Failed to load sessions:', error);
                    });
            });
        })
        .catch(error => {
            console.error('Failed to load cinema halls:', error);
        });
}

function loadSessionsInCinemaHall(cinemaHall) {
    console.log("Load sessions in " + cinemaHall.name);

    return fetch('http://127.0.0.2:30001/api/cinema-hall/sessions/' + cinemaHall.id)
        .then(response => response.json());
}

function loadMovieBySession(session) {
    console.log("load movie with id = " + session.movie_id)

    return fetch('http://127.0.0.2:30001/api/movies/' + session.movie_id)
        .then(response => response.json())
}

function showBuyForm(session, button) {
    // Создаем элементы модального окна
    const modalDiv = document.createElement('div');
    modalDiv.classList.add('modal', 'fade');
    modalDiv.setAttribute('data-bs-keyboard', 'false');
    modalDiv.setAttribute('id', 'buyModal');
    modalDiv.setAttribute('tabindex', '-1');
    modalDiv.setAttribute('aria-labelledby', 'buyModalLabel');
    modalDiv.setAttribute('aria-hidden', 'true');

    const modalDialogDiv = document.createElement('div');
    modalDialogDiv.classList.add('modal-dialog');

    const modalContentDiv = document.createElement('div');
    modalContentDiv.classList.add('modal-content');

    // Создаем заголовок модального окна
    const modalHeaderDiv = document.createElement('div');
    modalHeaderDiv.classList.add('modal-header');

    const modalTitleH1 = document.createElement('h1');
    modalTitleH1.classList.add('modal-title', 'fs-5');
    modalTitleH1.setAttribute('id', 'buyModalLabel');
    modalTitleH1.textContent = 'Оформление покупки';

    const closeButton = document.createElement('button');
    closeButton.setAttribute('type', 'button');
    closeButton.classList.add('btn-close');
    closeButton.setAttribute('data-bs-dismiss', 'modal');
    closeButton.setAttribute('aria-label', 'Close');

    modalHeaderDiv.appendChild(modalTitleH1);
    modalHeaderDiv.appendChild(closeButton);

    // Создаем тело модального окна
    const modalBodyDiv = document.createElement('div');
    modalBodyDiv.classList.add('modal-body');

    // Создаем форму в теле модального окна
    const form = document.createElement('form');
    form.addEventListener('submit', function (event) {
        event.preventDefault();
        // Обработка отправки формы и выполнение операции обновления в базе данных
        // Добавьте здесь соответствующий код
    });

    // Добавляем поля в форму
    const nameLabel = document.createElement('label');
    nameLabel.setAttribute('for', 'nameInput');
    nameLabel.textContent = 'Имя:';

    const nameInput = document.createElement('input');
    nameInput.setAttribute('type', 'text');
    nameInput.setAttribute('id', 'nameInput');
    nameInput.setAttribute('name', 'name');
    nameInput.setAttribute('required', 'true');

    const emailLabel = document.createElement('label');
    emailLabel.setAttribute('for', 'emailInput');
    emailLabel.textContent = 'Email:';

    const emailInput = document.createElement('input');
    emailInput.setAttribute('type', 'email');
    emailInput.setAttribute('id', 'emailInput');
    emailInput.setAttribute('name', 'email');
    emailInput.setAttribute('required', 'true');

    const submitButton = document.createElement('button');
    submitButton.setAttribute('type', 'submit');
    submitButton.classList.add('btn', 'btn-primary');
    submitButton.textContent = 'Оформить';

    // Добавляем поля в форму
    form.appendChild(nameLabel);
    form.appendChild(nameInput);
    form.appendChild(emailLabel);
    form.appendChild(emailInput);
    form.appendChild(submitButton);

    modalBodyDiv.appendChild(form);

    // Создаем подвал модального окна
    const modalFooterDiv = document.createElement('div');
    modalFooterDiv.classList.add('modal-footer');

    const closeButtonSecondary = document.createElement('button');
    closeButtonSecondary.setAttribute('type', 'button');
    closeButtonSecondary.classList.add('btn', 'btn-secondary');
    closeButtonSecondary.setAttribute('data-bs-dismiss', 'modal');
    closeButtonSecondary.textContent = 'Закрыть';

    modalFooterDiv.appendChild(closeButtonSecondary);

    // Собираем все элементы вместе
    modalContentDiv.appendChild(modalHeaderDiv);
    modalContentDiv.appendChild(modalBodyDiv);
    modalContentDiv.appendChild(modalFooterDiv);

    modalDialogDiv.appendChild(modalContentDiv);

    modalDiv.appendChild(modalDialogDiv);

    // Добавляем модальное окно в DOM
    document.body.appendChild(modalDiv);

    // Вызываем событие click для кнопки "Оформить"
    submitButton.dispatchEvent(new Event('click'));
}


function handleBuyFormSubmit(session, name, email, modal) {
    // ... (ваш существующий код)

    fetch('http://127.0.0.2:30001/api/buy-ticket', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            sessionId: session.id,
            name: name,
            email: email,
        }),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Покупка успешно оформлена:', data);
            modal.style.display = 'none'; // Закрываем модальное окно после успешной покупки
            // Дополнительная логика по обновлению интерфейса
        })
        .catch(error => {
            console.error('Ошибка при оформлении покупки:', error);
        });
}


