import React, { Component } from 'react'
import PlacementService from '../services/PlacementService'

class ViewPlacementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            placement: {}
        }
    }

    componentDidMount(){
        PlacementService.getPlacementById(this.state.id).then( res => {
            this.setState({placement: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Placement Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.placement.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> flight:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.placement.flight }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> goalImpressions:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.placement.goalImpressions }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPlacementComponent
