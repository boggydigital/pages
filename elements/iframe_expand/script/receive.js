window.addEventListener("message", e => {
    msg = JSON.parse(e.data);
    let iframe = document.querySelector(`iframe#${msg.context}`)
    iframe.style.height = msg.height + 0 + "px";
    iframe.classList.remove("loading")
});
