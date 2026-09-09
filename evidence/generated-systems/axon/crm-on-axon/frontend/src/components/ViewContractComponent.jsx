import React, { Component } from 'react'
import ContractService from '../services/ContractService'

class ViewContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            contract: {}
        }
    }

    componentDidMount(){
        ContractService.getContractById(this.state.id).then( res => {
            this.setState({contract: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Contract Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> contractNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contract.contractNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contract.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contract.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> renewalTermMonths:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contract.renewalTermMonths }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> autoRenew:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contract.autoRenew }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contract.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewContractComponent
