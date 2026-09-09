import React, { Component } from 'react'
import PaymentContractService from '../services/PaymentContractService';

class CreatePaymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                contractNumber: '',
                pricingPlanCode: '',
                status: ''
        }
        this.changecontractNumberHandler = this.changecontractNumberHandler.bind(this);
        this.changepricingPlanCodeHandler = this.changepricingPlanCodeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PaymentContractService.getPaymentContractById(this.state.id).then( (res) =>{
                let paymentContract = res.data;
                this.setState({
                    contractNumber: paymentContract.contractNumber,
                    pricingPlanCode: paymentContract.pricingPlanCode,
                    status: paymentContract.status
                });
            });
        }        
    }
    saveOrUpdatePaymentContract = (e) => {
        e.preventDefault();
        let paymentContract = {
                paymentContractId: this.state.id,
                contractNumber: this.state.contractNumber,
                pricingPlanCode: this.state.pricingPlanCode,
                status: this.state.status
            };
        console.log('paymentContract => ' + JSON.stringify(paymentContract));

        // step 5
        if(this.state.id === '_add'){
            paymentContract.paymentContractId=''
            PaymentContractService.createPaymentContract(paymentContract).then(res =>{
                this.props.history.push('/paymentContracts');
            });
        }else{
            PaymentContractService.updatePaymentContract(paymentContract).then( res => {
                this.props.history.push('/paymentContracts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PaymentContract</h3>
        }else{
            return <h3 className="text-center">Update PaymentContract</h3>
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
                                            <label> contractNumber:&emsp; </label>
                                                <input placeholder="contractNumber" name="contractNumber" className="form-control" value={this.state.contractNumber} onChange={this.changecontractNumberHandler}/>

                                            <label> pricingPlanCode:&emsp; </label>
                                                <input placeholder="pricingPlanCode" name="pricingPlanCode" className="form-control" value={this.state.pricingPlanCode} onChange={this.changepricingPlanCodeHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePaymentContract}>Save</button>
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

export default CreatePaymentContractComponent
