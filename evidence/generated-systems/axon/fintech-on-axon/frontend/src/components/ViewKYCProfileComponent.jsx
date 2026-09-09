import React, { Component } from 'react'
import KYCProfileService from '../services/KYCProfileService'

class ViewKYCProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            kYCProfile: {}
        }
    }

    componentDidMount(){
        KYCProfileService.getKYCProfileById(this.state.id).then( res => {
            this.setState({kYCProfile: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View KYCProfile Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> profileId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCProfile.profileId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCProfile.createdAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCProfile.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> VerificationLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.kYCProfile.verificationLevel }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewKYCProfileComponent
