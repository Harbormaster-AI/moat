import React, { Component } from 'react'
import WarrantyService from '../services/WarrantyService'

class ViewWarrantyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            warranty: {}
        }
    }

    componentDidMount(){
        WarrantyService.getWarrantyById(this.state.id).then( res => {
            this.setState({warranty: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Warranty Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> coverageMonths:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.warranty.coverageMonths }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> WarrantyType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.warranty.warrantyType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWarrantyComponent
