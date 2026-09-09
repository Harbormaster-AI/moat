import React, { Component } from 'react'
import DependentService from '../services/DependentService';

class CreateDependentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                birthDate: '',
                relationship: ''
        }
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changebirthDateHandler = this.changebirthDateHandler.bind(this);
        this.changeRelationshipHandler = this.changeRelationshipHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DependentService.getDependentById(this.state.id).then( (res) =>{
                let dependent = res.data;
                this.setState({
                    firstName: dependent.firstName,
                    lastName: dependent.lastName,
                    birthDate: dependent.birthDate,
                    relationship: dependent.relationship
                });
            });
        }        
    }
    saveOrUpdateDependent = (e) => {
        e.preventDefault();
        let dependent = {
                dependentId: this.state.id,
                firstName: this.state.firstName,
                lastName: this.state.lastName,
                birthDate: this.state.birthDate,
                relationship: this.state.relationship
            };
        console.log('dependent => ' + JSON.stringify(dependent));

        // step 5
        if(this.state.id === '_add'){
            dependent.dependentId=''
            DependentService.createDependent(dependent).then(res =>{
                this.props.history.push('/dependents');
            });
        }else{
            DependentService.updateDependent(dependent).then( res => {
                this.props.history.push('/dependents');
            });
        }
    }
    
    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changebirthDateHandler= (event) => {
        this.setState({birthDate: event.target.value});
    }
    changeRelationshipHandler= (event) => {
        this.setState({relationship: event.target.value});
    }

    cancel(){
        this.props.history.push('/dependents');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Dependent</h3>
        }else{
            return <h3 className="text-center">Update Dependent</h3>
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

                                            <label> birthDate:&emsp; </label>
                                                <input type="date" placeholder="birthDate" name="birthDate" className="form-control" value={this.state.birthDate} onChange={this.changebirthDateHandler}/>

                                            <label> Relationship:&emsp; </label>
                                                <select value={this.state.relationship} onChange={this.changeRelationshipHandler}>
                      <option name="Relationship" className="form-control" >
                          Spouse
                      </option>
                      <option name="Relationship" className="form-control" >
                          DomesticPartner
                      </option>
                      <option name="Relationship" className="form-control" >
                          Child
                      </option>
                      <option name="Relationship" className="form-control" >
                          Other
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDependent}>Save</button>
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

export default CreateDependentComponent
