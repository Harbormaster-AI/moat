import React, { Component } from 'react'
import InventorySourceService from '../services/InventorySourceService'

class ViewInventorySourceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inventorySource: {}
        }
    }

    componentDidMount(){
        InventorySourceService.getInventorySourceById(this.state.id).then( res => {
            this.setState({inventorySource: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InventorySource Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventorySource.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> domain:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventorySource.domain }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Channel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventorySource.channel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PrimaryFormat:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventorySource.primaryFormat }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInventorySourceComponent
