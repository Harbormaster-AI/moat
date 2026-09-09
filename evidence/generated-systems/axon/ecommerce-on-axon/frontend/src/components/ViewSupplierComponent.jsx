import React, { Component } from 'react'
import SupplierService from '../services/SupplierService'

class ViewSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            supplier: {}
        }
    }

    componentDidMount(){
        SupplierService.getSupplierById(this.state.id).then( res => {
            this.setState({supplier: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Supplier Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.supplier.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> contactEmail:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.supplier.contactEmail }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.supplier.website }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.supplier.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSupplierComponent
