import React, { Component } from 'react'
import StorageLocationService from '../services/StorageLocationService'

class ViewStorageLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            storageLocation: {}
        }
    }

    componentDidMount(){
        StorageLocationService.getStorageLocationById(this.state.id).then( res => {
            this.setState({storageLocation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View StorageLocation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.storageLocation.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> temperatureControlled:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.storageLocation.temperatureControlled }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> capacity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.storageLocation.capacity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> capacityUnit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.storageLocation.capacityUnit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LocationType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.storageLocation.locationType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewStorageLocationComponent
