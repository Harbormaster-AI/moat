import React, { Component } from 'react'
import ServiceProviderService from '../services/ServiceProviderService'

class ViewServiceProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            serviceProvider: {}
        }
    }

    componentDidMount(){
        ServiceProviderService.getServiceProviderById(this.state.id).then( res => {
            this.setState({serviceProvider: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ServiceProvider Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serviceProvider.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taxId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serviceProvider.taxId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ProviderType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serviceProvider.providerType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> NetworkStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serviceProvider.networkStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewServiceProviderComponent
