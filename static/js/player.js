
let lastEl,
    player = null;

function playsong(el, url) {
    if (player != null) {
        player.pause();
        player = null;
    }
    lastEl = el;
    player = new Audio(url);
    player.volume = 0.05;
    player.play()
}