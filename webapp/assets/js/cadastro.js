$("#formulario-cadastro").on("submit", criarUsuario); //quando rolar um submit nesse formulario, a funcao criarUsuario é chamada

function criarUsuario(evento) {
    evento.preventDefault(); //previne o comportamento padrao do formulario ao ser enviado

    if ($("#senha").val() != $("#confirmar-senha").val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
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
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $("#email").val(),
                        senha: $("#senha").val(),
                    }
                }).done(function() {
                    window.location.href = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Erro autenticar o usuário!", "error");
                })
            })
    }).fail(function() { //400 404 401 403 500
        Swal.fire("Ops...", "Erro ao cadastrar usuário!", "error");
    });
}
