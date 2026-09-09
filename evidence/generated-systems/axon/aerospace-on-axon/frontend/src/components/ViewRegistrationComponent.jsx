import React, { Component } from 'react'
import RegistrationService from '../services/RegistrationService'

class ViewRegistrationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            registration: {}
        }
    }

    componentDidMount(){
        RegistrationService.getRegistrationById(this.state.id).then( res => {
            this.setState({registration: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Registration Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> tailNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.registration.tailNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> registryCountry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.registration.registryCountry }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRegistrationComponent
