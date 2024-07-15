


// Function to extract account ID from URL query parameter
function getAccountIdFromUrl() {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get('account_id');
}


// window.onload = () => {
//     fetchAccounts();
//     document.getElementById('add-account-form').addEventListener('submit', addAccount);
// };
