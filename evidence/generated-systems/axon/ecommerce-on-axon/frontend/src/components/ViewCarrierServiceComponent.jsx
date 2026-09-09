import React, { Component } from 'react'
import CarrierServiceService from '../services/CarrierServiceService'

class ViewCarrierServiceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            carrierService: {}
        }
    }

    componentDidMount(){
        CarrierServiceService.getCarrierServiceById(this.state.id).then( res => {
            this.setState({carrierService: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CarrierService Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carrierService.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carrierService.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Carrier:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carrierService.carrier }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ServiceLevel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.carrierService.serviceLevel }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCarrierServiceComponent
