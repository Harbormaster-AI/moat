import React, { Component } from 'react'
import ClaimReserveService from '../services/ClaimReserveService'

class ViewClaimReserveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            claimReserve: {}
        }
    }

    componentDidMount(){
        ClaimReserveService.getClaimReserveById(this.state.id).then( res => {
            this.setState({claimReserve: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ClaimReserve Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimReserve.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> setDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimReserve.setDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ReserveType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimReserve.reserveType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.claimReserve.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewClaimReserveComponent
