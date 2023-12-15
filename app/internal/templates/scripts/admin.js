
document.addEventListener('DOMContentLoaded', function () {
    const addCinemaForm = document.getElementById('addCinemaForm');
    const submitBtn = document.getElementById('submitBtn');

    submitBtn.addEventListener('click', function () {
        const name = document.getElementById('name').value;
        const address = document.getElementById('address').value;

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

                const modal = document.getElementById('exampleModal');
                addCinemaForm.reset()
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});
