import React, { Component } from 'react'
import ReturnRequestService from '../services/ReturnRequestService';

class UpdateReturnRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                returnNumber: '',
                createdAt: '',
                refundAmount: '',
                status: ''
        }
        this.updateReturnRequest = this.updateReturnRequest.bind(this);

        this.changereturnNumberHandler = this.changereturnNumberHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changerefundAmountHandler = this.changerefundAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ReturnRequestService.getReturnRequestById(this.state.id).then( (res) =>{
            let returnRequest = res.data;
            this.setState({
                returnNumber: returnRequest.returnNumber,
                createdAt: returnRequest.createdAt,
                refundAmount: returnRequest.refundAmount,
                status: returnRequest.status
            });
        });
    }

    updateReturnRequest = (e) => {
        e.preventDefault();
        let returnRequest = {
            returnRequestId: this.state.id,
            returnNumber: this.state.returnNumber,
            createdAt: this.state.createdAt,
            refundAmount: this.state.refundAmount,
            status: this.state.status
        };
        console.log('returnRequest => ' + JSON.stringify(returnRequest));
        console.log('id => ' + JSON.stringify(this.state.id));
        ReturnRequestService.updateReturnRequest(returnRequest).then( res => {
            this.props.history.push('/returnRequests');
        });
    }

    changereturnNumberHandler= (event) => {
        this.setState({returnNumber: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changerefundAmountHandler= (event) => {
        this.setState({refundAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/returnRequests');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ReturnRequest</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> returnNumber: </label>
                                                <input placeholder="returnNumber" name="returnNumber" className="form-control" value={this.state.returnNumber} onChange={this.changereturnNumberHandler}/>

                                            <label> createdAt: </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> refundAmount: </label>
                                                <input placeholder="refundAmount" name="refundAmount" className="form-control" value={this.state.refundAmount} onChange={this.changerefundAmountHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Requested
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          InTransit
                      </option>
                      <option name="Status" className="form-control" >
                          Received
                      </option>
                      <option name="Status" className="form-control" >
                          Refunded
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateReturnRequest}>Save</button>
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

export default UpdateReturnRequestComponent
