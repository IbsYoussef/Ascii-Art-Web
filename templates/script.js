document.addEventListener('DOMContentLoaded', function() {
    const textarea = document.getElementById('textvalue');
    const form = document.getElementById('ascii-art-form');
    const refreshButton = document.getElementById('refreshButton');

    textarea.addEventListener('keydown', function(event) {
        if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault();
            form.submit();
        }
    });

    refreshButton.addEventListener('click', function() {
        location.reload();
    });
});
