import React, { Component } from 'react'
import CreativeApprovalService from '../services/CreativeApprovalService'

class ViewCreativeApprovalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            creativeApproval: {}
        }
    }

    componentDidMount(){
        CreativeApprovalService.getCreativeApprovalById(this.state.id).then( res => {
            this.setState({creativeApproval: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CreativeApproval Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewer:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeApproval.reviewer }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeApproval.reviewedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeApproval.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCreativeApprovalComponent
