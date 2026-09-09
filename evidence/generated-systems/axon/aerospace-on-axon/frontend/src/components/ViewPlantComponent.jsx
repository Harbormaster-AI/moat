import React, { Component } from 'react'
import PlantService from '../services/PlantService'

class ViewPlantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            plant: {}
        }
    }

    componentDidMount(){
        PlantService.getPlantById(this.state.id).then( res => {
            this.setState({plant: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Plant Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plant.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> plantCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plant.plantCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.plant.address }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPlantComponent
