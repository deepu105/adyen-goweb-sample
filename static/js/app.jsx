class Home extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
          <h1>Adyen Payments</h1>
          <p>Payments made easy</p>
        </div>
      </div>
    );
  }
}

class App extends React.Component {
  render() {
    return <Home />;
  }
}

ReactDOM.render(<App />, document.getElementById("app"));
