$("#formulario-cadastro").on("submit", criarUsuario); //quando rolar um submit nesse formulario, a funcao criarUsuario é chamada

function criarUsuario(evento) {
    evento.preventDefault(); //previne o comportamento padrao do formulario ao ser enviado

    if ($("#senha").val() != $("#confirmar-senha").val()) {
        console.log("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: $("#senha").val(),
        }
    }).done(function() { //201 200 204
        console.log("Usuário cadastrado com sucesso!");
    }).fail(function(erro) { //400 404 401 403 500
        console.log(erro)
        console.log("Erro ao cadastrar usuário!");
    });
}
