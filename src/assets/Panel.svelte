<script>
    import { onMount } from "svelte";

    export let TEXT, optionA, optionB, optionC, optionD

  
    const SPEED = 40;
    let dialogue = ""
    let isDialogueFinished = false
    let i = 0

    let selectedOption = ""
  
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
      <button on:click={() => {selectedOption = "a"}}>A: { optionA }</button>
      <button on:click={() => {selectedOption = "b"}}>A: { optionB }</button>
      <button on:click={() => {selectedOption = "c"}}>A: { optionC }</button>
      <button on:click={() => {selectedOption = "d"}}>A: { optionD }</button>
    </section>
  {/if}

  <svelte:window on:keydown|preventDefault={skipDialogueIfEnterKeyPressed} />

  <style>
  </style>