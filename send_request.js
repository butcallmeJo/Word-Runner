$(function() {
    $('#submit').click(function() {
		$('#progress').show()
        $.ajax({
            url: 'http://cp-0627-2.hbtn.io:8091/wiki/',
            data: $('form').serialize(),
            type: 'POST',
            success: function(response) {
                // $('#progress').css("display","block");
                // In case of success, we should display the result
                // For now we can debug with console response <-- return value of go function
                console.log(response);
				$('#progress').hide()
                $("#result").html(response)

                // var str1 = "<p>"
                // var str2 = str1.concat(response)
                // var printable_result = str2.concat(str1)
                //
                // $("result").html(printable_result)

            },
            error: function(error) {
                console.log(error);
            }
        });
    });
});
