import React, { Component } from 'react'
import OutboundAllocationService from '../services/OutboundAllocationService';

class CreateOutboundAllocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                allocationNumber: '',
                allocatedQuantity: '',
                allocationDate: '',
                status: ''
        }
        this.changeallocationNumberHandler = this.changeallocationNumberHandler.bind(this);
        this.changeallocatedQuantityHandler = this.changeallocatedQuantityHandler.bind(this);
        this.changeallocationDateHandler = this.changeallocationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OutboundAllocationService.getOutboundAllocationById(this.state.id).then( (res) =>{
                let outboundAllocation = res.data;
                this.setState({
                    allocationNumber: outboundAllocation.allocationNumber,
                    allocatedQuantity: outboundAllocation.allocatedQuantity,
                    allocationDate: outboundAllocation.allocationDate,
                    status: outboundAllocation.status
                });
            });
        }        
    }
    saveOrUpdateOutboundAllocation = (e) => {
        e.preventDefault();
        let outboundAllocation = {
                outboundAllocationId: this.state.id,
                allocationNumber: this.state.allocationNumber,
                allocatedQuantity: this.state.allocatedQuantity,
                allocationDate: this.state.allocationDate,
                status: this.state.status
            };
        console.log('outboundAllocation => ' + JSON.stringify(outboundAllocation));

        // step 5
        if(this.state.id === '_add'){
            outboundAllocation.outboundAllocationId=''
            OutboundAllocationService.createOutboundAllocation(outboundAllocation).then(res =>{
                this.props.history.push('/outboundAllocations');
            });
        }else{
            OutboundAllocationService.updateOutboundAllocation(outboundAllocation).then( res => {
                this.props.history.push('/outboundAllocations');
            });
        }
    }
    
    changeallocationNumberHandler= (event) => {
        this.setState({allocationNumber: event.target.value});
    }
    changeallocatedQuantityHandler= (event) => {
        this.setState({allocatedQuantity: event.target.value});
    }
    changeallocationDateHandler= (event) => {
        this.setState({allocationDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/outboundAllocations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add OutboundAllocation</h3>
        }else{
            return <h3 className="text-center">Update OutboundAllocation</h3>
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
                                            <label> allocationNumber:&emsp; </label>
                                                <input placeholder="allocationNumber" name="allocationNumber" className="form-control" value={this.state.allocationNumber} onChange={this.changeallocationNumberHandler}/>

                                            <label> allocatedQuantity:&emsp; </label>
                                                <input placeholder="allocatedQuantity" name="allocatedQuantity" className="form-control" value={this.state.allocatedQuantity} onChange={this.changeallocatedQuantityHandler}/>

                                            <label> allocationDate:&emsp; </label>
                                                <input type="date" placeholder="allocationDate" name="allocationDate" className="form-control" value={this.state.allocationDate} onChange={this.changeallocationDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Proposed
                      </option>
                      <option name="Status" className="form-control" >
                          Confirmed
                      </option>
                      <option name="Status" className="form-control" >
                          Picked
                      </option>
                      <option name="Status" className="form-control" >
                          Short
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOutboundAllocation}>Save</button>
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

export default CreateOutboundAllocationComponent
