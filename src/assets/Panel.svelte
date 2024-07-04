<script>
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher()

    export let TEXT, imageLink, imageClass, imageAlt, optionA, optionB, optionC, optionD
  
    const SPEED = 35;
    let dialogue = ""
    let isDialogueFinished = false
    let i = 0
  
    function typeWriter() {  
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

    function skipDialogue() {
      i = TEXT.length
      dialogue = TEXT
      isDialogueFinished = true
    }

    function skipDialogueIfEnterKeyPressed(e) {
      if (e.keyCode !== 13) { return }

      skipDialogue()
    }
    
  onMount(typeWriter)
  </script>
  
  <p class="text-typing">{@html dialogue}</p>
  {#if !isDialogueFinished}
    <button id="skip-button" on:click={skipDialogue}>Skip</button>
    <p id="skip-hint">Press Enter to skip.</p>
  {/if}
  
  {#if isDialogueFinished}
    {#if imageLink}
      <div class="image-container">
        <img class={imageClass} src={imageLink} alt={imageAlt}/>
      </div>
    {/if}

    <section id="story-options">
    {#if optionA}
      <button on:click={() => {dispatch("message", "a")}}>{ optionA }</button>
    {/if}
    {#if optionB}
      <button on:click={() => {dispatch("message", "b")}}>{ optionB }</button>
    {/if}
    {#if optionC}
      <button on:click={() => {dispatch("message", "c")}}>{ optionC }</button>
    {/if}
    {#if optionD}
      <button on:click={() => {dispatch("message", "d")}}>{ optionD }</button>
    {/if}
    </section>
  {/if}

  <svelte:window on:keydown|preventDefault={skipDialogueIfEnterKeyPressed} />

  <style>
  </style>