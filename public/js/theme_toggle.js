const getPreferredTheme = () => {
    return window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light";
};

window.toggleTheme = () => {
    const current = document.documentElement.getAttribute("theme");
    const newTheme = current === "dark" ? "light" : "dark";
    document.documentElement.setAttribute("theme", newTheme);
    updateToggleIcon(newTheme);
};

const updateToggleIcon = (theme) => {
    const icon = document.querySelector(".theme-toggle i");
    if (icon) {
        icon.className = theme === "dark" ? "nf nf-fa-sun" : "nf nf-fa-moon";
    }
};

document.addEventListener("DOMContentLoaded", () => {
    document.documentElement.setAttribute("theme", getPreferredTheme());
    updateToggleIcon(getPreferredTheme());
});
