if ($(".izimodal").length) {
    $("#report-modal").iziModal();


    // add class on file upload
    $('input[type="file"]').on("change", function (e) {
        var filename = $(this).val();
        var aux = filename.split('.');
        var extension = aux[aux.length -1].toUpperCase();
        console.log(extension)
        if (extension.toLowerCase() == "txt")  {
            $(".upload-btn-wrapper .btn").addClass("btn--fill")
            $(".upload-btn-wrapper .btn").text(filename.split('\\').pop())
            return true
        }
        alert("Chats must be a .txt file!")
        return false

    })


    $("#report-modal .btn--submit").on("click", function(e){
        var valid = true;
        e.preventDefault()

        $(':input[required]:visible').each(function() {
            if ($(this).val() == '') {
                alert("Please ensure you've selected a chat and inputted your email.")
                valid = false;
                return false

            }
        });
        if (valid) {
            var form = new FormData($("#report-modal form")[0]);
            $("form").hide()
            $(".btn--submit").hide()
            $(".loading").addClass("show")
            $.ajax({
                type: 'post',
                url: '/report',
                processData: false,
                contentType: false,
                // data: $('#report-modal form').serialize(),
                data: form,
                success: function (data, status) {
                    console.log(data)
                    $(".circle-loader").addClass("load-complete")
                    $(".checkmark").show()
                },
                error: function (error, status) {
                    console.log(error, status)
                }
            });
        }
    })

}