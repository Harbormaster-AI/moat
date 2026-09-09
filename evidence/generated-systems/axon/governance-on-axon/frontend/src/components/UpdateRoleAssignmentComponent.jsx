import React, { Component } from 'react'
import RoleAssignmentService from '../services/RoleAssignmentService';

class UpdateRoleAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                effectiveFrom: '',
                effectiveTo: ''
        }
        this.updateRoleAssignment = this.updateRoleAssignment.bind(this);

        this.changeeffectiveFromHandler = this.changeeffectiveFromHandler.bind(this);
        this.changeeffectiveToHandler = this.changeeffectiveToHandler.bind(this);
    }

    componentDidMount(){
        RoleAssignmentService.getRoleAssignmentById(this.state.id).then( (res) =>{
            let roleAssignment = res.data;
            this.setState({
                effectiveFrom: roleAssignment.effectiveFrom,
                effectiveTo: roleAssignment.effectiveTo
            });
        });
    }

    updateRoleAssignment = (e) => {
        e.preventDefault();
        let roleAssignment = {
            roleAssignmentId: this.state.id,
            effectiveFrom: this.state.effectiveFrom,
            effectiveTo: this.state.effectiveTo
        };
        console.log('roleAssignment => ' + JSON.stringify(roleAssignment));
        console.log('id => ' + JSON.stringify(this.state.id));
        RoleAssignmentService.updateRoleAssignment(roleAssignment).then( res => {
            this.props.history.push('/roleAssignments');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update RoleAssignment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> effectiveFrom: </label>
                                                <input type="date" placeholder="effectiveFrom" name="effectiveFrom" className="form-control" value={this.state.effectiveFrom} onChange={this.changeeffectiveFromHandler}/>

                                            <label> effectiveTo: </label>
                                                <input type="date" placeholder="effectiveTo" name="effectiveTo" className="form-control" value={this.state.effectiveTo} onChange={this.changeeffectiveToHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRoleAssignment}>Save</button>
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

export default UpdateRoleAssignmentComponent
