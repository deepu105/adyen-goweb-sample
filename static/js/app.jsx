const CHANNEL = "Web";
const COUNTRY = "NL";

async function paymentMethods(value, currency) {
  const response = await fetch("/api/paymentMethods", {
    method: "POST",
    body: JSON.stringify({
      countryCode: COUNTRY,
      channel: CHANNEL,
      amount: { currency, value }
    })
  });
  return await response.json();
}

async function payments(value, currency, paymentMethod, reference) {
  const response = await fetch("/api/payments", {
    method: "POST",
    body: JSON.stringify({
      returnUrl: "http://localhost:3000/",
      channel: CHANNEL,
      countryCode: COUNTRY,
      amount: { currency, value },
      reference,
      paymentMethod
    })
  });
  return await response.json();
}

const configuration = {
  locale: "en_NL",
  environment: "test",
  originKey: "pub.v2.8015833254571517.aHR0cDovL2xvY2FsaG9zdDozMDAw.XYkqfCl6Naupxo3swDluiREx7fjhRhTRu1GYjmo57GY"
};

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.cardPayment = React.createRef();
    this.paymentContainer = React.createRef();
    this.idealAction = React.createRef();
    this.state = {
      amount: 1.12,
      currency: "EUR",
      valid: false,
      paymentData: {},
      paymentDetails: {}
    };
    this.onChange = this.onChange.bind(this);
    this.onAdditionalDetails = this.onAdditionalDetails.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
    this.payNow = this.payNow.bind(this);
  }

  componentDidMount() {
    const { amount, currency } = this.state;
    paymentMethods(amount, currency).then(paymentMethodsResponse => {
      // @ts-ignore
      this.checkout = new AdyenCheckout({
        ...configuration,
        paymentMethodsResponse,
        onAdditionalDetails: this.onAdditionalDetails,
        onChange: this.onChange
      });
    });
  }

  onChange(state) {
    this.setState({
      valid: state.isValid,
      paymentData: state.data
    });
  }

  onAdditionalDetails(state) {
    this.setState({
      paymentDetails: state.data
    });
  }

  handleInputChange(event) {
    const { value, id, name } = event.target;
    this.setState({
      [`${name}Val`]: value,
      [`${name}Id`]: id,
      valid: false,
      paymentData: {}
    });
    this.checkout.create(id).mount(this.paymentContainer.current);
  }

  payNow() {
    const { amount, currency, paymentData } = this.state;
    payments(amount, currency, paymentData.paymentMethod, `${Date.now()}`).then(res => {
      if (res.action) {
        this.checkout.createFromAction(res.action).mount(this.idealAction.current);
      } else {
        //todo show result
        console.log(res);
      }
    });
  }

  render() {
    const { paymentData, valid, paymentMethodVal } = this.state;
    return (
      <div className="container d-flex justify-content-center">
        <div className="col-8 jumbotron">
          <div className="text-center">
            <h1>Adyen Checkout</h1>
            <p>Payments made easy</p>
          </div>
          <hr />
          <h4 className="title mb-3">Select Payment</h4>
          <div className="select-payment mb-5">
            <div className="form-check form-check-inline">
              <input
                className="form-check-input"
                type="radio"
                name="paymentMethod"
                id="ideal"
                value="iDEAL"
                onChange={this.handleInputChange}
              />
              <label className="form-check-label" htmlFor="ideal">
                iDEAL
              </label>
            </div>
            <div className="form-check form-check-inline">
              <input
                className="form-check-input"
                type="radio"
                name="paymentMethod"
                id="card"
                value="Credit Card"
                onChange={this.handleInputChange}
              />
              <label className="form-check-label" htmlFor="creditcard">
                Credit Card
              </label>
            </div>
          </div>
          <div className="mb-5">
            <h4 className="title mb-3">Pay with {paymentMethodVal}</h4>
            <div className="col-12">
              <div ref={this.paymentContainer}></div>
              <div ref={this.idealAction}></div>
            </div>
          </div>
          {/* <div className="row mb-5">
            <div className="col-12">
              <div ref={this.cardPayment}></div>
            </div>
          </div> */}
          <div>
            <button type="button" className="btn btn-primary" disabled={!valid} onClick={this.payNow}>
              Pay now
            </button>
          </div>
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

// @ts-ignore
ReactDOM.render(<App />, document.getElementById("app"));
