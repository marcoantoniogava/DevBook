$("#formulario-cadastro").on("submit", criarUsuario); //quando rolar um submit nesse formulario, a funcao criarUsuario é chamada

function criarUsuario(evento) {
    evento.preventDefault(); //previne o comportamento padrao do formulario ao ser enviado
    console.log("Dentro da função usuário!");

    if ($("#senha").val() != $("#confirmar-senha").val()) {
        console.log("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("nome").val(),
            email: $("email").val(),
            nick: $("nick").val(),
            senha: $("senha").val(),
        }
    })
}

