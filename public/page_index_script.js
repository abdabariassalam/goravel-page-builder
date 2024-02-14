$(document).ready(function(){
    // Handle click event on delete buttons
    $('#data-page').on('click', '.deleteButton', function() {
        const url = $(this).data('url');
        if (confirm('Are you sure you want to delete this page?')) {
            $.ajax({
                url: url,
                type: 'DELETE',
                contentType: 'application/json',
                success: function(response) {
                    alert(response);
                    location.reload();
                },
                error: function(response) {
                    alert(response.responseJSON);
                }
            });
        }
    });
});