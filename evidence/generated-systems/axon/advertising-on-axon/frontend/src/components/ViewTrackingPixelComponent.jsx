import React, { Component } from 'react'
import TrackingPixelService from '../services/TrackingPixelService'

class ViewTrackingPixelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            trackingPixel: {}
        }
    }

    componentDidMount(){
        TrackingPixelService.getTrackingPixelById(this.state.id).then( res => {
            this.setState({trackingPixel: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TrackingPixel Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trackingPixel.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> url:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trackingPixel.url }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> EventType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trackingPixel.eventType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PixelType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trackingPixel.pixelType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTrackingPixelComponent
