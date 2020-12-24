new Vue({
el: '#app',

data: {
    ws: null, // the websocket
    newMsg: '',
    chatContent: '',
    email: null,
    username: null,
    joined: false
},
})