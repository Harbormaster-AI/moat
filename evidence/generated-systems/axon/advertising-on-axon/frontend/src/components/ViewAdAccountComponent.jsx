import React, { Component } from 'react'
import AdAccountService from '../services/AdAccountService'

class ViewAdAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            adAccount: {}
        }
    }

    componentDidMount(){
        AdAccountService.getAdAccountById(this.state.id).then( res => {
            this.setState({adAccount: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AdAccount Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adAccount.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adAccount.accountCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> defaultCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adAccount.defaultCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> defaultTimezone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adAccount.defaultTimezone }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAdAccountComponent
