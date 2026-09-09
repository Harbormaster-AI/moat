import React, { Component } from 'react'
import WarehouseService from '../services/WarehouseService'

class ViewWarehouseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            warehouse: {}
        }
    }

    componentDidMount(){
        WarehouseService.getWarehouseById(this.state.id).then( res => {
            this.setState({warehouse: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Warehouse Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.warehouse.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> warehouseCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.warehouse.warehouseCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.warehouse.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> WarehouseType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.warehouse.warehouseType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWarehouseComponent
