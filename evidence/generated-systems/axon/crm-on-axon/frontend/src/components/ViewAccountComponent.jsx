import React, { Component } from 'react'
import AccountService from '../services/AccountService'

class ViewAccountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            account: {}
        }
    }

    componentDidMount(){
        AccountService.getAccountById(this.state.id).then( res => {
            this.setState({account: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Account Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.accountNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> industry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.industry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> billingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.billingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shippingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.shippingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.website }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> phone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.phone }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.asActive }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AccountType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.accountType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LifecycleStage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.account.lifecycleStage }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAccountComponent
