<script>
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher()

    export let TEXT, optionA, optionB, optionC, optionD
  
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
    <section id="story-options">
      <button on:click={() => {dispatch("message", "a")}}>A: { optionA }</button>
      <button on:click={() => {dispatch("message", "b")}}>B: { optionB }</button>
      <button on:click={() => {dispatch("message", "c")}}>C: { optionC }</button>
      <button on:click={() => {dispatch("message", "d")}}>D: { optionD }</button>
    </section>
  {/if}

  <svelte:window on:keydown|preventDefault={skipDialogueIfEnterKeyPressed} />

  <style>
  </style>