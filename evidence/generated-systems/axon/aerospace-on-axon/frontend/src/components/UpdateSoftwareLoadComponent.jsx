import React, { Component } from 'react'
import SoftwareLoadService from '../services/SoftwareLoadService';

class UpdateSoftwareLoadComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                version: '',
                loadType: ''
        }
        this.updateSoftwareLoad = this.updateSoftwareLoad.bind(this);

        this.changeversionHandler = this.changeversionHandler.bind(this);
        this.changeLoadTypeHandler = this.changeLoadTypeHandler.bind(this);
    }

    componentDidMount(){
        SoftwareLoadService.getSoftwareLoadById(this.state.id).then( (res) =>{
            let softwareLoad = res.data;
            this.setState({
                version: softwareLoad.version,
                loadType: softwareLoad.loadType
            });
        });
    }

    updateSoftwareLoad = (e) => {
        e.preventDefault();
        let softwareLoad = {
            softwareLoadId: this.state.id,
            version: this.state.version,
            loadType: this.state.loadType
        };
        console.log('softwareLoad => ' + JSON.stringify(softwareLoad));
        console.log('id => ' + JSON.stringify(this.state.id));
        SoftwareLoadService.updateSoftwareLoad(softwareLoad).then( res => {
            this.props.history.push('/softwareLoads');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SoftwareLoad</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> version: </label>
                                                <input placeholder="version" name="version" className="form-control" value={this.state.version} onChange={this.changeversionHandler}/>

                                            <label> LoadType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateSoftwareLoad}>Save</button>
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

export default UpdateSoftwareLoadComponent
