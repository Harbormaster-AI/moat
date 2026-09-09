import React, { Component } from 'react'
import PlantService from '../services/PlantService';

class CreatePlantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                plantCode: '',
                address: '',
                timeZone: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeplantCodeHandler = this.changeplantCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetimeZoneHandler = this.changetimeZoneHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PlantService.getPlantById(this.state.id).then( (res) =>{
                let plant = res.data;
                this.setState({
                    name: plant.name,
                    plantCode: plant.plantCode,
                    address: plant.address,
                    timeZone: plant.timeZone
                });
            });
        }        
    }
    saveOrUpdatePlant = (e) => {
        e.preventDefault();
        let plant = {
                plantId: this.state.id,
                name: this.state.name,
                plantCode: this.state.plantCode,
                address: this.state.address,
                timeZone: this.state.timeZone
            };
        console.log('plant => ' + JSON.stringify(plant));

        // step 5
        if(this.state.id === '_add'){
            plant.plantId=''
            PlantService.createPlant(plant).then(res =>{
                this.props.history.push('/plants');
            });
        }else{
            PlantService.updatePlant(plant).then( res => {
                this.props.history.push('/plants');
            });
        }
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
    changetimeZoneHandler= (event) => {
        this.setState({timeZone: event.target.value});
    }

    cancel(){
        this.props.history.push('/plants');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Plant</h3>
        }else{
            return <h3 className="text-center">Update Plant</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> plantCode:&emsp; </label>
                                                <input placeholder="plantCode" name="plantCode" className="form-control" value={this.state.plantCode} onChange={this.changeplantCodeHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> timeZone:&emsp; </label>
                                                <input placeholder="timeZone" name="timeZone" className="form-control" value={this.state.timeZone} onChange={this.changetimeZoneHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePlant}>Save</button>
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

export default CreatePlantComponent
