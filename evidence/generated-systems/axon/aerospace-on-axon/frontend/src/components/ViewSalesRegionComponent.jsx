import React, { Component } from 'react'
import SalesRegionService from '../services/SalesRegionService'

class ViewSalesRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            salesRegion: {}
        }
    }

    componentDidMount(){
        SalesRegionService.getSalesRegionById(this.state.id).then( res => {
            this.setState({salesRegion: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SalesRegion Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesRegion.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> regionCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesRegion.regionCode }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSalesRegionComponent
