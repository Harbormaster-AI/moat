import React, { Component } from 'react'
import TerminationService from '../services/TerminationService'

class ViewTerminationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            termination: {}
        }
    }

    componentDidMount(){
        TerminationService.getTerminationById(this.state.id).then( res => {
            this.setState({termination: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Termination Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> terminationNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.termination.terminationNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> terminationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.termination.terminationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> notes:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.termination.notes }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> eligibleForRehire:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.termination.eligibleForRehire }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Reason:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.termination.reason }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Type:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.termination.type }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTerminationComponent
