<script lang="ts">
    import type { LayoutData } from './$types';
    import "../app.css";
    
    export let data: LayoutData;
</script>

<div class="bg-bg w-screen h-screen px-32 py-10">
    <slot />
</div>