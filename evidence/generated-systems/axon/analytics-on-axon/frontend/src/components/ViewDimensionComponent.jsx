import React, { Component } from 'react'
import DimensionService from '../services/DimensionService'

class ViewDimensionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dimension: {}
        }
    }

    componentDidMount(){
        DimensionService.getDimensionById(this.state.id).then( res => {
            this.setState({dimension: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Dimension Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dimension.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> typeTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dimension.typeTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DimensionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dimension.dimensionType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDimensionComponent
