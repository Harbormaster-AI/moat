import React, { Component } from 'react'
import AircraftModelService from '../services/AircraftModelService';

class UpdateAircraftModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                modelDesignation: '',
                aircraftType: ''
        }
        this.updateAircraftModel = this.updateAircraftModel.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changemodelDesignationHandler = this.changemodelDesignationHandler.bind(this);
        this.changeAircraftTypeHandler = this.changeAircraftTypeHandler.bind(this);
    }

    componentDidMount(){
        AircraftModelService.getAircraftModelById(this.state.id).then( (res) =>{
            let aircraftModel = res.data;
            this.setState({
                name: aircraftModel.name,
                modelDesignation: aircraftModel.modelDesignation,
                aircraftType: aircraftModel.aircraftType
            });
        });
    }

    updateAircraftModel = (e) => {
        e.preventDefault();
        let aircraftModel = {
            aircraftModelId: this.state.id,
            name: this.state.name,
            modelDesignation: this.state.modelDesignation,
            aircraftType: this.state.aircraftType
        };
        console.log('aircraftModel => ' + JSON.stringify(aircraftModel));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftModelService.updateAircraftModel(aircraftModel).then( res => {
            this.props.history.push('/aircraftModels');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changemodelDesignationHandler= (event) => {
        this.setState({modelDesignation: event.target.value});
    }
    changeAircraftTypeHandler= (event) => {
        this.setState({aircraftType: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftModels');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AircraftModel</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> modelDesignation: </label>
                                                <input placeholder="modelDesignation" name="modelDesignation" className="form-control" value={this.state.modelDesignation} onChange={this.changemodelDesignationHandler}/>

                                            <label> AircraftType: </label>
                                                <select value={this.state.aircraftType} onChange={this.changeAircraftTypeHandler}>
                      <option name="AircraftType" className="form-control" >
                          NarrowBody
                      </option>
                      <option name="AircraftType" className="form-control" >
                          WideBody
                      </option>
                      <option name="AircraftType" className="form-control" >
                          RegionalJet
                      </option>
                      <option name="AircraftType" className="form-control" >
                          Turboprop
                      </option>
                      <option name="AircraftType" className="form-control" >
                          BusinessJet
                      </option>
                      <option name="AircraftType" className="form-control" >
                          Helicopter
                      </option>
                      <option name="AircraftType" className="form-control" >
                          eVTOL
                      </option>
                      <option name="AircraftType" className="form-control" >
                          CargoPlane
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAircraftModel}>Save</button>
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

export default UpdateAircraftModelComponent
