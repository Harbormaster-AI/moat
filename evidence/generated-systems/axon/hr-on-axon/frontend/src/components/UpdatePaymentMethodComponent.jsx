import React, { Component } from 'react'
import PaymentMethodService from '../services/PaymentMethodService';

class UpdatePaymentMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                preferred: '',
                methodType: ''
        }
        this.updatePaymentMethod = this.updatePaymentMethod.bind(this);

        this.changepreferredHandler = this.changepreferredHandler.bind(this);
        this.changeMethodTypeHandler = this.changeMethodTypeHandler.bind(this);
    }

    componentDidMount(){
        PaymentMethodService.getPaymentMethodById(this.state.id).then( (res) =>{
            let paymentMethod = res.data;
            this.setState({
                preferred: paymentMethod.preferred,
                methodType: paymentMethod.methodType
            });
        });
    }

    updatePaymentMethod = (e) => {
        e.preventDefault();
        let paymentMethod = {
            paymentMethodId: this.state.id,
            preferred: this.state.preferred,
            methodType: this.state.methodType
        };
        console.log('paymentMethod => ' + JSON.stringify(paymentMethod));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentMethodService.updatePaymentMethod(paymentMethod).then( res => {
            this.props.history.push('/paymentMethods');
        });
    }

    changepreferredHandler= (event) => {
        this.setState({preferred: event.target.value});
    }
    changeMethodTypeHandler= (event) => {
        this.setState({methodType: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentMethods');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PaymentMethod</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> preferred: </label>
                                                <input type="checkbox" placeholder="preferred" name="preferred" className="form-control" value={this.state.preferred} onChange={this.changepreferredHandler}/>


                                            <label> MethodType: </label>
                                                <select value={this.state.methodType} onChange={this.changeMethodTypeHandler}>
                      <option name="MethodType" className="form-control" >
                          DirectDeposit
                      </option>
                      <option name="MethodType" className="form-control" >
                          Check
                      </option>
                      <option name="MethodType" className="form-control" >
                          Cash
                      </option>
                      <option name="MethodType" className="form-control" >
                          InternationalTransfer
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePaymentMethod}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdatePaymentMethodComponent
