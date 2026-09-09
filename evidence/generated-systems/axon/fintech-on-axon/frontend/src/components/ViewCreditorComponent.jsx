import React, { Component } from 'react'
import CreditorService from '../services/CreditorService'

class ViewCreditorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            creditor: {}
        }
    }

    componentDidMount(){
        CreditorService.getCreditorById(this.state.id).then( res => {
            this.setState({creditor: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Creditor Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creditor.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bic:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creditor.bic }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creditor.address }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCreditorComponent
