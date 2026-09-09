import React, { Component } from 'react'
import LabResultService from '../services/LabResultService'

class ViewLabResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            labResult: {}
        }
    }

    componentDidMount(){
        LabResultService.getLabResultById(this.state.id).then( res => {
            this.setState({labResult: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LabResult Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> resultCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.labResult.resultCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> issuedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.labResult.issuedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.labResult.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLabResultComponent
