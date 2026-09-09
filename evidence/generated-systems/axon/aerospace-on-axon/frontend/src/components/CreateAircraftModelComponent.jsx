import React, { Component } from 'react'
import AircraftModelService from '../services/AircraftModelService';

class CreateAircraftModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                modelDesignation: '',
                aircraftType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changemodelDesignationHandler = this.changemodelDesignationHandler.bind(this);
        this.changeAircraftTypeHandler = this.changeAircraftTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftModelService.getAircraftModelById(this.state.id).then( (res) =>{
                let aircraftModel = res.data;
                this.setState({
                    name: aircraftModel.name,
                    modelDesignation: aircraftModel.modelDesignation,
                    aircraftType: aircraftModel.aircraftType
                });
            });
        }        
    }
    saveOrUpdateAircraftModel = (e) => {
        e.preventDefault();
        let aircraftModel = {
                aircraftModelId: this.state.id,
                name: this.state.name,
                modelDesignation: this.state.modelDesignation,
                aircraftType: this.state.aircraftType
            };
        console.log('aircraftModel => ' + JSON.stringify(aircraftModel));

        // step 5
        if(this.state.id === '_add'){
            aircraftModel.aircraftModelId=''
            AircraftModelService.createAircraftModel(aircraftModel).then(res =>{
                this.props.history.push('/aircraftModels');
            });
        }else{
            AircraftModelService.updateAircraftModel(aircraftModel).then( res => {
                this.props.history.push('/aircraftModels');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AircraftModel</h3>
        }else{
            return <h3 className="text-center">Update AircraftModel</h3>
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

                                            <label> modelDesignation:&emsp; </label>
                                                <input placeholder="modelDesignation" name="modelDesignation" className="form-control" value={this.state.modelDesignation} onChange={this.changemodelDesignationHandler}/>

                                            <label> AircraftType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraftModel}>Save</button>
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

export default CreateAircraftModelComponent
