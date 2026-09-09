import React, { Component } from 'react'
import PersonService from '../services/PersonService';

class CreatePersonComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                email: '',
                department: ''
        }
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changedepartmentHandler = this.changedepartmentHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePerson = (e) => {
        e.preventDefault();
        let person = {
                personId: this.state.id,
                firstName: this.state.firstName,
                lastName: this.state.lastName,
                email: this.state.email,
                department: this.state.department
            };
        console.log('person => ' + JSON.stringify(person));

        // step 5
        if(this.state.id === '_add'){
            person.personId=''
            PersonService.createPerson(person).then(res =>{
                this.props.history.push('/persons');
            });
        }else{
            PersonService.updatePerson(person).then( res => {
                this.props.history.push('/persons');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Person</h3>
        }else{
            return <h3 className="text-center">Update Person</h3>
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
                                            <label> firstName:&emsp; </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName:&emsp; </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> email:&emsp; </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> department:&emsp; </label>
                                                <input placeholder="department" name="department" className="form-control" value={this.state.department} onChange={this.changedepartmentHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePerson}>Save</button>
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

export default CreatePersonComponent
