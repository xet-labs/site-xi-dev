/**
 * Injects HTML content from external files into elements marked with a specific attribute.
 * 
 * Example:
 *   <div cntnt-src="/partials/header.html"></div>
 *
 * @param {string} attrName - Attribute to look for (e.g., "cntnt-src")
 * @param {Object} [options]
 * @param {"replace"|"append"|"prepend"} [options.mode="replace"] - How to insert content
 * @param {boolean} [options.observe=true] - Watch for dynamically added elements
 * @param {boolean} [options.cache=true] - Cache fetched HTML content
 */
function injectContent(attrName = "cntnt-src", options = {}) {
    const { mode = "replace", observe = true, cache = true } = options;
    const cacheMap = new Map();

    async function processElement(el) {
        const urlAttr = el.getAttribute(attrName);
        if (!urlAttr || el.dataset.cntntLoaded === "true") return;

        // Resolve relative to current location
        const url = new URL(urlAttr, location.href).href;
        el.setAttribute("data-cntnt-status", "loading");

        try {
            let content;
            if (cache && cacheMap.has(url)) {
                content = cacheMap.get(url);
            } else {
                const res = await fetch(url, { cache: cache ? "force-cache" : "no-store" });
                if (!res.ok) throw new Error(`HTTP ${res.status}`);
                content = await res.text();
                if (cache) cacheMap.set(url, content);
            }

            if (mode === "append") el.insertAdjacentHTML("beforeend", content);
            else if (mode === "prepend") el.insertAdjacentHTML("afterbegin", content);
            else el.innerHTML = content;

            el.dataset.cntntLoaded = "true";
            el.setAttribute("data-cntnt-status", "done");
        } catch (err) {
            console.warn(`Failed to inject [${attrName}] ${url}:`, err);
            el.setAttribute("data-cntnt-status", "error");
        }
    }

    function scan(root = document) {
        root.querySelectorAll(`[${attrName}]`).forEach(processElement);
    }

    // Initial load
    scan();

    // Watch for new elements
    if (observe) {
        const observer = new MutationObserver(mutations => {
            for (const m of mutations) {
                m.addedNodes.forEach(node => {
                    if (node.nodeType !== 1) return;
                    if (node.hasAttribute?.(attrName)) processElement(node);
                    else if (node.querySelectorAll) scan(node);
                });
            }
        });
        observer.observe(document.body, { childList: true, subtree: true });
    }
}

/**
 * Register auto-init of multiple attributes.
 * Example:
 *   injectContent.register("cntnt-src");
 *   injectContent.register("data-include", { mode: "append" });
 */
injectContent.register = function (attrName, options) {
    document.addEventListener("DOMContentLoaded", () => injectContent(attrName, options));
};

// Auto-init default attribute
document.addEventListener("DOMContentLoaded", () => injectContent("cntnt-src"));
