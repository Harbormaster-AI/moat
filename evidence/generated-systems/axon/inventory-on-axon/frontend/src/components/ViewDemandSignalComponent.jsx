import React, { Component } from 'react'
import DemandSignalService from '../services/DemandSignalService'

class ViewDemandSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            demandSignal: {}
        }
    }

    componentDidMount(){
        DemandSignalService.getDemandSignalById(this.state.id).then( res => {
            this.setState({demandSignal: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DemandSignal Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> externalReference:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.demandSignal.externalReference }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requestedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.demandSignal.requestedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.demandSignal.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DemandType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.demandSignal.demandType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDemandSignalComponent
