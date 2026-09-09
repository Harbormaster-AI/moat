import React, { Component } from 'react'
import MediaAssetService from '../services/MediaAssetService'

class ViewMediaAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            mediaAsset: {}
        }
    }

    componentDidMount(){
        MediaAssetService.getMediaAssetById(this.state.id).then( res => {
            this.setState({mediaAsset: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MediaAsset Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> url:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mediaAsset.url }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> altText:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mediaAsset.altText }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> position:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mediaAsset.position }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MediaType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mediaAsset.mediaType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMediaAssetComponent
