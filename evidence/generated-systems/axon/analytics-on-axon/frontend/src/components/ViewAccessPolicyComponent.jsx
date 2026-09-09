import React, { Component } from 'react'
import AccessPolicyService from '../services/AccessPolicyService'

class ViewAccessPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            accessPolicy: {}
        }
    }

    componentDidMount(){
        AccessPolicyService.getAccessPolicyById(this.state.id).then( res => {
            this.setState({accessPolicy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AccessPolicy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.accessPolicy.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subjectName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.accessPolicy.subjectName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AccessLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.accessPolicy.accessLevel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SubjectType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.accessPolicy.subjectType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAccessPolicyComponent
