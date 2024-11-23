const apiUrl = 'http://localhost:8080/v1/beers';

// Função para carregar todas as cervejas
async function loadBeers() {
    const response = await fetch(apiUrl);
    const beers = await response.json();
    const beerList = document.getElementById('beerList');
    
    // Limpar lista existente
    beerList.innerHTML = '';

    if (beers.length === 0) {
        beerList.innerHTML = '<p>No beers available.</p>';
        return;
    }

    beers.forEach(beer => {
        const beerDiv = document.createElement('div');
        beerDiv.className = 'beer-item';
        beerDiv.innerHTML = `
            <h3>${beer.name}</h3>
            <p>Type: ${beer.type.name}</p>
            <p>Style: ${beer.style.name}</p>
        `;
        beerList.appendChild(beerDiv);
    });
}

// Função para adicionar uma nova cerveja
async function addBeer(event) {
    event.preventDefault();

    const name = document.getElementById('name').value;
    const typeName = document.getElementById('type').value;
    const styleName = document.getElementById('style').value;

    const newBeer = {
        name: name,
        type: { name: typeName },
        style: { name: styleName }
    };

    const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newBeer)
    });

    if (response.ok) {
        alert('Beer added successfully!');
        document.getElementById('beerForm').reset();
    } else {
        alert('Failed to add beer');
    }
}

// Evento para adicionar cerveja
document.getElementById('beerForm').addEventListener('submit', addBeer);

// Evento para carregar cervejas ao clicar no botão
document.getElementById('loadBeersButton').addEventListener('click', loadBeers);
