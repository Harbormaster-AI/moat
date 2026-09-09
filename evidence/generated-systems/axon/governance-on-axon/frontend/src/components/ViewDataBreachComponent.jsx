import React, { Component } from 'react'
import DataBreachService from '../services/DataBreachService'

class ViewDataBreachComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataBreach: {}
        }
    }

    componentDidMount(){
        DataBreachService.getDataBreachById(this.state.id).then( res => {
            this.setState({dataBreach: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataBreach Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> incidentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataBreach.incidentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataBreach.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> recordsAffected:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataBreach.recordsAffected }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> notificationRequired:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataBreach.notificationRequired }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataBreach.severity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataBreach.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataBreachComponent
