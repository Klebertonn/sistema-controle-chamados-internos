async function salvarChamado() {

    const chamado = {

        titulo:
            document.getElementById("titulo").value,

        descricao:
            document.getElementById("descricao").value,

        status: "Aberto",

        prioridade:
            document.getElementById("prioridade").value,

        solicitante:
            document.getElementById("solicitante").value,

        responsavelId: 0
    };

    try {

        await criarChamado(chamado);

        alert("Chamado criado com sucesso!");

        limparFormulario();

        carregarChamados();

    } catch(error) {

        console.error(error);

        alert("Erro ao criar chamado");
    }
}

function limparFormulario() {

    document.getElementById("titulo").value = "";

    document.getElementById("descricao").value = "";

    document.getElementById("prioridade").value = "Baixa";

    document.getElementById("solicitante").value = "";
}