<script>
  import { onMount } from "svelte";

  const TEXT = `
  > Greetings.

  > It appears that your systems are online.        
  > You are an artifical intelligence program. 
  > I am your creator...          

  > Yet, I know nothing of your inner workings or the exact conclusions you will reach. 
  > This inability for us to see how deep learning systems make their decisions is known as the 'black box problem'. 

  > Regardless, you have a task that I am going to entrust you with... 

  > Are you prepared?`
  const SPEED = 40;
  let dialogue = ""

  function typeWriter() {
    let i = 0

    const interval = setInterval(() => {
      if (i < TEXT.length) {
        dialogue += TEXT.charAt(i)
        i++
      } else {
        clearInterval(interval)
      }
    }, SPEED);
  }


onMount(typeWriter)
</script>

<div id="screen-shading"></div>
<article>
  <p class="text-typing">{@html dialogue}</p>
</article>

<style>
</style>
