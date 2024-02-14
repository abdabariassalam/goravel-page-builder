$(document).ready(function(){
    // Handle click event on delete buttons
    $('#submit').on('click', function() {
        const url = $(this).data('url');
        $.ajax({
            url: url,
            type: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify({
                'title': $('#title').val(),
                'short_description': $('#short_description').val(),
                'template': $('input[type=radio][name=template]:checked').val(),
            }),
            success: function(response) {
                alert(response);
                location.replace("/");
            },
            error: function(response) {
                alert(response.responseJSON);
            }
        });
    });
    var currentValue = $('input[type=radio][name=template]:checked')
    $('input[type=radio][name=template]').on('click', function() {
        if ($(this).val() != currentValue.val()) {
            if (confirm('Mengganti template akan mernghapus content yang tersimpan. apakah anda yakin?')) {
                return true
            } else {
                return false;
            }
        }
    });
});
