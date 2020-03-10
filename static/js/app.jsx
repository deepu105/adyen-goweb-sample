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

async function payments(value, currency, paymentData, reference, clientIP) {
  const response = await fetch("/api/payments", {
    method: "POST",
    body: JSON.stringify({
      returnUrl: "http://localhost:3000/redirect",
      channel: CHANNEL,
      countryCode: COUNTRY,
      amount: { currency, value },
      paymentMethod: paymentData.paymentMethod,
      browserInfo: paymentData.browserInfo,
      billingAddress: paymentData.billingAddress,
      shopperIP: clientIP,
      // @ts-ignore
      origin: document.location.origin,
      reference
    })
  });
  return await response.json();
}

async function paymentDetails(paymentData, details) {
  const response = await fetch("/api/paymentDetails", {
    method: "POST",
    body: JSON.stringify({
      paymentData,
      details
    })
  });
  return await response.json();
}

async function clientIP() {
  const response = await fetch("/api/clientIP");
  return await response.text();
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
      clientIP: "",
      amount: 113.5,
      currency: "EUR",
      valid: false,
      paid: false,
      paymentData: {},
      paymentDetails: {},
      paymentRes: {}
    };
    this.onChange = this.onChange.bind(this);
    this.onAdditionalDetails = this.onAdditionalDetails.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
    this.payNow = this.payNow.bind(this);
  }

  componentDidMount() {
    const { amount, currency } = this.state;
    // @ts-ignore
    let params = new URL(document.location).searchParams;
    const paymentRes = {
      pspReference: params.get("PspReference"),
      resultCode: params.get("ResultCode"),
      refusalReason: params.get("RefusalReason")
    };
    if (paymentRes.pspReference) {
      // TODO decode params
      // pspReference = decodeURI(pspReference);
      // resultCode = decodeURI(resultCode);
      // refusalReason = decodeURI(refusalReason);
      // @ts-ignore
      this.setState({ paid: true, paymentRes });
    } else {
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
    clientIP().then(clientIP => this.setState({ clientIP }));
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
    this.checkout
      .create(id, {
        hasHolderName: true,
        holderNameRequired: true,
        billingAddressRequired: true
      })
      .mount(this.paymentContainer.current);
  }

  payNow() {
    const { amount, currency, paymentData, clientIP } = this.state;
    // @ts-ignore
    sessionStorage.clear();
    payments(amount, currency, paymentData, `${Date.now()}`, clientIP).then(res => {
      if (res.action) {
        this.checkout.createFromAction(res.action).mount(this.idealAction.current);
      } else {
        this.setState({ paid: true, paymentRes: res });
      }
    });
  }

  render() {
    const { paid, valid, paymentMethodVal, paymentRes } = this.state;
    return (
      <div className="container d-flex justify-content-center">
        <div className="col-8 jumbotron">
          <div className="text-center">
            <h1>
              <a href="/">Adyen Checkout</a>
            </h1>
            <p>Payments made easy</p>
          </div>
          <hr />
          {paid ? (
            paymentRes && paymentRes.resultCode ? (
              <React.Fragment>
                <h4 className="title mb-3">Payment {paymentRes.resultCode}</h4>
                <div>{paymentRes.refusalReason ? <span>Reason: {paymentRes.refusalReason}</span> : ""}</div>
                <div>Payment reference: {paymentRes.pspReference}</div>
              </React.Fragment>
            ) : (
              <h4 className="title mb-3">Payment is being processed</h4>
            )
          ) : (
            <React.Fragment>
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
                  <label className="form-check-label" htmlFor="card">
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
              <div>
                <button type="button" className="btn btn-primary" disabled={!valid} onClick={this.payNow}>
                  Pay now
                </button>
              </div>
            </React.Fragment>
          )}
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
