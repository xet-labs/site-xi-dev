function timeAgo(dateStr) {
  const diffInSeconds = Math.floor((Date.now() - new Date(dateStr)) / 1000);
  if (isNaN(diffInSeconds)) return "";

  const units = [
    { label: "year",   seconds: 31536000 },
    { label: "month",  seconds: 2592000 },
    { label: "day",    seconds: 86400 },
    { label: "hour",   seconds: 3600 },
    { label: "minute", seconds: 60 },
    { label: "second", seconds: 1 },
  ];

  for (const unit of units) {
    const count = Math.floor(diffInSeconds / unit.seconds);
    if (count >= 1) {
      return `${count} ${unit.label}${count > 1 ? "s" : ""} ago`;
    }
  }

  return "just now";
}

function renderBlogsCard(blog) {
  const user = blog.user || {};
  const blogUrl = blog.path || `/blog/@${user.username}/${blog.slug}`;
  const blogTitle = blog.title?.trim() || blog.headline?.trim() || "Untitled";
  const blogHeadline = blog.headline?.trim() || blog.title?.trim() || "Untitled";
  const userVerified = user.verified
    ? `<i class="icon meta-verified" title="Verified">
        <svg xmlns="http://www.w3.org/2000/svg" width="1.15em" viewBox="0 0 24 24">
          <path d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10 10-4.5 10-10S17.5 2 12 2zM9.8 17.3l-4.2-4.1L7 11.8l2.8 2.7L17 7.4l1.4 1.4-8.6 8.5z"></path>
        </svg>
      </i>` : "";

  return `
<article class="card" data-href="${blogUrl}" tabindex="0" role="link" aria-label="Read more about ${blog.title || blog.short_title}">
  <figure class="card-hero">
    <a href="${blogUrl}">
      <img src="${blog.featured_img}" alt="${blog.title}" loading="lazy" class="lazyload">
    </a>
  </figure>

  <div class="card-info">
    <h2 class="card-title" title="${blogTitle}">
      <a href="${blogUrl}">${blogHeadline}</a>
    </h2>
    <p class="card-exrpt">${blog.description}</p>

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
            ${userVerified}
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

let blogCardsFetching = false;
let Blogs_Page = 1;
const Blogs_Limit = 6;
let blogsExhausted = false;

async function BlogsCard_fetch() {
  if (blogsExhausted || blogCardsFetching) return;
  blogCardsFetching = true;

  const loader = document.getElementById("blogCards_loading");
  loader.style.display = "block";
  loader.style.opacity = 1;

  try {
    const res = await fetch(`/api/blog?Page=${Blogs_Page}&Limit=${Blogs_Limit}`);
    const response = await res.json();

    const blogs = response.blogs || [];
    if (response.blogsExhausted || blogs.length === 0) {
      blogsExhausted = true;
      loader.style.display = "none";
    } else {
      document.getElementById("BlogCards").insertAdjacentHTML(
        "beforeend",
        blogs.map(renderBlogsCard).join("")
      );
      Blogs_Page++;
    }
  } catch (err) {
    console.error("Failed to fetch blogs:", err);
  }

  loader.style.opacity = 0;
  loader.style.display = "none";
  blogCardsFetching = false;
}

// Initial fetch
BlogsCard_fetch();

// Infinite scroll with vanilla JS
window.addEventListener("scroll", () => {
  if (window.scrollY + window.innerHeight > document.documentElement.scrollHeight - 1200) {
    if (!blogCardsFetching) BlogsCard_fetch();
  }
});
