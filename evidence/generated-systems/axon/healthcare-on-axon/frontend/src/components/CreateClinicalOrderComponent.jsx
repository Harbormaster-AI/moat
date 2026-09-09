import React, { Component } from 'react'
import ClinicalOrderService from '../services/ClinicalOrderService';

class CreateClinicalOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                orderNumber: '',
                status: '',
                orderType: '',
                priority: ''
        }
        this.changeorderNumberHandler = this.changeorderNumberHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeOrderTypeHandler = this.changeOrderTypeHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateClinicalOrder = (e) => {
        e.preventDefault();
        let clinicalOrder = {
                clinicalOrderId: this.state.id,
                orderNumber: this.state.orderNumber,
                status: this.state.status,
                orderType: this.state.orderType,
                priority: this.state.priority
            };
        console.log('clinicalOrder => ' + JSON.stringify(clinicalOrder));

        // step 5
        if(this.state.id === '_add'){
            clinicalOrder.clinicalOrderId=''
            ClinicalOrderService.createClinicalOrder(clinicalOrder).then(res =>{
                this.props.history.push('/clinicalOrders');
            });
        }else{
            ClinicalOrderService.updateClinicalOrder(clinicalOrder).then( res => {
                this.props.history.push('/clinicalOrders');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ClinicalOrder</h3>
        }else{
            return <h3 className="text-center">Update ClinicalOrder</h3>
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
                                            <label> orderNumber:&emsp; </label>
                                                <input placeholder="orderNumber" name="orderNumber" className="form-control" value={this.state.orderNumber} onChange={this.changeorderNumberHandler}/>

                                            <label> Status:&emsp; </label>
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

                                            <label> OrderType:&emsp; </label>
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

                                            <label> Priority:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateClinicalOrder}>Save</button>
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

export default CreateClinicalOrderComponent
