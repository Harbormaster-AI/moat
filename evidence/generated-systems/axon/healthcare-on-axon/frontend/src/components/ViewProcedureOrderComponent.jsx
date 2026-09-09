import React, { Component } from 'react'
import ProcedureOrderService from '../services/ProcedureOrderService'

class ViewProcedureOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            procedureOrder: {}
        }
    }

    componentDidMount(){
        ProcedureOrderService.getProcedureOrderById(this.state.id).then( res => {
            this.setState({procedureOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ProcedureOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> procedureCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedureOrder.procedureCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> consentObtained:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedureOrder.consentObtained }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AnesthesiaType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedureOrder.anesthesiaType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProcedureOrderComponent
