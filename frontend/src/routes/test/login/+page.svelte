<script>
    import back from '@/plugins/back.js';

    let resp = "";

    back.get('/v1/auth/connected')
        .then((res) => {
            resp = JSON.stringify(res.data);
        })
        .catch((err) => {
            resp = JSON.stringify(err.response.data);
        });
</script>

<p>{resp}</p>