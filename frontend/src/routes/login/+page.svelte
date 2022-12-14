<script>
    import back from '@/plugins/back.js';
    import { Recaptcha, recaptcha, observer } from "svelte-recaptcha-v2";
    import Nav from '@/routes/nav.svelte';
    import "@/app.css";

    const googleRecaptchaSiteKey=import.meta.env.VITE_JDB_RECAPTCHA_PKEY;

    let data = {
        token: "",
        login: "",
        password: "",
    };  

    let alert = {
        show: false,
        message: "",
    };

    let enableForm = false;

    async function onSubmit() {
        recaptcha.execute();

        // wait for the data.token to be set
        while (data.token === "") {
            await new Promise(r => setTimeout(r, 100));
        }

        let body = {
            captchaToken: data.token,
            username: data.login,
            password: data.password,
        };
        back.post('/v1/auth/login', body)
            .then(() => {
                window.location.href = '/';
            })
            .catch((err) => {
                alert.show = true;
                alert.message = err.response.data.status;
            });
    }

    const onCaptchaReady = (event) => {
        enableForm = true;
    };

    const onCaptchaSuccess = (event) => {
        data.token = event.detail.token;
    };
</script>

<div class="bg-gray-100 min-h-screen">

    <Nav />

    <div class="flex items-center justify-center bg-gray-100">
        <div class="px-8 py-6 mt-4 text-left bg-white shadow-lg columns-1">
            <h3 class="text-2xl font-bold text-center">Connexion à Stock Race</h3>
            {#if alert.show}
                <div role="alert" class="mt-5">
                    <div class="bg-red-500 text-white font-bold rounded-t px-4 py-2">
                        Erreur
                    </div>
                    <div class="border border-t-0 border-red-400 rounded-b bg-red-100 px-3 py-3 text-red-700 break-normal">
                        <p>{alert.message}</p>
                    </div>
                </div>
            {/if}
            <form>
                <div class="mt-4">
                    <div>
                        <label class="block" for="username">Nom d'utilisateur / Email<label>
                            <input bind:value={data.login} type="text" placeholder="Nom d'utilisateur / Email"
                                class="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600">
                    </div>
                    <div class="mt-4">
                        <label class="block">Mot de passe<label>
                            <input bind:value={data.password} type="password" placeholder="Mot de passe"
                                class="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600">
                    </div>
                    <div class="flex items-baseline justify-center">
                        <button disabled={!enableForm} type="submit" on:click|preventDefault={onSubmit} class="px-6 py-2 mt-4 text-white bg-blue-600 rounded-lg hover:bg-blue-900">Connexion</button>
                    </div>
                    <div class="flex items-baseline justify-center mt-4">
                        <a href="/register" class="text-blue-600 hover:text-blue-900">Pas encore de compte ?</a>
                    </div>

                    <Recaptcha
                        sitekey={googleRecaptchaSiteKey}
                        badge={"top"}
                        size={"invisible"}
                        on:success={onCaptchaSuccess}
                        on:ready={onCaptchaReady}
                    />
                </div>
            </form>
        </div>
    </div>

</div>