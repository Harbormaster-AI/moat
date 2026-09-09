import React, { Component } from 'react'
import PlantService from '../services/PlantService';

class UpdatePlantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                plantCode: '',
                address: ''
        }
        this.updatePlant = this.updatePlant.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeplantCodeHandler = this.changeplantCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
    }

    componentDidMount(){
        PlantService.getPlantById(this.state.id).then( (res) =>{
            let plant = res.data;
            this.setState({
                name: plant.name,
                plantCode: plant.plantCode,
                address: plant.address
            });
        });
    }

    updatePlant = (e) => {
        e.preventDefault();
        let plant = {
            plantId: this.state.id,
            name: this.state.name,
            plantCode: this.state.plantCode,
            address: this.state.address
        };
        console.log('plant => ' + JSON.stringify(plant));
        console.log('id => ' + JSON.stringify(this.state.id));
        PlantService.updatePlant(plant).then( res => {
            this.props.history.push('/plants');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeplantCodeHandler= (event) => {
        this.setState({plantCode: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }

    cancel(){
        this.props.history.push('/plants');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Plant</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> plantCode: </label>
                                                <input placeholder="plantCode" name="plantCode" className="form-control" value={this.state.plantCode} onChange={this.changeplantCodeHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePlant}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdatePlantComponent
