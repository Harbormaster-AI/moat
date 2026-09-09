import React, { Component } from 'react'
import DependentService from '../services/DependentService'

class ListDependentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dependents: []
        }
        this.addDependent = this.addDependent.bind(this);
        this.editDependent = this.editDependent.bind(this);
        this.deleteDependent = this.deleteDependent.bind(this);
    }

    deleteDependent(id){
        DependentService.deleteDependent(id).then( res => {
            this.setState({dependents: this.state.dependents.filter(dependent => dependent.dependentId !== id)});
        });
    }
    viewDependent(id){
        this.props.history.push(`/view-dependent/${id}`);
    }
    editDependent(id){
        this.props.history.push(`/add-dependent/${id}`);
    }

    componentDidMount(){
        DependentService.getDependents().then((res) => {
            this.setState({ dependents: res.data});
        });
    }

    addDependent(){
        this.props.history.push('/add-dependent/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Dependent List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDependent}> Add Dependent</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> BirthDate </th>
                                    <th> Relationship </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dependents.map(
                                        dependent => 
                                        <tr key = {dependent.dependentId}>
                                             <td> { dependent.firstName } </td>
                                             <td> { dependent.lastName } </td>
                                             <td> { dependent.birthDate } </td>
                                             <td> { dependent.relationship } </td>
                                             <td>
                                                 <button onClick={ () => this.editDependent(dependent.dependentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDependent(dependent.dependentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDependent(dependent.dependentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDependentComponent
