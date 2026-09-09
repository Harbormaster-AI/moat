import React, { Component } from 'react'
import InspectionLotService from '../services/InspectionLotService'

class ViewInspectionLotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inspectionLot: {}
        }
    }

    componentDidMount(){
        InspectionLotService.getInspectionLotById(this.state.id).then( res => {
            this.setState({inspectionLot: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InspectionLot Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lotNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionLot.lotNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionLot.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> sampleSize:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionLot.sampleSize }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdOn:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionLot.createdOn }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> InspectionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionLot.inspectionType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionLot.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInspectionLotComponent
