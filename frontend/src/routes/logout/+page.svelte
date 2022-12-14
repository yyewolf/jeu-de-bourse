<script>
    import back from '@/plugins/back.js';
    back.get('/v1/auth/logout')
        .then(() => {
            window.location.href = '/';
        })
        .catch(() => {
            window.location.href = '/';
        });
</script>