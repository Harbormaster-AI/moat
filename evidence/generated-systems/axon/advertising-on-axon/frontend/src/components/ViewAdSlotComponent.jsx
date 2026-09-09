import React, { Component } from 'react'
import AdSlotService from '../services/AdSlotService'

class ViewAdSlotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            adSlot: {}
        }
    }

    componentDidMount(){
        AdSlotService.getAdSlotById(this.state.id).then( res => {
            this.setState({adSlot: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AdSlot Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> slotCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adSlot.slotCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> width:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adSlot.width }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> height:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adSlot.height }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> floorPrice:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adSlot.floorPrice }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Format:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.adSlot.format }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAdSlotComponent
