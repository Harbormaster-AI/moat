import React, { Component } from 'react'
import ImagingCenterService from '../services/ImagingCenterService'

class ViewImagingCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            imagingCenter: {}
        }
    }

    componentDidMount(){
        ImagingCenterService.getImagingCenterById(this.state.id).then( res => {
            this.setState({imagingCenter: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ImagingCenter Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingCenter.name }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewImagingCenterComponent
