import React, { Component } from 'react'
import RoleAssignmentService from '../services/RoleAssignmentService';

class CreateRoleAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                effectiveFrom: '',
                effectiveTo: ''
        }
        this.changeeffectiveFromHandler = this.changeeffectiveFromHandler.bind(this);
        this.changeeffectiveToHandler = this.changeeffectiveToHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RoleAssignmentService.getRoleAssignmentById(this.state.id).then( (res) =>{
                let roleAssignment = res.data;
                this.setState({
                    effectiveFrom: roleAssignment.effectiveFrom,
                    effectiveTo: roleAssignment.effectiveTo
                });
            });
        }        
    }
    saveOrUpdateRoleAssignment = (e) => {
        e.preventDefault();
        let roleAssignment = {
                roleAssignmentId: this.state.id,
                effectiveFrom: this.state.effectiveFrom,
                effectiveTo: this.state.effectiveTo
            };
        console.log('roleAssignment => ' + JSON.stringify(roleAssignment));

        // step 5
        if(this.state.id === '_add'){
            roleAssignment.roleAssignmentId=''
            RoleAssignmentService.createRoleAssignment(roleAssignment).then(res =>{
                this.props.history.push('/roleAssignments');
            });
        }else{
            RoleAssignmentService.updateRoleAssignment(roleAssignment).then( res => {
                this.props.history.push('/roleAssignments');
            });
        }
    }
    
    changeeffectiveFromHandler= (event) => {
        this.setState({effectiveFrom: event.target.value});
    }
    changeeffectiveToHandler= (event) => {
        this.setState({effectiveTo: event.target.value});
    }

    cancel(){
        this.props.history.push('/roleAssignments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add RoleAssignment</h3>
        }else{
            return <h3 className="text-center">Update RoleAssignment</h3>
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
                                            <label> effectiveFrom:&emsp; </label>
                                                <input type="date" placeholder="effectiveFrom" name="effectiveFrom" className="form-control" value={this.state.effectiveFrom} onChange={this.changeeffectiveFromHandler}/>

                                            <label> effectiveTo:&emsp; </label>
                                                <input type="date" placeholder="effectiveTo" name="effectiveTo" className="form-control" value={this.state.effectiveTo} onChange={this.changeeffectiveToHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRoleAssignment}>Save</button>
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

export default CreateRoleAssignmentComponent
