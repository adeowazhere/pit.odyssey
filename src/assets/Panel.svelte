<script>
    import { onMount } from "svelte";

    export let TEXT
  
    const SPEED = 40;
    let dialogue = ""
    let isDialogueFinished = false
  
    function typeWriter() {
      let i = 0
  
      const interval = setInterval(() => {
        if (i < TEXT.length) {
          dialogue += TEXT.charAt(i)
          i++
        } else {
          clearInterval(interval)
          isDialogueFinished = true
        }
      }, SPEED);
    }
  
  onMount(typeWriter)
  </script>
  
  <p class="text-typing">{@html dialogue}</p>
  
  {#if isDialogueFinished}
    <section id="input-section"><span>>></span><input autofocus/></section>
  {/if}

  <style>
  </style>