const copyContent = async () => {
    let text = document.getElementById('resultTab').innerHTML;
    let button = document.getElementById('copyButton');
    try {
        await navigator.clipboard.writeText(text);
        button.innerHTML = "Copied!";
        console.log('Content copied to clipboard')
    } catch (err) {
        console.error('Failed to copy: ', err);
    }
}

const toggleEye = (e) => {
  try {
    const parent = e.target.closest(".passwordContainer");
    if (!parent) return;
    
    const classList = parent.querySelector(".passwordEye").classList;
    const passwordField = parent.querySelector(".passwordField");

    if (classList && passwordField) {
      // Toggle password visibility
      const type = classList.contains("show") ? "password" : "text";
      passwordField.setAttribute("type", type);

      // Toggle the eye icon visibility
      classList.toggle("show");
      classList.toggle("hide");
    }
  } catch (err) {
    console.error("Failed to toggle password visibility.", err);
  }
};

document.querySelectorAll(".passwordEye").forEach((input) => {
  input.addEventListener("click", toggleEye);
});

