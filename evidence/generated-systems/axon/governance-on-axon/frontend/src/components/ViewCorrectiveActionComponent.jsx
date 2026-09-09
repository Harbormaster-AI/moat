import React, { Component } from 'react'
import CorrectiveActionService from '../services/CorrectiveActionService'

class ViewCorrectiveActionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            correctiveAction: {}
        }
    }

    componentDidMount(){
        CorrectiveActionService.getCorrectiveActionById(this.state.id).then( res => {
            this.setState({correctiveAction: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CorrectiveAction Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> actionTitle:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.correctiveAction.actionTitle }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> owner:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.correctiveAction.owner }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> targetDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.correctiveAction.targetDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.correctiveAction.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCorrectiveActionComponent
