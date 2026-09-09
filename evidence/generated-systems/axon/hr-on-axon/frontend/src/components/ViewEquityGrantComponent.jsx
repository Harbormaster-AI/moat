import React, { Component } from 'react'
import EquityGrantService from '../services/EquityGrantService'

class ViewEquityGrantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            equityGrant: {}
        }
    }

    componentDidMount(){
        EquityGrantService.getEquityGrantById(this.state.id).then( res => {
            this.setState({equityGrant: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View EquityGrant Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> grantId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.equityGrant.grantId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> grantedUnits:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.equityGrant.grantedUnits }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> vestingStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.equityGrant.vestingStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> GrantType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.equityGrant.grantType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEquityGrantComponent
