import React, { Component } from 'react'
import LaboratoryOrderService from '../services/LaboratoryOrderService'

class ViewLaboratoryOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            laboratoryOrder: {}
        }
    }

    componentDidMount(){
        LaboratoryOrderService.getLaboratoryOrderById(this.state.id).then( res => {
            this.setState({laboratoryOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LaboratoryOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> testCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.laboratoryOrder.testCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> fastingRequired:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.laboratoryOrder.fastingRequired }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SpecimenType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.laboratoryOrder.specimenType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLaboratoryOrderComponent
