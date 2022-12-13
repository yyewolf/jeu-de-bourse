<script>
    import "@/app.css";
    let pkey = import.meta.env.VITE_JDB_RECAPTCHA_PKEY;
    
    let data = {
        username: "",
        password: "",
        passwordverify: "",
    };  
    
    function onSubmit() {
        doRecaptcha();
    }
    
    function doRecaptcha() {
        grecaptcha.ready(function() {
            grecaptcha.execute(pkey, { action: "submit" }).then(function(t) {
                // send to backend
                let body = {
                    captchaToken: t,
                    username: data.username,
                    password: data.password,
                    passwordverify: data.passwordverify,
                };
                back.post('/v1/captcha', body).then((res) => {
                    console.log(res.data);
                });
            });
        });
    }
    </script>

<svelte:head>
    <script src="https://www.google.com/recaptcha/api.js?render={pkey}"></script>
</svelte:head>

<div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="px-8 py-6 mt-4 text-left bg-white shadow-lg">
        <h3 class="text-2xl font-bold text-center">Inscription à Jeu de Bourse</h3>
        <form action="">
            <div class="mt-4">
                <div>
                    <label class="block" for="username">Nom d'utilisateur<label>
                        <input bind:value={data.username} type="text" placeholder="Nom d'utilisateur"
                            class="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600">
                </div>
                <div class="mt-4">
                    <label class="block">Mot de passe<label>
                        <input bind:value={data.password} type="password" placeholder="Mot de passe"
                            class="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600">
                </div>
                <div class="mt-4">
                    <label class="block">Confirmation du mot de passe<label>
                        <input bind:value={data.passwordverify} type="passwordverify" placeholder="Mot de passe"
                            class="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600">
                </div>
                <div class="flex items-baseline justify-between">
                    <button class="px-6 py-2 mt-4 text-white bg-blue-600 rounded-lg hover:bg-blue-900">Inscription</button>
                </div>
            </div>
        </form>
    </div>
</div>