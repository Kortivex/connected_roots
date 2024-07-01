document.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll('.nav.collapse.parent').forEach(navCollapseParent => {
        if (navCollapseParent.querySelector('li a.active')) {
            navCollapseParent.classList.add('show');
            const toggleLink = navCollapseParent.closest('.parent-wrapper')?.previousElementSibling;
            if (toggleLink?.tagName.toLowerCase() === 'a') {
                toggleLink.setAttribute('aria-expanded', 'true');
                toggleLink.classList.add('collapsed');
            }
        }
    });
    const username = document.getElementById("user-name-surname").textContent;
    document.getElementById("avatar").src = generateAvatar(username);
});

function generateAvatar(text) {
    const canvas = document.createElement("canvas");
    const context = canvas.getContext("2d");

    canvas.width = 200;
    canvas.height = 200;

    // Set default color values
    let foregroundColor =  "#222834";
    let backgroundColor =  "#8B94AD";

    // Extract initials
    const initials = text.split(' ').map(word => word[0].toUpperCase()).join('');

    // Draw background
    context.fillStyle = backgroundColor;
    context.fillRect(0, 0, canvas.width, canvas.height);

    // Draw text
    context.font = "bold 100px Assistant";
    context.fillStyle = foregroundColor;
    context.textAlign = "center";
    context.textBaseline = "middle";
    context.fillText(initials, canvas.width / 2, canvas.height / 2);

    return canvas.toDataURL("image/png");
}
