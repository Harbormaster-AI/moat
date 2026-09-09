import React, { Component } from 'react'
import RoleAssignmentService from '../services/RoleAssignmentService'

class ViewRoleAssignmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            roleAssignment: {}
        }
    }

    componentDidMount(){
        RoleAssignmentService.getRoleAssignmentById(this.state.id).then( res => {
            this.setState({roleAssignment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View RoleAssignment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveFrom:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.roleAssignment.effectiveFrom }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveTo:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.roleAssignment.effectiveTo }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRoleAssignmentComponent
