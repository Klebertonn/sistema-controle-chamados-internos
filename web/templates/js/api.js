const API_URL = "http://localhost:8080";

async function listarChamados() {
    const response = await fetch(`${API_URL}/chamados`);

    if (!response.ok) {
        throw new Error("Erro ao listar chamados");
    }

    return await response.json();
}

async function criarChamado(chamado) {

    const response = await fetch(`${API_URL}/chamados`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(chamado)
    });

    if (!response.ok) {
        throw new Error("Erro ao criar chamado");
    }

    return response;
}

async function buscarChamado(id) {

    const response =
        await fetch(`${API_URL}/chamados/${id}`);

    if (!response.ok) {
        throw new Error("Chamado não encontrado");
    }

    return await response.json();
}

async function atualizarChamado(id, chamado) {

    const response = await fetch(
        `${API_URL}/chamados/${id}`,
        {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(chamado)
        }
    );

    if (!response.ok) {
        throw new Error("Erro ao atualizar");
    }

    return response;
}