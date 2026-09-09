import React, { Component } from 'react'
import AdvertiserService from '../services/AdvertiserService'

class ViewAdvertiserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            advertiser: {}
        }
    }

    componentDidMount(){
        AdvertiserService.getAdvertiserById(this.state.id).then( res => {
            this.setState({advertiser: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Advertiser Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.advertiser.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> legalName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.advertiser.legalName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> industry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.advertiser.industry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.advertiser.website }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAdvertiserComponent
