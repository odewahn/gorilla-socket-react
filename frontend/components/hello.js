var React = require('react')
var $ = require('jquery'); //Use to load data from the DB


module.exports = React.createClass({
  getInitialState: function() {
    return {
      Config: {
        Username: "",
        Password: ""
      }
    }
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
    $.ajax({
      type: "POST",
      url: "/db",
      data: JSON.stringify(msg),
      datatype: "JSON"
    })

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
