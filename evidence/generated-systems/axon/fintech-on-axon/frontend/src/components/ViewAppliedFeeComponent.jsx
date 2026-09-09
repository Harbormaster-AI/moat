import React, { Component } from 'react'
import AppliedFeeService from '../services/AppliedFeeService'

class ViewAppliedFeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            appliedFee: {}
        }
    }

    componentDidMount(){
        AppliedFeeService.getAppliedFeeById(this.state.id).then( res => {
            this.setState({appliedFee: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AppliedFee Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appliedFee.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appliedFee.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> FeeType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.appliedFee.feeType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAppliedFeeComponent
