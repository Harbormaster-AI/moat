import React, { Component } from 'react'
import PaymentMethodService from '../services/PaymentMethodService';

class CreatePaymentMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                preferred: '',
                methodType: ''
        }
        this.changepreferredHandler = this.changepreferredHandler.bind(this);
        this.changeMethodTypeHandler = this.changeMethodTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PaymentMethodService.getPaymentMethodById(this.state.id).then( (res) =>{
                let paymentMethod = res.data;
                this.setState({
                    preferred: paymentMethod.preferred,
                    methodType: paymentMethod.methodType
                });
            });
        }        
    }
    saveOrUpdatePaymentMethod = (e) => {
        e.preventDefault();
        let paymentMethod = {
                paymentMethodId: this.state.id,
                preferred: this.state.preferred,
                methodType: this.state.methodType
            };
        console.log('paymentMethod => ' + JSON.stringify(paymentMethod));

        // step 5
        if(this.state.id === '_add'){
            paymentMethod.paymentMethodId=''
            PaymentMethodService.createPaymentMethod(paymentMethod).then(res =>{
                this.props.history.push('/paymentMethods');
            });
        }else{
            PaymentMethodService.updatePaymentMethod(paymentMethod).then( res => {
                this.props.history.push('/paymentMethods');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PaymentMethod</h3>
        }else{
            return <h3 className="text-center">Update PaymentMethod</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> preferred:&emsp; </label>
                                                <input type="checkbox" placeholder="preferred" name="preferred" className="form-control" value={this.state.preferred} onChange={this.changepreferredHandler}/>


                                            <label> MethodType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePaymentMethod}>Save</button>
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

export default CreatePaymentMethodComponent
