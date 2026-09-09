import React, { Component } from 'react'
import QuarantineService from '../services/QuarantineService'

class ViewQuarantineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            quarantine: {}
        }
    }

    componentDidMount(){
        QuarantineService.getQuarantineById(this.state.id).then( res => {
            this.setState({quarantine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Quarantine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reason:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quarantine.reason }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quarantine.startedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> releasedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quarantine.releasedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Disposition:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.quarantine.disposition }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewQuarantineComponent
