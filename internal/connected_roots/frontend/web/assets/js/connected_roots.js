document.addEventListener("DOMContentLoaded", () => {
    // Add active class to parent.
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

    // Get username.
    let username = document.getElementById("user-name-surname");
    if (username) username = username.textContent;

    // Generate avatar.
    const avatar = document.getElementById("avatar");
    if (avatar) avatar.src = generateAvatar(username);

    // Generate avatar detail profile.
    const avatarDetailProfile = document.getElementById("avatar-detail-profile");
    if (avatarDetailProfile) avatarDetailProfile.src = generateAvatar(username);

    // Show notification.
    const toastEl = document.getElementById('notification-toast');
    if (toastEl) {
        const toast = new bootstrap.Toast(toastEl, {
            autohide: true,
            delay: 5000
        });
        toast.show();
    }
});

function generateAvatar(text) {
    const canvas = document.createElement("canvas");
    const context = canvas.getContext("2d");

    canvas.width = 200;
    canvas.height = 200;

    // Set default color values
    let foregroundColor = "#222834";
    let backgroundColor = "#8B94AD";

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
