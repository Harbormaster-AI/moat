import React, { Component } from 'react'
import CreativeAssetService from '../services/CreativeAssetService'

class ViewCreativeAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            creativeAsset: {}
        }
    }

    componentDidMount(){
        CreativeAssetService.getCreativeAssetById(this.state.id).then( res => {
            this.setState({creativeAsset: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CreativeAsset Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> clickUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.clickUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> landingPage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.landingPage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> width:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.width }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> height:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.height }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> durationSeconds:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.durationSeconds }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CreativeType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.creativeType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AdFormat:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeAsset.adFormat }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCreativeAssetComponent
