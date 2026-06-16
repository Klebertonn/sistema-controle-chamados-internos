async function carregarChamados() {

    const response = await fetch("http://localhost:8080/chamados");

    const chamados = await response.json();

    console.log("Resposta API:", chamados);

    const tabela = document.getElementById("tabelaChamados");

    tabela.innerHTML = "";

    chamados.forEach(c => {

        const linha = `
            <tr>
                <td>${c.ordemId}</td>
                <td>${c.titulo}</td>
                <td>${c.descricao}</td>
                <td>${c.status}</td>
                <td>${c.prioridade}</td>
                <td>${c.solicitante}</td>
            </tr>
        `;

        tabela.innerHTML += linha;
    });
}

window.onload = carregarChamados;