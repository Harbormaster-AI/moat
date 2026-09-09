import React, { Component } from 'react'
import OrganizationService from '../services/OrganizationService'

class ViewOrganizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            organization: {}
        }
    }

    componentDidMount(){
        OrganizationService.getOrganizationById(this.state.id).then( res => {
            this.setState({organization: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Organization Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> legalName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.legalName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> jurisdiction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.jurisdiction }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> industrySector:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.industrySector }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOrganizationComponent
