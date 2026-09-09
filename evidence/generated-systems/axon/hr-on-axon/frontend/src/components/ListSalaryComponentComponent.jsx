import React, { Component } from 'react'
import SalaryComponentService from '../services/SalaryComponentService'

class ListSalaryComponentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                salaryComponents: []
        }
        this.addSalaryComponent = this.addSalaryComponent.bind(this);
        this.editSalaryComponent = this.editSalaryComponent.bind(this);
        this.deleteSalaryComponent = this.deleteSalaryComponent.bind(this);
    }

    deleteSalaryComponent(id){
        SalaryComponentService.deleteSalaryComponent(id).then( res => {
            this.setState({salaryComponents: this.state.salaryComponents.filter(salaryComponent => salaryComponent.salaryComponentId !== id)});
        });
    }
    viewSalaryComponent(id){
        this.props.history.push(`/view-salaryComponent/${id}`);
    }
    editSalaryComponent(id){
        this.props.history.push(`/add-salaryComponent/${id}`);
    }

    componentDidMount(){
        SalaryComponentService.getSalaryComponents().then((res) => {
            this.setState({ salaryComponents: res.data});
        });
    }

    addSalaryComponent(){
        this.props.history.push('/add-salaryComponent/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SalaryComponent List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSalaryComponent}> Add SalaryComponent</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Amount </th>
                                    <th> Recurring </th>
                                    <th> ComponentType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.salaryComponents.map(
                                        salaryComponent => 
                                        <tr key = {salaryComponent.salaryComponentId}>
                                             <td> { salaryComponent.amount } </td>
                                             <td> { salaryComponent.recurring } </td>
                                             <td> { salaryComponent.componentType } </td>
                                             <td>
                                                 <button onClick={ () => this.editSalaryComponent(salaryComponent.salaryComponentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSalaryComponent(salaryComponent.salaryComponentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSalaryComponent(salaryComponent.salaryComponentId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListSalaryComponentComponent
