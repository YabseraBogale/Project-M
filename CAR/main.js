let menu = document.querySelector("#menu-icon");
let navbar = document.querySelector(".navbar");
menu.onclick = () => {
  menu.classList.toggle("bx-x");
  navbar.classList.toggle("active");
};
window.onscroll = () => {
  menu.classList.remove("bx-x");
  navbar.classList.remove("active");
};
const form = document.getElementById("form");
const username = document.getElementById("username");
const phoneno = document.getElementById("phoneno");
const email = document.getElementById("email");
const password = document.getElementById("password");
const confirmpassord = document.getElementById("cpassord");

form.addEventListener("submit", (e) => {
  e.preventDefault();
  validateInput();
});

const setError = (element, message) => {
  const inputcontrol = element.parentElement;
  const errorDisplay = inputControl.querySelector(".error");

  erroeDisplay.innerText = message;
  inputControl.classList.add("error");
  inputControl.classList.remove("success");
};

const setSuccess = (element) => {
  const inputcontrol = element.parentElement;
  const errorDisplay = inputControl.querySelector(".error");

  erroeDisplay.innerText = "";
  inputControl.classList.add("success");
  inputControl.classList.remove("error");
};

const validateInput = () => {
  const usernameValue = username.value.trim();
  const emailValue = email.value.trim();
  const phonenoValue = phoneno.value.trim();
  const passwordValue = password.value.trim();
  const confirmpasswordValue = confirmpassword.value.trim();

  if (usernameValue === "") {
    setError(username, "username is required");
  } else {
    setSuccess(username);
  }
  if (phonenoValue === "") {
    setError(phoneno, "phone number is required");
  } else {
    setSuccess(phoneno);
  }
  if (emailValue === "") {
    setError(email, "email is required");
  } else {
    setSuccess(email);
  }
  if (passwordValue === "") {
    setError(password, "password is required");
  } else if (passwordValue.length < 8) {
    setError(password, "password must be at least 8 character");
  } else {
    setSuccess(passord);
  }
  if (confirmpasswordValue === "") {
    setError(password, "please confirm your password");
  } else if (confirmpassordValue != passwordValue) {
    setError(confirmpassword, "password doesn't match");
  } else {
    setSuccess(confirmpassord);
  }
};
