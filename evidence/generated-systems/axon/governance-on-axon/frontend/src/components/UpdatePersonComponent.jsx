import React, { Component } from 'react'
import PersonService from '../services/PersonService';

class UpdatePersonComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                email: '',
                department: ''
        }
        this.updatePerson = this.updatePerson.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changedepartmentHandler = this.changedepartmentHandler.bind(this);
    }

    componentDidMount(){
        PersonService.getPersonById(this.state.id).then( (res) =>{
            let person = res.data;
            this.setState({
                firstName: person.firstName,
                lastName: person.lastName,
                email: person.email,
                department: person.department
            });
        });
    }

    updatePerson = (e) => {
        e.preventDefault();
        let person = {
            personId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            email: this.state.email,
            department: this.state.department
        };
        console.log('person => ' + JSON.stringify(person));
        console.log('id => ' + JSON.stringify(this.state.id));
        PersonService.updatePerson(person).then( res => {
            this.props.history.push('/persons');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changedepartmentHandler= (event) => {
        this.setState({department: event.target.value});
    }

    cancel(){
        this.props.history.push('/persons');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Person</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> department: </label>
                                                <input placeholder="department" name="department" className="form-control" value={this.state.department} onChange={this.changedepartmentHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePerson}>Save</button>
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

export default UpdatePersonComponent
