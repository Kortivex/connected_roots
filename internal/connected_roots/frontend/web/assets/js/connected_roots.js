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
});
