var React = require('react')

module.exports = React.createClass({
  getInitialState: function() {
    var serverSocket = new WebSocket("ws://localhost:3000/db")
    return {
      serverSocket: serverSocket,
      Config: {
        Username: "",
        Password: ""
      }
    }
  },
  componentDidMount: function() {
    var _this = this
    // Write message on receive
    this.state.serverSocket.onmessage = function(e) {
      console.log("got  message back")
      _this.setState({response: e.data})
    };
  },
  setField: function(e) {
    var s = this.state.Config
    s[e.target.name] = e.target.value
    this.setState({Config: s})
  },
  saveValues: function() {
    console.log("Sending ", this.state.Config)
    var msg = {
      key: "LaunchbotCredentials",
      value: this.state.Config
    }
    this.state.serverSocket.send(JSON.stringify(msg))
  },
  render: function() {
    return (
      <div>
        <h1>Hello, World!</h1>
        Username: <input type="text" name="Username" onChange={this.setField}/>
        <br/>
        Password: <input type="text" name="Password" onChange={this.setField}/>
        <br/>
        <button onClick={this.saveValues}>Do Something</button>
      </div>
    )
  }
})
