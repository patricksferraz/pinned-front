function countWords(elem) {
  let words = document.getElementById('words');
  words.innerText = elem.getAttribute('maxlength') - elem.value.length;
}

function toast(msg, type) {
  let types = {
    success: '#00c851',
    error: '#ff4444',
    info: '#0099cc',
    warning: '#ffbb33',
  };

  Toastify({
    text: msg,
    duration: 3000,
    // destination: 'https://github.com/apvarun/toastify-js',
    newWindow: true,
    close: true,
    gravity: 'bottom', // `top` or `bottom`
    position: 'right', // `left`, `center` or `right`
    stopOnFocus: true, // Prevents dismissing of toast on hover
    style: {
      background: types[type],
    },
    onClick: function () {}, // Callback after click
  }).showToast();
}
