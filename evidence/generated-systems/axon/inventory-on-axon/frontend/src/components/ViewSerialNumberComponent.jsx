import React, { Component } from 'react'
import SerialNumberService from '../services/SerialNumberService'

class ViewSerialNumberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            serialNumber: {}
        }
    }

    componentDidMount(){
        SerialNumberService.getSerialNumberById(this.state.id).then( res => {
            this.setState({serialNumber: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SerialNumber Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> serial:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serialNumber.serial }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> activationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serialNumber.activationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serialNumber.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSerialNumberComponent
