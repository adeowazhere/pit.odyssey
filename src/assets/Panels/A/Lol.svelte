<script>
    import Panel from "/src/assets/Panel.svelte"

    const panelSettings = {
        TEXT: `
> thats great!
> okay
> so
> ahem.
     
> the boy you will meet goes by...          
<b><i>Pit.</i></b>     
> he sounds like a 12 year old and tells a few tall tales          
> not anything malicious just uh... a teeny bit <i>mischievous</i>
> fortunately, he doesn't like dad jokes such as puns     
> so that should give you some leverage if things get <i>too</i> out of hand.     
> im sure you will get along fine, dont worry too much about it :)
          
> oh right you need to know what he even looks like
> hold on I have a drawing somewhere...
> even without a corporeal form it shouldn't look too bad
> ...                   
> hopefully.
                    
> here!`,

    imageLink: "/src/assets/Images/pit_drawing.png",
    imageAlt: "a drawing of Pit!",
    optionA: "Woah, yummy art! (eat the drawing.)",
    optionB: "Where did the paper come from?",
    optionC: "I would like to know more.",
    optionD: "Okay got it, what's next?" 
    }
</script>

<Panel {...panelSettings} />