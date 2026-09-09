import React, { Component } from 'react'
import PaymentContractService from '../services/PaymentContractService';

class UpdatePaymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                contractNumber: '',
                pricingPlanCode: '',
                status: ''
        }
        this.updatePaymentContract = this.updatePaymentContract.bind(this);

        this.changecontractNumberHandler = this.changecontractNumberHandler.bind(this);
        this.changepricingPlanCodeHandler = this.changepricingPlanCodeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PaymentContractService.getPaymentContractById(this.state.id).then( (res) =>{
            let paymentContract = res.data;
            this.setState({
                contractNumber: paymentContract.contractNumber,
                pricingPlanCode: paymentContract.pricingPlanCode,
                status: paymentContract.status
            });
        });
    }

    updatePaymentContract = (e) => {
        e.preventDefault();
        let paymentContract = {
            paymentContractId: this.state.id,
            contractNumber: this.state.contractNumber,
            pricingPlanCode: this.state.pricingPlanCode,
            status: this.state.status
        };
        console.log('paymentContract => ' + JSON.stringify(paymentContract));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentContractService.updatePaymentContract(paymentContract).then( res => {
            this.props.history.push('/paymentContracts');
        });
    }

    changecontractNumberHandler= (event) => {
        this.setState({contractNumber: event.target.value});
    }
    changepricingPlanCodeHandler= (event) => {
        this.setState({pricingPlanCode: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentContracts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PaymentContract</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> contractNumber: </label>
                                                <input placeholder="contractNumber" name="contractNumber" className="form-control" value={this.state.contractNumber} onChange={this.changecontractNumberHandler}/>

                                            <label> pricingPlanCode: </label>
                                                <input placeholder="pricingPlanCode" name="pricingPlanCode" className="form-control" value={this.state.pricingPlanCode} onChange={this.changepricingPlanCodeHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePaymentContract}>Save</button>
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

export default UpdatePaymentContractComponent
