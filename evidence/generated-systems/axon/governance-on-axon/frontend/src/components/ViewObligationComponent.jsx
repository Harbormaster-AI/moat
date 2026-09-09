import React, { Component } from 'react'
import ObligationService from '../services/ObligationService'

class ViewObligationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            obligation: {}
        }
    }

    componentDidMount(){
        ObligationService.getObligationById(this.state.id).then( res => {
            this.setState({obligation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Obligation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> referenceNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.obligation.referenceNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> descriptionText:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.obligation.descriptionText }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ObligationType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.obligation.obligationType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ReviewFrequency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.obligation.reviewFrequency }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewObligationComponent
