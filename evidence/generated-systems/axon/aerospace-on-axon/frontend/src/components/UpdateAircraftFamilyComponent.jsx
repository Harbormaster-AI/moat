import React, { Component } from 'react'
import AircraftFamilyService from '../services/AircraftFamilyService';

class UpdateAircraftFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                familyCode: ''
        }
        this.updateAircraftFamily = this.updateAircraftFamily.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changefamilyCodeHandler = this.changefamilyCodeHandler.bind(this);
    }

    componentDidMount(){
        AircraftFamilyService.getAircraftFamilyById(this.state.id).then( (res) =>{
            let aircraftFamily = res.data;
            this.setState({
                name: aircraftFamily.name,
                familyCode: aircraftFamily.familyCode
            });
        });
    }

    updateAircraftFamily = (e) => {
        e.preventDefault();
        let aircraftFamily = {
            aircraftFamilyId: this.state.id,
            name: this.state.name,
            familyCode: this.state.familyCode
        };
        console.log('aircraftFamily => ' + JSON.stringify(aircraftFamily));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftFamilyService.updateAircraftFamily(aircraftFamily).then( res => {
            this.props.history.push('/aircraftFamilys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AircraftFamily</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> familyCode: </label>
                                                <input placeholder="familyCode" name="familyCode" className="form-control" value={this.state.familyCode} onChange={this.changefamilyCodeHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAircraftFamily}>Save</button>
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

export default UpdateAircraftFamilyComponent
