$(function() {
    $('#submit').click(function() {
        console.log("This is being called")
        console.log($('form').serialize())
        $.ajax({
            url: 'http://cp-0627-2.hbtn.io:8091/wiki/',
            data: $('form').serialize(),
            type: 'POST',
            success: function(response) {
                // $('#progress').css("display","block");
				console.log("hi")
                // In case of success, we should display the result
                // For now we can debug with console response <-- return value of go function
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
