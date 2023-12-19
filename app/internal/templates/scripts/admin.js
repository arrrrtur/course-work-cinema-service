// add cinema
document.addEventListener('DOMContentLoaded', function () {
    const addCinemaForm = document.getElementById('AddCinemaForm');
    const submitBtn = document.getElementById('AddCinemaBtn');

    submitBtn.addEventListener('click', function () {
        const name = document.getElementById('AddCinemaName').value;
        const address = document.getElementById('AddCinemaAddress').value;

        const formData = {
            name: name,
            address: address
        };

        fetch('http://127.0.0.2:30001/api/cinemas', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => {})
            .then(data => {
                console.log(data);

                const modal = document.getElementById('AddCinemaModal');
                addCinemaForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});

// edit cinema
document.getElementById('CheckCinemaBtn').addEventListener('click', function() {
    // Get data from the form
    let id = document.getElementById('EditCinemaId').value
    let name = document.getElementById('EditCinemaName')
    let address = document.getElementById('EditCinemaAddress');



    fetch('http://127.0.0.2:30001/api/cinemas/' + id)
        .then(response => response.json())
        .then(cinema => {
            console.log(cinema);
            name.value = cinema.name
            address.value = cinema.address

        })
        .catch(error => {
            console.error('Error:', error);
        });


});
document.addEventListener('DOMContentLoaded', function () {
    const editCinemaForm = document.getElementById('EditCinemaForm');
    const editBtn = document.getElementById('EditCinemaBtn');

    editBtn.addEventListener('click', function () {
        const id = document.getElementById('EditCinemaId').value
        const name = document.getElementById('EditCinemaName').value
        const address = document.getElementById('EditCinemaAddress').value

        const formData = {
            id: parseInt(id),
            name: name,
            address: address,

        };

        fetch('http://127.0.0.2:30001/api/cinemas', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => {console.log(response.json())})
            .then(data => {
                console.log(data);

                const modal = document.getElementById('EditMovieModal');
                editCinemaForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});

// delete cinema
document.addEventListener('DOMContentLoaded', function () {
    const deleteCinemaForm = document.getElementById('DeleteCinemaForm');
    const deleteBtn = document.getElementById('DeleteCinemaForm');

    deleteBtn.addEventListener('click', function () {
        const id = document.getElementById('DeleteCinemaId').value

        fetch('http://127.0.0.2:30001/api/cinemas/:' + id, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
        })
            .then(response => {console.log(response.json())})
            .then(data => {
                console.log(data);

                const modal = document.getElementById('DeleteCinemaModal');
                deleteCinemaForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});


function marshallRating(rating) {
    const keyValuePairs = rating.split(',').map(pair => pair.trim());
    const resultObject = {};
    keyValuePairs.forEach(pair => {
        const [key, value] = pair.split(' => ');
        resultObject[key] = value;
    });

    return resultObject
}

// add movie
document.addEventListener('DOMContentLoaded', function () {
    const addMovieForm = document.getElementById('AddMovieForm');
    const addBtn = document.getElementById('AddMovieBtn');

    addBtn.addEventListener('click', function () {
        const title = document.getElementById('AddMovieTitle').value
        const description = document.getElementById('AddMovieDescription').value
        const duration = document.getElementById('AddMovieDuration').value
        const release_year = document.getElementById('AddMovieReleaseYear').value
        const director = document.getElementById('AddMovieDirector').value
        const rating = marshallRating(document.getElementById('AddMovieRating').value)



        const formData = {
            ID: parseInt(id),
            title: title,
            description: description,
            duration: parseInt(duration),
            release_year: parseInt(release_year),
            director: parseInt(director),
            rating: rating,
        };

        console.log(JSON.stringify(formData))

        fetch('http://127.0.0.2:30001/api/movies', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => {console.log(response.json())})
            .then(data => {
                console.log(data);

                const modal = document.getElementById('EditMovieModal');
                addMovieForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});


// edit movie
document.getElementById('CheckMovieBtn').addEventListener('click', function() {
    // Get data from the form
    let id = document.getElementById('EditMovieId').value
    let title = document.getElementById('EditMovieTitle')
    let description = document.getElementById('EditMovieDescription');
    let duration = document.getElementById('EditMovieDuration');
    let release_year = document.getElementById('EditMovieReleaseYear');
    let director = document.getElementById('EditMovieDirector');
    let rating = document.getElementById('EditMovieRating');


    fetch('http://127.0.0.2:30001//api/movies/' + id)
        .then(response => response.json())
        .then(movie => {
            console.log(movie);

            title.value = movie.title
            description.value = movie.description
            duration.value = movie.duration
            release_year.value = movie.release_year
            director.value = movie.director
            let getRating = (movie) => {
                let resultString = '';

                for (let [key, value] of Object.entries(movie.rating)) {
                    resultString += `${key} => ${value}, `;
                }

                return resultString.trim();
            }
            rating.value = getRating(movie)
        })
        .catch(error => {
            console.error('Error:', error);
        });


});
document.addEventListener('DOMContentLoaded', function () {
    const editMovieForm = document.getElementById('EditMovieForm');
    const editBtn = document.getElementById('EditMovieBtn');

    editBtn.addEventListener('click', function () {
        const id = document.getElementById('EditMovieId').value
        const title = document.getElementById('EditMovieTitle').value
        const description = document.getElementById('EditMovieDescription').value
        const duration = document.getElementById('EditMovieDuration').value
        const release_year = document.getElementById('EditMovieReleaseYear').value
        const director = document.getElementById('EditMovieDirector').value
        const rating = marshallRating(document.getElementById('EditMovieRating').value)

        const formData = {
            ID: parseInt(id),
            title: title,
            description: description,
            duration: parseInt(duration),
            release_year: parseInt(release_year),
            director: parseInt(director),
            rating: rating,
        };

        console.log(JSON.stringify(formData))

        fetch('http://127.0.0.2:30001/api/movies', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => {console.log(response.json())})
            .then(data => {
                console.log(data);

                const modal = document.getElementById('EditMovieModal');
                editMovieForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});

// delete movie
document.addEventListener('DOMContentLoaded', function () {
    const deleteMovieForm = document.getElementById('DeleteMovieForm');
    const deleteBtn = document.getElementById('DeleteMovieForm');

    deleteBtn.addEventListener('click', function () {
        const id = document.getElementById('DeleteMovieId').value

        fetch('http://127.0.0.2:30001/api/movies/:' + id, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
        })
            .then(response => {console.log(response.json())})
            .then(data => {
                console.log(data);

                const modal = document.getElementById('DeleteCinemaModal');
                deleteMovieForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});
