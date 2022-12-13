<script>
    import back from '@/plugins/back.js';

    let pkey = import.meta.env.VITE_JDB_RECAPTCHA_PKEY;

    function onSubmit() {
        doRecaptcha();
    }

    function doRecaptcha() {
        grecaptcha.ready(function() {
            grecaptcha.execute(pkey, { action: "submit" }).then(function(t) {
                // send to backend
                back.post('/v1/captcha', { captchaToken: t }).then((res) => {
                    console.log(res.data);
                });
            });
        });
    }
</script>

<svelte:head>
    <script src="https://www.google.com/recaptcha/api.js?render={pkey}"></script>
</svelte:head>

<h1>Test recaptcha</h1>

<style>
    main {
      font-family: sans-serif;
      text-align: center;
    }
</style>
  
<main>
    <form>
        <button type="submit" on:click|preventDefault={onSubmit}>submit</button>
    </form>
</main>