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

<div class="flex flex-col max-w-full text-white mt-5">
    <div class="flex flex-col items-center">
        <div class="items-center">
            <p class="text-2xl">Profil de <span class="font-extrabold">{infos.username}</span></p>
            <div class="divider"></div> 
        </div>
        <div class="items-center">
            <p>Valorisation : {(infos.valorisation/10000).toFixed(2)}€ 
                {#if infos.var_valorisation > 0}
                    <span class="text-green-500">(+{(infos.var_valorisation/10000).toFixed(2)}%=</span>
                {:else if infos.var_valorisation < 0}
                    <span class="text-red-500">(-{(infos.var_valorisation/10000).toFixed(2)}%)</span>
                {:else}
                    <span class="text-white">(+0.00%)</span>
                {/if}
            </p> 
            <p>Disponible au comptant : {(infos.comptant/10000).toFixed(2)}€</p>  
            <p>Disponible au SRD : {(infos.srd/10000).toFixed(2)}€</p>     
        </div>
    </div>
</div>