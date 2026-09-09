import React, { Component } from 'react'
import ClinicalOrderService from '../services/ClinicalOrderService';

class UpdateClinicalOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                orderNumber: '',
                status: '',
                orderType: '',
                priority: ''
        }
        this.updateClinicalOrder = this.updateClinicalOrder.bind(this);

        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeOrderTypeHandler = this.changeOrderTypeHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    componentDidMount(){
        ClinicalOrderService.getClinicalOrderById(this.state.id).then( (res) =>{
            let clinicalOrder = res.data;
            this.setState({
                orderNumber: clinicalOrder.orderNumber,
                status: clinicalOrder.status,
                orderType: clinicalOrder.orderType,
                priority: clinicalOrder.priority
            });
        });
    }

    updateClinicalOrder = (e) => {
        e.preventDefault();
        let clinicalOrder = {
            clinicalOrderId: this.state.id,
            orderNumber: this.state.orderNumber,
            status: this.state.status,
            orderType: this.state.orderType,
            priority: this.state.priority
        };
        console.log('clinicalOrder => ' + JSON.stringify(clinicalOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        ClinicalOrderService.updateClinicalOrder(clinicalOrder).then( res => {
            this.props.history.push('/clinicalOrders');
        });
    }

    changeorderNumberHandler= (event) => {
        this.setState({orderNumber: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeOrderTypeHandler= (event) => {
        this.setState({orderType: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }

    cancel(){
        this.props.history.push('/clinicalOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ClinicalOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> orderNumber: </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> OrderType: </label>
                                                <select value={this.state.orderType} onChange={this.changeOrderTypeHandler}>
                      <option name="OrderType" className="form-control" >
                          Medication
                      </option>
                      <option name="OrderType" className="form-control" >
                          Laboratory
                      </option>
                      <option name="OrderType" className="form-control" >
                          Imaging
                      </option>
                      <option name="OrderType" className="form-control" >
                          Procedure
                      </option>
                      <option name="OrderType" className="form-control" >
                          Consultation
                      </option>
                    </select>

                                            <label> Priority: </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Routine
                      </option>
                      <option name="Priority" className="form-control" >
                          Urgent
                      </option>
                      <option name="Priority" className="form-control" >
                          Stat
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateClinicalOrder}>Save</button>
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

export default UpdateClinicalOrderComponent
