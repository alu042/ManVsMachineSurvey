<script lang="ts">
    export let width: string;
    export let height: string;
</script>

<div class="group inline-block cursor-pointer">
    <div class="flex items-center justify-center p-1 rounded-full border-solid border-2 border-red-500 transition-colors duration-300 hover:bg-red-500">
        <svg xmlns="http://www.w3.org/2000/svg" height={height} width={width} viewBox="0 0 64 512" class="fill-current text-red-500 group-hover:text-white transition-colors duration-300">
            <path d="M64 64c0-17.7-14.3-32-32-32S0 46.3 0 64V320c0 17.7 14.3 32 32 32s32-14.3 32-32V64zM32 480a40 40 0 1 0 0-80 40 40 0 1 0 0 80z"/>
        </svg>
    </div>
</div>
