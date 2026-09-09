import React, { Component } from 'react'
import CreativeVariationService from '../services/CreativeVariationService'

class ViewCreativeVariationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            creativeVariation: {}
        }
    }

    componentDidMount(){
        CreativeVariationService.getCreativeVariationById(this.state.id).then( res => {
            this.setState({creativeVariation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CreativeVariation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeVariation.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> language:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeVariation.language }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> headline:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeVariation.headline }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bodyText:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeVariation.bodyText }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> callToAction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeVariation.callToAction }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCreativeVariationComponent
