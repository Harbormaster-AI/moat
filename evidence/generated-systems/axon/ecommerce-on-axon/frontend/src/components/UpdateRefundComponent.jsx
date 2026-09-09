import React, { Component } from 'react'
import RefundService from '../services/RefundService';

class UpdateRefundComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                refundNumber: '',
                amount: '',
                reason: '',
                createdAt: '',
                status: ''
        }
        this.updateRefund = this.updateRefund.bind(this);

        this.changerefundNumberHandler = this.changerefundNumberHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        RefundService.getRefundById(this.state.id).then( (res) =>{
            let refund = res.data;
            this.setState({
                refundNumber: refund.refundNumber,
                amount: refund.amount,
                reason: refund.reason,
                createdAt: refund.createdAt,
                status: refund.status
            });
        });
    }

    updateRefund = (e) => {
        e.preventDefault();
        let refund = {
            refundId: this.state.id,
            refundNumber: this.state.refundNumber,
            amount: this.state.amount,
            reason: this.state.reason,
            createdAt: this.state.createdAt,
            status: this.state.status
        };
        console.log('refund => ' + JSON.stringify(refund));
        console.log('id => ' + JSON.stringify(this.state.id));
        RefundService.updateRefund(refund).then( res => {
            this.props.history.push('/refunds');
        });
    }

    changerefundNumberHandler= (event) => {
        this.setState({refundNumber: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/refunds');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Refund</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> refundNumber: </label>
                                                <input placeholder="refundNumber" name="refundNumber" className="form-control" value={this.state.refundNumber} onChange={this.changerefundNumberHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> createdAt: </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Requested
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Declined
                      </option>
                      <option name="Status" className="form-control" >
                          Processed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRefund}>Save</button>
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

export default UpdateRefundComponent
