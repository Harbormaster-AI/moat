import React, { Component } from 'react'
import SettlementBatchService from '../services/SettlementBatchService'

class ViewSettlementBatchComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            settlementBatch: {}
        }
    }

    componentDidMount(){
        SettlementBatchService.getSettlementBatchById(this.state.id).then( res => {
            this.setState({settlementBatch: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SettlementBatch Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> batchId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.settlementBatch.batchId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> periodStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.settlementBatch.periodStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> periodEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.settlementBatch.periodEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalVolume:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.settlementBatch.totalVolume }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> totalCount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.settlementBatch.totalCount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.settlementBatch.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSettlementBatchComponent
