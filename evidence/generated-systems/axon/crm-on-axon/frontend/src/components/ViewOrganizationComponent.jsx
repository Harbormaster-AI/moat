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
                            <div className = "col" style={{textAlign:"right"}}><label> defaultCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.defaultCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> defaultLocale:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.defaultLocale }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.organization.website }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOrganizationComponent
