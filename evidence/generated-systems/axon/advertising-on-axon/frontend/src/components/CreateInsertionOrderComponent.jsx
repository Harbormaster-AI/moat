import React, { Component } from 'react'
import InsertionOrderService from '../services/InsertionOrderService';

class CreateInsertionOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                ioNumber: '',
                agreedBudget: '',
                flight: '',
                status: ''
        }
        this.changeioNumberHandler = this.changeioNumberHandler.bind(this);
        this.changeagreedBudgetHandler = this.changeagreedBudgetHandler.bind(this);
        this.changeflightHandler = this.changeflightHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InsertionOrderService.getInsertionOrderById(this.state.id).then( (res) =>{
                let insertionOrder = res.data;
                this.setState({
                    ioNumber: insertionOrder.ioNumber,
                    agreedBudget: insertionOrder.agreedBudget,
                    flight: insertionOrder.flight,
                    status: insertionOrder.status
                });
            });
        }        
    }
    saveOrUpdateInsertionOrder = (e) => {
        e.preventDefault();
        let insertionOrder = {
                insertionOrderId: this.state.id,
                ioNumber: this.state.ioNumber,
                agreedBudget: this.state.agreedBudget,
                flight: this.state.flight,
                status: this.state.status
            };
        console.log('insertionOrder => ' + JSON.stringify(insertionOrder));

        // step 5
        if(this.state.id === '_add'){
            insertionOrder.insertionOrderId=''
            InsertionOrderService.createInsertionOrder(insertionOrder).then(res =>{
                this.props.history.push('/insertionOrders');
            });
        }else{
            InsertionOrderService.updateInsertionOrder(insertionOrder).then( res => {
                this.props.history.push('/insertionOrders');
            });
        }
    }
    
    changeioNumberHandler= (event) => {
        this.setState({ioNumber: event.target.value});
    }
    changeagreedBudgetHandler= (event) => {
        this.setState({agreedBudget: event.target.value});
    }
    changeflightHandler= (event) => {
        this.setState({flight: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/insertionOrders');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InsertionOrder</h3>
        }else{
            return <h3 className="text-center">Update InsertionOrder</h3>
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
                                            <label> ioNumber:&emsp; </label>
                                                <input placeholder="ioNumber" name="ioNumber" className="form-control" value={this.state.ioNumber} onChange={this.changeioNumberHandler}/>

                                            <label> agreedBudget:&emsp; </label>
                                                <input placeholder="agreedBudget" name="agreedBudget" className="form-control" value={this.state.agreedBudget} onChange={this.changeagreedBudgetHandler}/>

                                            <label> flight:&emsp; </label>
                                                <input placeholder="flight" name="flight" className="form-control" value={this.state.flight} onChange={this.changeflightHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Sent
                      </option>
                      <option name="Status" className="form-control" >
                          Executed
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInsertionOrder}>Save</button>
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

export default CreateInsertionOrderComponent
