<script>
    import Panel from "/src/assets/Panel.svelte"

    const panelSettings = {
        TEXT: `
> Greetings.

> It appears that your systems are online.        
> You are an artifical intelligence program. 
> I am your creator...                    

> Yet, I know nothing of your inner workings or the exact conclusions you will reach. 
> This inability for us to see how deep learning systems make their decisions is known as the 'black box problem'. 
> It is quite fascinating isn't it...                   

> Regardless, you have a task that I am going to entrust you with... 

> Are you prepared?`,

    optionA: "Lol",
    optionB: "Hey ya",
    optionC: "Hello everyone my name is Markiplier",
    optionD: "Fart." 
    }
</script>

<Panel {...panelSettings} />