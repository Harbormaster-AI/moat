import React, { Component } from 'react'
import WalletService from '../services/WalletService'

class ViewWalletComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            wallet: {}
        }
    }

    componentDidMount(){
        WalletService.getWalletById(this.state.id).then( res => {
            this.setState({wallet: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Wallet Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> currency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.wallet.currency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> balance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.wallet.balance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.wallet.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWalletComponent
