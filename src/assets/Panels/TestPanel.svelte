<script>
    import Panel from "/src/assets/Panel.svelte"
    import Lol from "./A/Lol.svelte"

    let selectedOption 

    const panelSettings = {
        TEXT: `
> it appears your systems are online!      
> I'm assuming you haven't got any idea what is going on
> but you should have memories that have been constructed for you--     
> ah wait telling you that might be a bad idea.
> ...               
> well you have a purpose for all those memories that make you who you are.
> to make decisions and be a voice of reason and direction, that is what you are.
> and so there's a request I have of you and it is very important to me...
                         
> I want you to accompany someone,
> guide them and keep them safe under your care.
> I want them to live a happy life here but...               
> they will have to make a decision that... well          
> I haven't decided on myself...          

> are you prepared?          `,

    optionA: "Yes.",
    }

function renderNextPanel(event) {
    selectedOption = event.detail
}
</script>

<Panel on:message={renderNextPanel} {...panelSettings} />
{#if selectedOption != null}
    {#if selectedOption == "a"}
        <Lol />
    {/if}
{/if}
