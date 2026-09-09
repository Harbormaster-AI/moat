import React, { Component } from 'react'
import DirectDebitMandateService from '../services/DirectDebitMandateService'

class ViewDirectDebitMandateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            directDebitMandate: {}
        }
    }

    componentDidMount(){
        DirectDebitMandateService.getDirectDebitMandateById(this.state.id).then( res => {
            this.setState({directDebitMandate: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DirectDebitMandate Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> mandateId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.directDebitMandate.mandateId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> signedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.directDebitMandate.signedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Scheme:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.directDebitMandate.scheme }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.directDebitMandate.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDirectDebitMandateComponent
