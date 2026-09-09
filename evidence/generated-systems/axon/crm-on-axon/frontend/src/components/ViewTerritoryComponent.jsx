import React, { Component } from 'react'
import TerritoryService from '../services/TerritoryService'

class ViewTerritoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            territory: {}
        }
    }

    componentDidMount(){
        TerritoryService.getTerritoryById(this.state.id).then( res => {
            this.setState({territory: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Territory Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.territory.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> region:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.territory.region }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TerritoryType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.territory.territoryType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTerritoryComponent
