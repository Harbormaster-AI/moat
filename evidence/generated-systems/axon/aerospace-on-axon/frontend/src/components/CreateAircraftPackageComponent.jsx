import React, { Component } from 'react'
import AircraftPackageService from '../services/AircraftPackageService';

class CreateAircraftPackageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                packageType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changePackageTypeHandler = this.changePackageTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftPackageService.getAircraftPackageById(this.state.id).then( (res) =>{
                let aircraftPackage = res.data;
                this.setState({
                    name: aircraftPackage.name,
                    packageType: aircraftPackage.packageType
                });
            });
        }        
    }
    saveOrUpdateAircraftPackage = (e) => {
        e.preventDefault();
        let aircraftPackage = {
                aircraftPackageId: this.state.id,
                name: this.state.name,
                packageType: this.state.packageType
            };
        console.log('aircraftPackage => ' + JSON.stringify(aircraftPackage));

        // step 5
        if(this.state.id === '_add'){
            aircraftPackage.aircraftPackageId=''
            AircraftPackageService.createAircraftPackage(aircraftPackage).then(res =>{
                this.props.history.push('/aircraftPackages');
            });
        }else{
            AircraftPackageService.updateAircraftPackage(aircraftPackage).then( res => {
                this.props.history.push('/aircraftPackages');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changePackageTypeHandler= (event) => {
        this.setState({packageType: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftPackages');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AircraftPackage</h3>
        }else{
            return <h3 className="text-center">Update AircraftPackage</h3>
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

                                            <label> PackageType:&emsp; </label>
                                                <select value={this.state.packageType} onChange={this.changePackageTypeHandler}>
                      <option name="PackageType" className="form-control" >
                          PerformancePack
                      </option>
                      <option name="PackageType" className="form-control" >
                          CabinPack
                      </option>
                      <option name="PackageType" className="form-control" >
                          ConnectivityPack
                      </option>
                      <option name="PackageType" className="form-control" >
                          CompliancePack
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraftPackage}>Save</button>
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

export default CreateAircraftPackageComponent
