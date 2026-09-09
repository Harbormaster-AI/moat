import React, { Component } from 'react'
import SalaryComponentService from '../services/SalaryComponentService';

class UpdateSalaryComponentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                amount: '',
                recurring: '',
                componentType: ''
        }
        this.updateSalaryComponent = this.updateSalaryComponent.bind(this);

        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changerecurringHandler = this.changerecurringHandler.bind(this);
        this.changeComponentTypeHandler = this.changeComponentTypeHandler.bind(this);
    }

    componentDidMount(){
        SalaryComponentService.getSalaryComponentById(this.state.id).then( (res) =>{
            let salaryComponent = res.data;
            this.setState({
                amount: salaryComponent.amount,
                recurring: salaryComponent.recurring,
                componentType: salaryComponent.componentType
            });
        });
    }

    updateSalaryComponent = (e) => {
        e.preventDefault();
        let salaryComponent = {
            salaryComponentId: this.state.id,
            amount: this.state.amount,
            recurring: this.state.recurring,
            componentType: this.state.componentType
        };
        console.log('salaryComponent => ' + JSON.stringify(salaryComponent));
        console.log('id => ' + JSON.stringify(this.state.id));
        SalaryComponentService.updateSalaryComponent(salaryComponent).then( res => {
            this.props.history.push('/salaryComponents');
        });
    }

    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changerecurringHandler= (event) => {
        this.setState({recurring: event.target.value});
    }
    changeComponentTypeHandler= (event) => {
        this.setState({componentType: event.target.value});
    }

    cancel(){
        this.props.history.push('/salaryComponents');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SalaryComponent</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> recurring: </label>
                                                <input type="checkbox" placeholder="recurring" name="recurring" className="form-control" value={this.state.recurring} onChange={this.changerecurringHandler}/>


                                            <label> ComponentType: </label>
                                                <select value={this.state.componentType} onChange={this.changeComponentTypeHandler}>
                      <option name="ComponentType" className="form-control" >
                          BaseSalary
                      </option>
                      <option name="ComponentType" className="form-control" >
                          Allowance
                      </option>
                      <option name="ComponentType" className="form-control" >
                          OvertimeRate
                      </option>
                      <option name="ComponentType" className="form-control" >
                          Commission
                      </option>
                      <option name="ComponentType" className="form-control" >
                          ShiftDifferential
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSalaryComponent}>Save</button>
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

export default UpdateSalaryComponentComponent
