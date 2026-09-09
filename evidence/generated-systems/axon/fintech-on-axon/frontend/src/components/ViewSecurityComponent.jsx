import React, { Component } from 'react'
import SecurityService from '../services/SecurityService'

class ViewSecurityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            security: {}
        }
    }

    componentDidMount(){
        SecurityService.getSecurityById(this.state.id).then( res => {
            this.setState({security: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Security Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> symbol:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.security.symbol }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> isin:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.security.isin }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> cusip:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.security.cusip }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.security.currency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SecurityType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.security.securityType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSecurityComponent
