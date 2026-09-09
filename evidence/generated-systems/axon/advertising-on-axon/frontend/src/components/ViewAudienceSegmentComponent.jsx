import React, { Component } from 'react'
import AudienceSegmentService from '../services/AudienceSegmentService'

class ViewAudienceSegmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            audienceSegment: {}
        }
    }

    componentDidMount(){
        AudienceSegmentService.getAudienceSegmentById(this.state.id).then( res => {
            this.setState({audienceSegment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AudienceSegment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.audienceSegment.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> estimatedReach:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.audienceSegment.estimatedReach }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.audienceSegment.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ProviderType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.audienceSegment.providerType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAudienceSegmentComponent
