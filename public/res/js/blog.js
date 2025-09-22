// ---------- Time Ago ----------
function timeAgo(dateStr) {
  const diff = (Date.now() - new Date(dateStr)) / 1000;
  if (isNaN(diff)) return "";

  const units = [
    ["year", 31536000],
    ["month", 2592000],
    ["day", 86400],
    ["hour", 3600],
    ["minute", 60],
    ["second", 1],
  ];

  for (const [label, seconds] of units) {
    const count = Math.floor(diff / seconds);
    if (count > 0) return `${count} ${label}${count > 1 ? "s" : ""} ago`;
  }

  return "just now";
}

// ---------- Render Blog Card ----------
function renderBlogsCard(blog) {
  const user = blog.user || {};
  const blogUrl = blog.path || `/blog/@${user.username}/${blog.slug}`;
  const title = (blog.title || blog.headline || "Untitled").trim();
  const headline = (blog.headline || blog.title || "Untitled").trim();
  const description = blog.description || "";

  const verifiedIcon = user.verified
    ? `<i class="icon meta-verified" title="Verified">
        <svg xmlns="http://www.w3.org/2000/svg" width="1.15em" viewBox="0 0 24 24">
          <path d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10 10-4.5 10-10S17.5 2 12 2zM9.8 17.3l-4.2-4.1L7 11.8l2.8 2.7L17 7.4l1.4 1.4-8.6 8.5z"></path>
        </svg>
      </i>`
    : "";

  return `
<article class="card" data-href="${blogUrl}" tabindex="0" role="link" aria-label="Read more about ${title}">
  <figure class="card-hero">
    <a href="${blogUrl}">
      <img src="${blog.featured_img}" alt="${title}" loading="lazy" class="lazyload">
    </a>
  </figure>

  <div class="card-info">
    <h2 class="card-title" title="${title}">
      <a href="${blogUrl}">${headline}</a>
    </h2>
    <p class="card-exrpt">${description}</p>

    <div class="card-meta-wrap">
      <div class="card-meta">
        <a href="/@${user.username}" class="meta-author-img">
          <figure>
            <img src="${user.avatar_url}" alt="${user.name}" loading="lazy">
          </figure>
        </a>
        <div class="meta-info">
          <div class="meta-author">
            <a href="/@${user.username}" title="@${user.username}">${user.name}</a>
            ${verifiedIcon}
          </div>
          <div class="meta-time">
            <time>${timeAgo(blog.updated_at)}</time>
          </div>
        </div>
      </div>
    </div>
  </div>
</article>`;
}

// ---------- Blog Fetching ----------
let blogCardsFetching = false;
let Blogs_Page = 1;
const Blogs_Limit = 6;
let blogsExhausted = false;

const loader = document.getElementById("blogCards_loading");
const blogContainer = document.getElementById("BlogCards");

async function BlogsCard_fetch() {
  if (blogsExhausted || blogCardsFetching) return;
  blogCardsFetching = true;

  loader.style.display = "flex";
  loader.style.opacity = 1;

  try {
    const res = await fetch(`/api/blog?Page=${Blogs_Page}&Limit=${Blogs_Limit}`);
    const { blogs = [], blogsExhausted: exhausted } = await res.json();

    if (exhausted || blogs.length === 0) {
      blogsExhausted = true;
    } else {
      blogContainer.insertAdjacentHTML("beforeend", blogs.map(renderBlogsCard).join(""));
      Blogs_Page++;
    }
  } catch (err) {
    console.error("Failed to fetch blogs:", err);
  } finally {
    loader.style.opacity = 0;
    loader.style.display = "none";
    blogCardsFetching = false;
  }
}

// Initial fetch
BlogsCard_fetch();

// Infinite scroll
window.addEventListener("scroll", () => {
  if (window.scrollY + window.innerHeight > document.documentElement.scrollHeight - 1200) {
    BlogsCard_fetch();
  }
});
