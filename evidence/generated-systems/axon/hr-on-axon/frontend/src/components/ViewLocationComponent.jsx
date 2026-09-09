import React, { Component } from 'react'
import LocationService from '../services/LocationService'

class ViewLocationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            location: {}
        }
    }

    componentDidMount(){
        LocationService.getLocationById(this.state.id).then( res => {
            this.setState({location: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Location Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.location.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.location.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> timezone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.location.timezone }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLocationComponent
