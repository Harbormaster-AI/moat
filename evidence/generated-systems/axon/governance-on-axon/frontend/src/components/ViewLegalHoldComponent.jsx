import React, { Component } from 'react'
import LegalHoldService from '../services/LegalHoldService'

class ViewLegalHoldComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            legalHold: {}
        }
    }

    componentDidMount(){
        LegalHoldService.getLegalHoldById(this.state.id).then( res => {
            this.setState({legalHold: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LegalHold Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.legalHold.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reason:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.legalHold.reason }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> issuedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.legalHold.issuedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> releaseDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.legalHold.releaseDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> HoldStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.legalHold.holdStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLegalHoldComponent
