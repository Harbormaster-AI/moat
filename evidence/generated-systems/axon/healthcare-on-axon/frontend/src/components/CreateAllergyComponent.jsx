import React, { Component } from 'react'
import AllergyService from '../services/AllergyService';

class CreateAllergyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                substance: '',
                reaction: '',
                severity: '',
                status: ''
        }
        this.changesubstanceHandler = this.changesubstanceHandler.bind(this);
        this.changereactionHandler = this.changereactionHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AllergyService.getAllergyById(this.state.id).then( (res) =>{
                let allergy = res.data;
                this.setState({
                    substance: allergy.substance,
                    reaction: allergy.reaction,
                    severity: allergy.severity,
                    status: allergy.status
                });
            });
        }        
    }
    saveOrUpdateAllergy = (e) => {
        e.preventDefault();
        let allergy = {
                allergyId: this.state.id,
                substance: this.state.substance,
                reaction: this.state.reaction,
                severity: this.state.severity,
                status: this.state.status
            };
        console.log('allergy => ' + JSON.stringify(allergy));

        // step 5
        if(this.state.id === '_add'){
            allergy.allergyId=''
            AllergyService.createAllergy(allergy).then(res =>{
                this.props.history.push('/allergys');
            });
        }else{
            AllergyService.updateAllergy(allergy).then( res => {
                this.props.history.push('/allergys');
            });
        }
    }
    
    changesubstanceHandler= (event) => {
        this.setState({substance: event.target.value});
    }
    changereactionHandler= (event) => {
        this.setState({reaction: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/allergys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Allergy</h3>
        }else{
            return <h3 className="text-center">Update Allergy</h3>
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
                                            <label> substance:&emsp; </label>
                                                <input placeholder="substance" name="substance" className="form-control" value={this.state.substance} onChange={this.changesubstanceHandler}/>

                                            <label> reaction:&emsp; </label>
                                                <input placeholder="reaction" name="reaction" className="form-control" value={this.state.reaction} onChange={this.changereactionHandler}/>

                                            <label> Severity:&emsp; </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Mild
                      </option>
                      <option name="Severity" className="form-control" >
                          Moderate
                      </option>
                      <option name="Severity" className="form-control" >
                          Severe
                      </option>
                      <option name="Severity" className="form-control" >
                          LifeThreatening
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Inactive
                      </option>
                      <option name="Status" className="form-control" >
                          Resolved
                      </option>
                      <option name="Status" className="form-control" >
                          EnteredInError
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAllergy}>Save</button>
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

export default CreateAllergyComponent
