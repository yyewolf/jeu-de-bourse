<script>
    import back from '@/plugins/back.js';

    let infos = {};

    back.get('/v1/profile/infos')
        .then((res) => {
            infos = res.data.infos;
        })
        .catch((err) => {
            console.log(err);
        });
</script>

<span>{JSON.stringify(infos)}</span>