import React, { Component } from 'react'
import OperationService from '../services/OperationService';

class CreateOperationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                operationNumber: '',
                name: '',
                setupTime: '',
                standardCycleTime: '',
                operationType: ''
        }
        this.changeoperationNumberHandler = this.changeoperationNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesetupTimeHandler = this.changesetupTimeHandler.bind(this);
        this.changestandardCycleTimeHandler = this.changestandardCycleTimeHandler.bind(this);
        this.changeOperationTypeHandler = this.changeOperationTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OperationService.getOperationById(this.state.id).then( (res) =>{
                let operation = res.data;
                this.setState({
                    operationNumber: operation.operationNumber,
                    name: operation.name,
                    setupTime: operation.setupTime,
                    standardCycleTime: operation.standardCycleTime,
                    operationType: operation.operationType
                });
            });
        }        
    }
    saveOrUpdateOperation = (e) => {
        e.preventDefault();
        let operation = {
                operationId: this.state.id,
                operationNumber: this.state.operationNumber,
                name: this.state.name,
                setupTime: this.state.setupTime,
                standardCycleTime: this.state.standardCycleTime,
                operationType: this.state.operationType
            };
        console.log('operation => ' + JSON.stringify(operation));

        // step 5
        if(this.state.id === '_add'){
            operation.operationId=''
            OperationService.createOperation(operation).then(res =>{
                this.props.history.push('/operations');
            });
        }else{
            OperationService.updateOperation(operation).then( res => {
                this.props.history.push('/operations');
            });
        }
    }
    
    changeoperationNumberHandler= (event) => {
        this.setState({operationNumber: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changesetupTimeHandler= (event) => {
        this.setState({setupTime: event.target.value});
    }
    changestandardCycleTimeHandler= (event) => {
        this.setState({standardCycleTime: event.target.value});
    }
    changeOperationTypeHandler= (event) => {
        this.setState({operationType: event.target.value});
    }

    cancel(){
        this.props.history.push('/operations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Operation</h3>
        }else{
            return <h3 className="text-center">Update Operation</h3>
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
                                            <label> operationNumber:&emsp; </label>
                                                <input placeholder="operationNumber" name="operationNumber" className="form-control" value={this.state.operationNumber} onChange={this.changeoperationNumberHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> setupTime:&emsp; </label>
                                                <input placeholder="setupTime" name="setupTime" className="form-control" value={this.state.setupTime} onChange={this.changesetupTimeHandler}/>

                                            <label> standardCycleTime:&emsp; </label>
                                                <input placeholder="standardCycleTime" name="standardCycleTime" className="form-control" value={this.state.standardCycleTime} onChange={this.changestandardCycleTimeHandler}/>

                                            <label> OperationType:&emsp; </label>
                                                <select value={this.state.operationType} onChange={this.changeOperationTypeHandler}>
                      <option name="OperationType" className="form-control" >
                          Setup
                      </option>
                      <option name="OperationType" className="form-control" >
                          Run
                      </option>
                      <option name="OperationType" className="form-control" >
                          Teardown
                      </option>
                      <option name="OperationType" className="form-control" >
                          Inspection
                      </option>
                      <option name="OperationType" className="form-control" >
                          Transfer
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOperation}>Save</button>
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

export default CreateOperationComponent
