import React, { Component } from 'react'
import PaymentProcessorService from '../services/PaymentProcessorService'

class ListPaymentProcessorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                paymentProcessors: []
        }
        this.addPaymentProcessor = this.addPaymentProcessor.bind(this);
        this.editPaymentProcessor = this.editPaymentProcessor.bind(this);
        this.deletePaymentProcessor = this.deletePaymentProcessor.bind(this);
    }

    deletePaymentProcessor(id){
        PaymentProcessorService.deletePaymentProcessor(id).then( res => {
            this.setState({paymentProcessors: this.state.paymentProcessors.filter(paymentProcessor => paymentProcessor.paymentProcessorId !== id)});
        });
    }
    viewPaymentProcessor(id){
        this.props.history.push(`/view-paymentProcessor/${id}`);
    }
    editPaymentProcessor(id){
        this.props.history.push(`/add-paymentProcessor/${id}`);
    }

    componentDidMount(){
        PaymentProcessorService.getPaymentProcessors().then((res) => {
            this.setState({ paymentProcessors: res.data});
        });
    }

    addPaymentProcessor(){
        this.props.history.push('/add-paymentProcessor/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PaymentProcessor List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPaymentProcessor}> Add PaymentProcessor</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ProcessorCode </th>
                                    <th> NetworkSupport </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.paymentProcessors.map(
                                        paymentProcessor => 
                                        <tr key = {paymentProcessor.paymentProcessorId}>
                                             <td> { paymentProcessor.name } </td>
                                             <td> { paymentProcessor.processorCode } </td>
                                             <td> { paymentProcessor.networkSupport } </td>
                                             <td>
                                                 <button onClick={ () => this.editPaymentProcessor(paymentProcessor.paymentProcessorId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePaymentProcessor(paymentProcessor.paymentProcessorId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPaymentProcessor(paymentProcessor.paymentProcessorId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListPaymentProcessorComponent
