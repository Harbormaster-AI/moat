import React, { Component } from 'react'
import APIClientService from '../services/APIClientService'

class ViewAPIClientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aPIClient: {}
        }
    }

    componentDidMount(){
        APIClientService.getAPIClientById(this.state.id).then( res => {
            this.setState({aPIClient: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View APIClient Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aPIClient.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> clientId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aPIClient.clientId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> redirectUri:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aPIClient.redirectUri }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ClientType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aPIClient.clientType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAPIClientComponent
