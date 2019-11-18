const jwt = localStorage.getItem("jwt");
const email = localStorage.getItem("email");
const role = localStorage.getItem("role");

if(jwt === null || email === null || role === null){
    $('#navigation').html('');
    $('#navigation').append('<a class="text-muted" href="/contact">Contact Us</a>');

    $('#auth-box').append('<a class="btn btn-sm btn-outline-secondary" href="/sign-in">Sign In</a>');
}
else{
    $('#navigation').html('');
    $('#navigation').append('<a class="text-muted" href="/manage-articles">Manage Articles</a>');
    $('#navigation').append('<a class="text-muted" href="/see-feedbacks">See Feedbacks</a>');

    $('#auth-box').append('<a class="btn btn-sm btn-outline-secondary" href="#" id="sign-out">Sign Out</a>');
}

$("#sign-out").on("click", function (e) {
    e.preventDefault();
    localStorage.clear();
    window.location = '/';
});