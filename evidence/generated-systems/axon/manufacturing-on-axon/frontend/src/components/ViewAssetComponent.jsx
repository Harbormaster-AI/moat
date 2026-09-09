import React, { Component } from 'react'
import AssetService from '../services/AssetService'

class ViewAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            asset: {}
        }
    }

    componentDidMount(){
        AssetService.getAssetById(this.state.id).then( res => {
            this.setState({asset: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Asset Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assetTag:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.asset.assetTag }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> assetName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.asset.assetName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> commissioningDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.asset.commissioningDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AssetStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.asset.assetStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAssetComponent
