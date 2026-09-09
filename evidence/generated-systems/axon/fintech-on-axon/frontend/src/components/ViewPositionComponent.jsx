import React, { Component } from 'react'
import PositionService from '../services/PositionService'

class ViewPositionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            position: {}
        }
    }

    componentDidMount(){
        PositionService.getPositionById(this.state.id).then( res => {
            this.setState({position: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Position Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.position.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> averageCost:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.position.averageCost }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> marketValue:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.position.marketValue }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPositionComponent
