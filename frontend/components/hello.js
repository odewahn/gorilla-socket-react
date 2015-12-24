var React = require('react')

module.exports = React.createClass({
  getInitialState: function() {
    var serverSocket = new WebSocket("ws://localhost:3000/long")
    return {
      serverSocket: serverSocket,
      response: "Nothing yet!",
      command: "nothing"
    }
  },
  componentDidMount: function() {
    var _this = this
    //this.state.serverSocket.send("frontend has started")
    // Write message on receive
    this.state.serverSocket.onmessage = function(e) {
      console.log("got  message back")
      _this.setState({response: e.data})
    };
  },
  setCommand: function(e) {
    this.setState({command: e.target.value})
  },
  startLongProcess: function() {
    console.log("Sending ", this.state.command)
    this.state.serverSocket.send(this.state.command)
  },
  render: function() {
    return (
      <div>
        <h1>Hello, World!</h1>
        <input type="text" name="command" onChange={this.setCommand}/>
        <button onClick={this.startLongProcess}>Do Something</button>
        <br/>
        Here's something you can muck around with: <input type="text" />
        <br/>
        {this.state.response}
      </div>
    )
  }
})
