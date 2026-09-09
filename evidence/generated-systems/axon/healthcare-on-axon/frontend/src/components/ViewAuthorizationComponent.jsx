import React, { Component } from 'react'
import AuthorizationService from '../services/AuthorizationService'

class ViewAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            authorization: {}
        }
    }

    componentDidMount(){
        AuthorizationService.getAuthorizationById(this.state.id).then( res => {
            this.setState({authorization: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Authorization Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> authNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.authorization.authNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestedService:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.authorization.requestedService }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.authorization.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAuthorizationComponent
