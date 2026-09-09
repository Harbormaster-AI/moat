import React, { Component } from 'react'
import SoftwareLoadService from '../services/SoftwareLoadService';

class CreateSoftwareLoadComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                version: '',
                loadType: ''
        }
        this.changeversionHandler = this.changeversionHandler.bind(this);
        this.changeLoadTypeHandler = this.changeLoadTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SoftwareLoadService.getSoftwareLoadById(this.state.id).then( (res) =>{
                let softwareLoad = res.data;
                this.setState({
                    version: softwareLoad.version,
                    loadType: softwareLoad.loadType
                });
            });
        }        
    }
    saveOrUpdateSoftwareLoad = (e) => {
        e.preventDefault();
        let softwareLoad = {
                softwareLoadId: this.state.id,
                version: this.state.version,
                loadType: this.state.loadType
            };
        console.log('softwareLoad => ' + JSON.stringify(softwareLoad));

        // step 5
        if(this.state.id === '_add'){
            softwareLoad.softwareLoadId=''
            SoftwareLoadService.createSoftwareLoad(softwareLoad).then(res =>{
                this.props.history.push('/softwareLoads');
            });
        }else{
            SoftwareLoadService.updateSoftwareLoad(softwareLoad).then( res => {
                this.props.history.push('/softwareLoads');
            });
        }
    }
    
    changeversionHandler= (event) => {
        this.setState({version: event.target.value});
    }
    changeLoadTypeHandler= (event) => {
        this.setState({loadType: event.target.value});
    }

    cancel(){
        this.props.history.push('/softwareLoads');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SoftwareLoad</h3>
        }else{
            return <h3 className="text-center">Update SoftwareLoad</h3>
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
                                            <label> version:&emsp; </label>
                                                <input placeholder="version" name="version" className="form-control" value={this.state.version} onChange={this.changeversionHandler}/>

                                            <label> LoadType:&emsp; </label>
                                                <select value={this.state.loadType} onChange={this.changeLoadTypeHandler}>
                      <option name="LoadType" className="form-control" >
                          FlightDeckSoftware
                      </option>
                      <option name="LoadType" className="form-control" >
                          MaintenanceTools
                      </option>
                      <option name="LoadType" className="form-control" >
                          CabinIFE
                      </option>
                      <option name="LoadType" className="form-control" >
                          ConnectivityModem
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSoftwareLoad}>Save</button>
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

export default CreateSoftwareLoadComponent
