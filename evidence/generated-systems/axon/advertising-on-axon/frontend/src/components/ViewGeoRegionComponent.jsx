import React, { Component } from 'react'
import GeoRegionService from '../services/GeoRegionService'

class ViewGeoRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            geoRegion: {}
        }
    }

    componentDidMount(){
        GeoRegionService.getGeoRegionById(this.state.id).then( res => {
            this.setState({geoRegion: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View GeoRegion Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.geoRegion.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.geoRegion.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RegionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.geoRegion.regionType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewGeoRegionComponent
