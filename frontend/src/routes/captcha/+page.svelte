<script>
    let pkey = import.meta.env.VITE_JDB_RECAPTCHA_PKEY;
    let token = '';

    function onSubmit() {
        doRecaptcha();
    }

    function doRecaptcha() {
        grecaptcha.ready(function() {
            grecaptcha.execute(pkey, { action: "submit" }).then(function(t) {
                token = t;
            });
        });
    }
</script>

<svelte:head>
    <script src="https://www.google.com/recaptcha/api.js?render={pkey}"></script>
</svelte:head>

<h1>Test recaptcha</h1>

<p>Token : {token}</p>

<style>
    main {
      font-family: sans-serif;
      text-align: center;
    }
</style>
  
<main>
    <form on:submit|preventDefault={onSubmit}>
      <button type="submit">submit</button>
    </form>
</main>