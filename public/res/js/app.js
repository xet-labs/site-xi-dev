// toggle underlay
function toggleUnderlay(name, state) {
  const underlay = document.querySelector(`.underlay[data-name="${name}"]`);
  if (!underlay) return console.warn(`Underlay '${name}' not found`);

  // No state provided → toggle current
  if (state === undefined) {
    underlay.classList.toggle('active');
    return;
  }

  // Force show/hide
  underlay.classList.toggle('active', !!state);
}

// hide show password-field
document.querySelectorAll(".pw-toggle").forEach(btn => {
  btn.onclick = function () {
    const wrapper = this.closest(".input-field");
    const input = wrapper.querySelector(".pw-input");
    const isHidden = input.type === "password";

    input.type = isHidden ? "text" : "password";

    // Toggle active state to switch icons
    this.classList.toggle("active", isHidden);
    this.title = isHidden ? "Hide password" : "Show password";
  };
});

document.addEventListener("DOMContentLoaded", () => {
  // Theme toggle (light/dark) ===
  document.documentElement.classList.add(localStorage.getItem('theme') || 'light');

  const themeSwitch = document.getElementById("id-themeswitch");
  if (themeSwitch) {
    themeSwitch.addEventListener("click", () => {
      document.documentElement.classList.toggle("dark");
      const theme = document.documentElement.classList.contains("dark") ? "dark" : "light";
      localStorage.setItem("theme", theme);
    });
  }


  // Nav hide|reveal on scroll
  let lastScrollTop = 0;
  let ticking = false;
  const header = document.querySelector('header');
  const headerHeight = header.offsetHeight;
  const scrollDelta = 5;

  window.addEventListener('scroll', function () {
    if (!ticking) {
      window.requestAnimationFrame(function () {
        const currentScroll = window.pageYOffset || document.documentElement.scrollTop;

        // Ignore small scroll changes (jitter)
        if (Math.abs(currentScroll - lastScrollTop) <= scrollDelta) {
          ticking = false;
          return;
        }

        // Down scroll & past header -> hide
        if (currentScroll > lastScrollTop && currentScroll > headerHeight) {
          header.style.top = `-${headerHeight}px`;
        } else {
          // Up scroll or at top -> show
          header.style.top = '0';
        }

        lastScrollTop = currentScroll <= 0 ? 0 : currentScroll;
        ticking = false;
      });

      ticking = true;
    }
  });

  // === Fix modal flicker for login/signup switch ===
  const loginBtn = document.getElementById('id-login-btn');
  const signupBtn = document.getElementById('id-signup-btn');
  const loginWrap = document.querySelector('.login-wrap');
  const signupWrap = document.querySelector('.signup-wrap');

  function toggleAuthForm() {
    if (loginBtn?.checked) {
      if (loginWrap) loginWrap.style.display = "flex";
      if (signupWrap) signupWrap.style.display = "";
    } else if (signupBtn?.checked) {
      if (signupWrap) signupWrap.style.display = "flex";
      if (loginWrap) loginWrap.style.display = "";
    }
  }

  if (loginBtn) loginBtn.addEventListener('click', toggleAuthForm);
  if (signupBtn) signupBtn.addEventListener('click', toggleAuthForm);


  // handle login
  document.getElementById('loginForm').addEventListener('submit', function (event) {
    event.preventDefault(); // Prevent default form submission

    // Create a JSON object with the form data
    var formData = {
      email: document.getElementById('loginEmail').value,
      password: document.getElementById('loginPassword').value
    };

    // Send a POST request with JSON data
    fetch('/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
    })
      .then(response => response.json())
      .then(data => {
        if (data.access_token) {
          alert('Login successful!');
          // Do something with the token, e.g., save it or redirect
        } else {
          alert('Login failed: ' + data.error);
        }
      })
      .catch(error => {
        alert('Error: ' + error.message);
      });
  });



});

document.addEventListener("DOMContentLoaded", () => {
  const signupForm = document.getElementById("signupForm");

  signupForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // stop normal form submission

    // Collect form values
    const formData = {
      username: document.getElementById("signupUsername")?.value || "",
      name: document.getElementById("signupName")?.value || "",
      email: document.getElementById("signupEmail").value,
      password: document.getElementById("signupPassword").value,
      confirm_password: document.getElementById("signupConfirmPassword").value,
    };

    // Basic client-side validation
    if (formData.password !== formData.confirm_password) {
      alert("Passwords do not match!");
      return;
    }

    try {
      // Send JSON POST request to your backend
      const response = await fetch("/api/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      const data = await response.json();

      if (response.ok) {
        alert("Signup successful! You can now log in.");
        // Optional: redirect to login page
        // window.location.href = "/login";
      } else {
        // Backend returned an error
        alert("Signup failed: " + (data.error || "Unknown error"));
      }
    } catch (error) {
      alert("Network error: " + error.message);
    }
  });
});

