<script>
	import "../app.css";
  import back from '@/plugins/back.js';
  
  let connected = false;

  back.get('/v1/auth/connected')
    .then((res) => {
      connected = true;
    })
    .catch((err) => {
      connected = false;
    });
</script>

<nav class="flex items-center justify-between flex-wrap p-6  bg-blue-400">
  <div class="flex items-center flex-shrink-0 text-white mr-6">
    <a href="/" class="font-semibold text-xl tracking-tight">Stock Race</a>
  </div>
  <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
    <div class="text-sm lg:flex-grow">
      <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4">
        Docs
      </a>
    </div>
    <div>
      {#if connected}
        <a href="/logout" class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Logout</a>
      {:else}
        <a href="/login" class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Login</a>
      {/if}
    </div>
  </div>
</nav>