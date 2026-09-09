import React, { Component } from 'react'
import ReturnRequestService from '../services/ReturnRequestService'

class ViewReturnRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            returnRequest: {}
        }
    }

    componentDidMount(){
        ReturnRequestService.getReturnRequestById(this.state.id).then( res => {
            this.setState({returnRequest: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ReturnRequest Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> returnNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.returnRequest.returnNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.returnRequest.createdAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> refundAmount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.returnRequest.refundAmount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.returnRequest.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewReturnRequestComponent
