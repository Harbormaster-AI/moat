import React, { Component } from 'react'
import OperatorService from '../services/OperatorService';

class UpdateOperatorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                icaoDesignator: '',
                operatorType: ''
        }
        this.updateOperator = this.updateOperator.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeicaoDesignatorHandler = this.changeicaoDesignatorHandler.bind(this);
        this.changeOperatorTypeHandler = this.changeOperatorTypeHandler.bind(this);
    }

    componentDidMount(){
        OperatorService.getOperatorById(this.state.id).then( (res) =>{
            let operator = res.data;
            this.setState({
                name: operator.name,
                icaoDesignator: operator.icaoDesignator,
                operatorType: operator.operatorType
            });
        });
    }

    updateOperator = (e) => {
        e.preventDefault();
        let operator = {
            operatorId: this.state.id,
            name: this.state.name,
            icaoDesignator: this.state.icaoDesignator,
            operatorType: this.state.operatorType
        };
        console.log('operator => ' + JSON.stringify(operator));
        console.log('id => ' + JSON.stringify(this.state.id));
        OperatorService.updateOperator(operator).then( res => {
            this.props.history.push('/operators');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Operator</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> icaoDesignator: </label>
                                                <input placeholder="icaoDesignator" name="icaoDesignator" className="form-control" value={this.state.icaoDesignator} onChange={this.changeicaoDesignatorHandler}/>

                                            <label> OperatorType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateOperator}>Save</button>
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

export default UpdateOperatorComponent
