import React, { Component } from 'react'
import ProcedureService from '../services/ProcedureService'

class ViewProcedureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            procedure: {}
        }
    }

    componentDidMount(){
        ProcedureService.getProcedureById(this.state.id).then( res => {
            this.setState({procedure: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Procedure Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> procedureCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedure.procedureCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedure.startDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedure.endDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.procedure.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProcedureComponent
