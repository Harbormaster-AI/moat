import React, { Component } from 'react'
import OperationService from '../services/OperationService';

class UpdateOperationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                operationNumber: '',
                name: '',
                setupTime: '',
                standardCycleTime: '',
                operationType: ''
        }
        this.updateOperation = this.updateOperation.bind(this);

        this.changeoperationNumberHandler = this.changeoperationNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesetupTimeHandler = this.changesetupTimeHandler.bind(this);
        this.changestandardCycleTimeHandler = this.changestandardCycleTimeHandler.bind(this);
        this.changeOperationTypeHandler = this.changeOperationTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateOperation = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        OperationService.updateOperation(operation).then( res => {
            this.props.history.push('/operations');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Operation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> operationNumber: </label>
                                                <input placeholder="operationNumber" name="operationNumber" className="form-control" value={this.state.operationNumber} onChange={this.changeoperationNumberHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> setupTime: </label>
                                                <input placeholder="setupTime" name="setupTime" className="form-control" value={this.state.setupTime} onChange={this.changesetupTimeHandler}/>

                                            <label> standardCycleTime: </label>
                                                <input placeholder="standardCycleTime" name="standardCycleTime" className="form-control" value={this.state.standardCycleTime} onChange={this.changestandardCycleTimeHandler}/>

                                            <label> OperationType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateOperation}>Save</button>
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

export default UpdateOperationComponent
