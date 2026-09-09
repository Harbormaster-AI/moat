import React, { Component } from 'react'
import OperatorService from '../services/OperatorService';

class CreateOperatorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                icaoDesignator: '',
                operatorType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeicaoDesignatorHandler = this.changeicaoDesignatorHandler.bind(this);
        this.changeOperatorTypeHandler = this.changeOperatorTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OperatorService.getOperatorById(this.state.id).then( (res) =>{
                let operator = res.data;
                this.setState({
                    name: operator.name,
                    icaoDesignator: operator.icaoDesignator,
                    operatorType: operator.operatorType
                });
            });
        }        
    }
    saveOrUpdateOperator = (e) => {
        e.preventDefault();
        let operator = {
                operatorId: this.state.id,
                name: this.state.name,
                icaoDesignator: this.state.icaoDesignator,
                operatorType: this.state.operatorType
            };
        console.log('operator => ' + JSON.stringify(operator));

        // step 5
        if(this.state.id === '_add'){
            operator.operatorId=''
            OperatorService.createOperator(operator).then(res =>{
                this.props.history.push('/operators');
            });
        }else{
            OperatorService.updateOperator(operator).then( res => {
                this.props.history.push('/operators');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeicaoDesignatorHandler= (event) => {
        this.setState({icaoDesignator: event.target.value});
    }
    changeOperatorTypeHandler= (event) => {
        this.setState({operatorType: event.target.value});
    }

    cancel(){
        this.props.history.push('/operators');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Operator</h3>
        }else{
            return <h3 className="text-center">Update Operator</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> icaoDesignator:&emsp; </label>
                                                <input placeholder="icaoDesignator" name="icaoDesignator" className="form-control" value={this.state.icaoDesignator} onChange={this.changeicaoDesignatorHandler}/>

                                            <label> OperatorType:&emsp; </label>
                                                <select value={this.state.operatorType} onChange={this.changeOperatorTypeHandler}>
                      <option name="OperatorType" className="form-control" >
                          Airline
                      </option>
                      <option name="OperatorType" className="form-control" >
                          Cargo
                      </option>
                      <option name="OperatorType" className="form-control" >
                          Government
                      </option>
                      <option name="OperatorType" className="form-control" >
                          Private
                      </option>
                      <option name="OperatorType" className="form-control" >
                          Lessor
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOperator}>Save</button>
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

export default CreateOperatorComponent
