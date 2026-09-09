import React, { Component } from 'react'
import PersonService from '../services/PersonService'

class ListPersonComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                persons: []
        }
        this.addPerson = this.addPerson.bind(this);
        this.editPerson = this.editPerson.bind(this);
        this.deletePerson = this.deletePerson.bind(this);
    }

    deletePerson(id){
        PersonService.deletePerson(id).then( res => {
            this.setState({persons: this.state.persons.filter(person => person.personId !== id)});
        });
    }
    viewPerson(id){
        this.props.history.push(`/view-person/${id}`);
    }
    editPerson(id){
        this.props.history.push(`/add-person/${id}`);
    }

    componentDidMount(){
        PersonService.getPersons().then((res) => {
            this.setState({ persons: res.data});
        });
    }

    addPerson(){
        this.props.history.push('/add-person/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Person List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPerson}> Add Person</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> Email </th>
                                    <th> Department </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.persons.map(
                                        person => 
                                        <tr key = {person.personId}>
                                             <td> { person.firstName } </td>
                                             <td> { person.lastName } </td>
                                             <td> { person.email } </td>
                                             <td> { person.department } </td>
                                             <td>
                                                 <button onClick={ () => this.editPerson(person.personId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePerson(person.personId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPerson(person.personId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPersonComponent
