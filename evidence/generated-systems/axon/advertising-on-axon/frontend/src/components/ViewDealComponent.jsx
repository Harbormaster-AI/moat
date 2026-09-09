import React, { Component } from 'react'
import DealService from '../services/DealService'

class ViewDealComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            deal: {}
        }
    }

    componentDidMount(){
        DealService.getDealById(this.state.id).then( res => {
            this.setState({deal: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Deal Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> floorPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.deal.floorPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DealType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.deal.dealType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDealComponent
