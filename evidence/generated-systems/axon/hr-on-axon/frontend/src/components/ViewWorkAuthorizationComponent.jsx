import React, { Component } from 'react'
import WorkAuthorizationService from '../services/WorkAuthorizationService'

class ViewWorkAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            workAuthorization: {}
        }
    }

    componentDidMount(){
        WorkAuthorizationService.getWorkAuthorizationById(this.state.id).then( res => {
            this.setState({workAuthorization: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View WorkAuthorization Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> country:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workAuthorization.country }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expirationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workAuthorization.expirationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workAuthorization.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWorkAuthorizationComponent
