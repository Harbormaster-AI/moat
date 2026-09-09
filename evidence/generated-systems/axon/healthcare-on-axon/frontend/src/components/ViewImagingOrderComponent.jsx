import React, { Component } from 'react'
import ImagingOrderService from '../services/ImagingOrderService'

class ViewImagingOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            imagingOrder: {}
        }
    }

    componentDidMount(){
        ImagingOrderService.getImagingOrderById(this.state.id).then( res => {
            this.setState({imagingOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ImagingOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bodySite:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingOrder.bodySite }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> contrast:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingOrder.contrast }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Modality:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingOrder.modality }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewImagingOrderComponent
