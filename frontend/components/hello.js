var React = require('react')

module.exports = React.createClass({
  getInitialState: function() {
    return {
      lastTick: "Never!"
    }
  },
  componentDidMount: function() {
    var serversocket = new WebSocket("ws://localhost:3000/pulsar");
    var _this = this
    // Write message on receive
    serversocket.onmessage = function(e) {
      _this.setState({lastTick: e.data})
    };

  },
  render: function() {
    return (
      <div>
        <h1>Hello, World!</h1>
        {this.state.lastTick}
      </div>
    )
  }
})
