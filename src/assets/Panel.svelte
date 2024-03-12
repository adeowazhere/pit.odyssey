<script>
    import { onMount } from "svelte";

    export let TEXT, optionA, optionB, optionC, optionD

  
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

    function skipDialogue() {
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
    <button>Skip</button>
  {/if}
  
  {#if isDialogueFinished}
    <section id="story-options">
      <a>A: { optionA }</a>
      <a>B: { optionB }</a>
      <a>C: { optionC }</a>
      <a>D: { optionD }</a>
    </section>
  {/if}

  <svelte:window on:keydown|preventDefault={skipDialogueIfEnterKeyPressed} />

  <style>
  </style>