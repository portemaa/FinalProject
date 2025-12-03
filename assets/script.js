//Password Checker

// Show a message when the page loads
window.onload = function() {
    console.log("Password Checker app loaded.");
};

// Visual if password field is empty before submit
function validateForm() {
    var input = document.getElementById("password");
    if (input && input.value.trim() === "") {
        alert("Please enter a password before checking.");
        return false; // halts the form submission
    }
    return true;
}
