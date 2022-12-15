<script>
    import back from '@/plugins/back.js';
    import Nav from '@/routes/nav.svelte';
    import { Tooltip } from 'flowbite-svelte';
    import "@/app.css";

    let stocks = [];
    let page = 1;
    let pages = 1;
    
    function getStocks() {
        let body = {
            page: page
        }
        back.post('/v1/stocks/', body)
            .then((res) => {
                stocks = res.data.stocks;
                pages = res.data.pages;
            })
            .catch((err) => {
                console.log(err);
            });
    }

    getStocks();
</script>

<Nav />

<!-- Center container that will list all stocks & page navigation -->
<div class="flex items-center justify-center bg-slate-600 h-full">
    <!-- Use a card as a container -->
    <div class="w-full justify-center max-w-3xl p-4 bg-white border rounded-lg shadow-md sm:p-8 dark:bg-slate-800 dark:border-slate-700 mt-5 mb-5">
        <div class="flex items-center justify-between mb-7">
            <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white">Liste des actions</h5>
        </div>
        <ul class="divide-y divide-gray-200 dark:divide-gray-700">
            <!--  for each stock -->
            {#each stocks as stock}
                <li class="py-3 sm:py-4">
                    <div class="flex items-center space-x-4">
                        <div class="flex-1 min-w-0">
                            <div class="text-sm text-gray-900 truncate dark:text-white flex justify-start">
                                <p class="font-extrabold">{stock.name}</p>
                                <p class="font-small  ml-3">{stock.symbol}</p>
                                <!-- Tooltip if SRD / Loan -->
                                {#if stock.srd}
                                    <button class="ml-2">
                                        <span class="sr-only">SRD</span>
                                        <svg class="h-5 w-5 text-gray-400" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
                                            <path fill-rule="evenodd" d="M10 2a8 8 0 100 16 8 8 0 000-16zm0 14a6 6 0 110-12 6 6 0 010 12z" clip-rule="evenodd" />
                                            <path fill-rule="evenodd" d="M10 6a1 1 0 00-1 1v4a1 1 0 102 0V7a1 1 0 00-1-1z" clip-rule="evenodd" />
                                            <path fill-rule="evenodd" d="M10 12a1 1 0 100 2 1 1 0 000-2z" clip-rule="evenodd" />
                                        </svg>
                                    </button>
                                    <Tooltip class="bg-slate-900">SRD disponible</Tooltip>
                                {/if}
                                {#if stock.loan}
                                    <button class="ml-2">
                                        <svg class="h-5 w-5 text-gray-400" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
                                            <path fill-rule="evenodd" d="M10 2a8 8 0 100 16 8 8 0 000-16zm0 14a6 6 0 110-12 6 6 0 010 12z" clip-rule="evenodd" />
                                            <path fill-rule="evenodd" d="M10 6a1 1 0 00-1 1v4a1 1 0 102 0V7a1 1 0 00-1-1z" clip-rule="evenodd" />
                                            <path fill-rule="evenodd" d="M10 12a1 1 0 100 2 1 1 0 000-2z" clip-rule="evenodd" />
                                        </svg>
                                    </button>
                                    <Tooltip class="bg-slate-900">Vente à perte disponible</Tooltip>
                                {/if}
                            </div>
                            <div class="text-sm text-gray-500 truncate dark:text-gray-400">
                                {stock.id}
                            </div>
                        </div>
                        <div class="inline-flex items-center text-base font-normal text-gray-900 dark:text-white">
                            <p class="mr-5">{stock.last_price/1000}€</p>

                            {#if stock.last_price > stock.opening_price}
                                <svg class="h-5 w-5 text-green-400" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
                                    <path fill-rule="evenodd" d="M10 2a8 8 0 100 16 8 8 0 000-16zm0 14a6 6 0 110-12 6 6 0 010 12z" clip-rule="evenodd" />
                                    <path fill-rule="evenodd" d="M10 6a1 1 0 00-1 1v4a1 1 0 102 0V7a1 1 0 00-1-1z" clip-rule="evenodd" />
                                    <path fill-rule="evenodd" d="M10 12a1 1 0 100 2 1 1 0 000-2z" clip-rule="evenodd" />
                                </svg>
                            {:else if stock.last_price < stock.opening_price}
                                <svg class="h-5 w-5 text-red-400" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
                                    <path fill-rule="evenodd" d="M10 2a8 8 0 100 16 8 8 0 000-16zm0 14a6 6 0 110-12 6 6 0 010 12z" clip-rule="evenodd" />
                                    <path fill-rule="evenodd" d="M10 6a1 1 0 00-1 1v4a1 1 0 102 0V7a1 1 0 00-1-1z" clip-rule="evenodd" />
                                    <path fill-rule="evenodd" d="M10 12a1 1 0 100 2 1 1 0 000-2z" clip-rule="evenodd" />
                                </svg>
                            {:else}
                                <svg class="h-5 w-5 text-gray-400" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
                                    <path fill-rule="evenodd" d="M10 2a8 8 0 100 16 8 8 0 000-16zm0 14a6 6 0 110-12 6 6 0 010 12z" clip-rule="evenodd" />
                                    <path fill-rule="evenodd" d="M10 6a1 1 0 00-1 1v4a1 1 0 102 0V7a1 1 0 00-1-1z" clip-rule="evenodd" />
                                    <path fill-rule="evenodd" d="M10 12a1 1 0 100 2 1 1 0 000-2z" clip-rule="evenodd" />
                                </svg>
                            {/if}
                            {#if stock.last_price > stock.opening_price}
                                <p class="ml-2 text-green-400">+{(100*stock.last_price/stock.opening_price-100).toFixed(2)}%</p>
                            {:else if stock.last_price < stock.opening_price}
                                <p class="ml-2 text-red-400">-{(100-100*stock.last_price/stock.opening_price).toFixed(2)}%</p>
                            {:else}
                                <p class="ml-2">{(100-100*stock.last_price/stock.opening_price).toFixed(2)}%</p>
                            {/if}
                        </div>
                    </div>
                </li>
            {/each}
        </ul>

        <!-- Pages selector 5 pages in advance -->
        <div class="flex justify-center mt-4">
            <div class="flex space-x-2">
                {#if page > 1}
                <button class="px-4 py-2 rounded-md text-sm font-bold text-gray-700 bg-white border border-gray-300 hover:bg-gray-50 dark:bg-gray-800 dark:text-gray-200 dark:hover:bg-gray-700 dark:border-gray-700 dark:hover:text-white" on:click={() => { page = page-1; getStocks() }}>
                    &lt;
                </button>
                {/if}
                <!-- Always 5 indexes, if negative, offsets the first one -->
                {#each [page - 2, page - 1, page, page + 1, page + 2] as page_number}
                    {#if page_number <= pages && page_number > 0}
                        {#if page_number == page}
                            <button class="px-4 py-2 rounded-md text-sm font-bold text-white bg-gray-900 border border-gray-300 hover:bg-gray-700 dark:bg-gray-700 dark:hover:bg-gray-600 dark:border-gray-700 dark:hover:text-white" on:click={() => { page = page_number; getStocks() }}>
                                {page_number}
                            </button>
                        {:else}
                            <button class="px-4 py-2 rounded-md text-sm font-medium text-gray-700 bg-white border border-gray-300 hover:bg-gray-50 dark:bg-gray-800 dark:text-gray-200 dark:hover:bg-gray-700 dark:border-gray-700 dark:hover:text-white" on:click={() =>{ page = page_number; getStocks() }}>
                                {page_number}
                            </button>
                        {/if}
                    {/if}
                {/each}
                {#if page < pages}
                    <button class="px-4 py-2 rounded-md text-sm font-bold text-gray-700 bg-white border border-gray-300 hover:bg-gray-50 dark:bg-gray-800 dark:text-gray-200 dark:hover:bg-gray-700 dark:border-gray-700 dark:hover:text-white" on:click={() =>{ page = page + 1; getStocks() }}>
                        >
                    </button>
                {/if}
            </div> 
        </div>  
    </div>
</div>
