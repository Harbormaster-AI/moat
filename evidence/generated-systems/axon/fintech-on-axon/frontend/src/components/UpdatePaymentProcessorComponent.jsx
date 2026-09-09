import React, { Component } from 'react'
import PaymentProcessorService from '../services/PaymentProcessorService';

class UpdatePaymentProcessorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                processorCode: '',
                networkSupport: ''
        }
        this.updatePaymentProcessor = this.updatePaymentProcessor.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeprocessorCodeHandler = this.changeprocessorCodeHandler.bind(this);
        this.changenetworkSupportHandler = this.changenetworkSupportHandler.bind(this);
    }

    componentDidMount(){
        PaymentProcessorService.getPaymentProcessorById(this.state.id).then( (res) =>{
            let paymentProcessor = res.data;
            this.setState({
                name: paymentProcessor.name,
                processorCode: paymentProcessor.processorCode,
                networkSupport: paymentProcessor.networkSupport
            });
        });
    }

    updatePaymentProcessor = (e) => {
        e.preventDefault();
        let paymentProcessor = {
            paymentProcessorId: this.state.id,
            name: this.state.name,
            processorCode: this.state.processorCode,
            networkSupport: this.state.networkSupport
        };
        console.log('paymentProcessor => ' + JSON.stringify(paymentProcessor));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentProcessorService.updatePaymentProcessor(paymentProcessor).then( res => {
            this.props.history.push('/paymentProcessors');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PaymentProcessor</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> processorCode: </label>
                                                <input placeholder="processorCode" name="processorCode" className="form-control" value={this.state.processorCode} onChange={this.changeprocessorCodeHandler}/>

                                            <label> networkSupport: </label>
                                                <input placeholder="networkSupport" name="networkSupport" className="form-control" value={this.state.networkSupport} onChange={this.changenetworkSupportHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePaymentProcessor}>Save</button>
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

export default UpdatePaymentProcessorComponent
