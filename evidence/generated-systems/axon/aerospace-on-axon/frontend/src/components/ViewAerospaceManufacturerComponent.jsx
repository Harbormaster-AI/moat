import React, { Component } from 'react'
import AerospaceManufacturerService from '../services/AerospaceManufacturerService'

class ViewAerospaceManufacturerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aerospaceManufacturer: {}
        }
    }

    componentDidMount(){
        AerospaceManufacturerService.getAerospaceManufacturerById(this.state.id).then( res => {
            this.setState({aerospaceManufacturer: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AerospaceManufacturer Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aerospaceManufacturer.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> legalName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aerospaceManufacturer.legalName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> headquartersCountry:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aerospaceManufacturer.headquartersCountry }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aerospaceManufacturer.website }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAerospaceManufacturerComponent
