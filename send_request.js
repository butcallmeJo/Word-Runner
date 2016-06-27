$(function() {
    $('#submit').click(function() {

        $.ajax({
            url: '/*INSERT PHP URI HERE*/',
            data: $('form').serialize(),
            type: 'POST',
            success: function(response) {
                // In case of success, we should display the result
                // For now we can debug with console response <-- return value of php function?
                console.log(response);
                
                /*
                $("textarea").html(// code to display result here)
                */
            },
            error: function(error) {
                console.log(error);
            }
        });
    });
});
