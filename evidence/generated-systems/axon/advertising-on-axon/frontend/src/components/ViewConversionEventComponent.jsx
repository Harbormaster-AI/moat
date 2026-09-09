import React, { Component } from 'react'
import ConversionEventService from '../services/ConversionEventService'

class ViewConversionEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            conversionEvent: {}
        }
    }

    componentDidMount(){
        ConversionEventService.getConversionEventById(this.state.id).then( res => {
            this.setState({conversionEvent: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ConversionEvent Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> timestamp:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.conversionEvent.timestamp }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> value:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.conversionEvent.value }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> EventType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.conversionEvent.eventType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AttributionModel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.conversionEvent.attributionModel }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewConversionEventComponent
