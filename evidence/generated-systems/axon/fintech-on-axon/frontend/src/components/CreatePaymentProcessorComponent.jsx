import React, { Component } from 'react'
import PaymentProcessorService from '../services/PaymentProcessorService';

class CreatePaymentProcessorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                processorCode: '',
                networkSupport: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeprocessorCodeHandler = this.changeprocessorCodeHandler.bind(this);
        this.changenetworkSupportHandler = this.changenetworkSupportHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PaymentProcessorService.getPaymentProcessorById(this.state.id).then( (res) =>{
                let paymentProcessor = res.data;
                this.setState({
                    name: paymentProcessor.name,
                    processorCode: paymentProcessor.processorCode,
                    networkSupport: paymentProcessor.networkSupport
                });
            });
        }        
    }
    saveOrUpdatePaymentProcessor = (e) => {
        e.preventDefault();
        let paymentProcessor = {
                paymentProcessorId: this.state.id,
                name: this.state.name,
                processorCode: this.state.processorCode,
                networkSupport: this.state.networkSupport
            };
        console.log('paymentProcessor => ' + JSON.stringify(paymentProcessor));

        // step 5
        if(this.state.id === '_add'){
            paymentProcessor.paymentProcessorId=''
            PaymentProcessorService.createPaymentProcessor(paymentProcessor).then(res =>{
                this.props.history.push('/paymentProcessors');
            });
        }else{
            PaymentProcessorService.updatePaymentProcessor(paymentProcessor).then( res => {
                this.props.history.push('/paymentProcessors');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeprocessorCodeHandler= (event) => {
        this.setState({processorCode: event.target.value});
    }
    changenetworkSupportHandler= (event) => {
        this.setState({networkSupport: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentProcessors');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PaymentProcessor</h3>
        }else{
            return <h3 className="text-center">Update PaymentProcessor</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> processorCode:&emsp; </label>
                                                <input placeholder="processorCode" name="processorCode" className="form-control" value={this.state.processorCode} onChange={this.changeprocessorCodeHandler}/>

                                            <label> networkSupport:&emsp; </label>
                                                <input placeholder="networkSupport" name="networkSupport" className="form-control" value={this.state.networkSupport} onChange={this.changenetworkSupportHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePaymentProcessor}>Save</button>
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

export default CreatePaymentProcessorComponent
