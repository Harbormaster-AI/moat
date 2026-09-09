import React, { Component } from 'react'
import AircraftFamilyService from '../services/AircraftFamilyService';

class CreateAircraftFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                familyCode: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changefamilyCodeHandler = this.changefamilyCodeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftFamilyService.getAircraftFamilyById(this.state.id).then( (res) =>{
                let aircraftFamily = res.data;
                this.setState({
                    name: aircraftFamily.name,
                    familyCode: aircraftFamily.familyCode
                });
            });
        }        
    }
    saveOrUpdateAircraftFamily = (e) => {
        e.preventDefault();
        let aircraftFamily = {
                aircraftFamilyId: this.state.id,
                name: this.state.name,
                familyCode: this.state.familyCode
            };
        console.log('aircraftFamily => ' + JSON.stringify(aircraftFamily));

        // step 5
        if(this.state.id === '_add'){
            aircraftFamily.aircraftFamilyId=''
            AircraftFamilyService.createAircraftFamily(aircraftFamily).then(res =>{
                this.props.history.push('/aircraftFamilys');
            });
        }else{
            AircraftFamilyService.updateAircraftFamily(aircraftFamily).then( res => {
                this.props.history.push('/aircraftFamilys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changefamilyCodeHandler= (event) => {
        this.setState({familyCode: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftFamilys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AircraftFamily</h3>
        }else{
            return <h3 className="text-center">Update AircraftFamily</h3>
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

                                            <label> familyCode:&emsp; </label>
                                                <input placeholder="familyCode" name="familyCode" className="form-control" value={this.state.familyCode} onChange={this.changefamilyCodeHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraftFamily}>Save</button>
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

export default CreateAircraftFamilyComponent
