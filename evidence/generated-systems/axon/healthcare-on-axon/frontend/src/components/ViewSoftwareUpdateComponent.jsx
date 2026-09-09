import React, { Component } from 'react'
import SoftwareUpdateService from '../services/SoftwareUpdateService'

class ViewSoftwareUpdateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            softwareUpdate: {}
        }
    }

    componentDidMount(){
        SoftwareUpdateService.getSoftwareUpdateById(this.state.id).then( res => {
            this.setState({softwareUpdate: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SoftwareUpdate Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> version:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.softwareUpdate.version }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> appliedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.softwareUpdate.appliedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> UpdateType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.softwareUpdate.updateType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSoftwareUpdateComponent
