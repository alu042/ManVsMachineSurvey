<script lang="ts">
    import type { LayoutData } from './$types';
    import "../app.css";
    
    export let data: LayoutData;
</script>

<div class="bg-bg w-screen h-screen px-28 py-14">
    <slot />
</div>