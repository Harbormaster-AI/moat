import React, { Component } from 'react'
import SoftwareUpdateService from '../services/SoftwareUpdateService';

class CreateSoftwareUpdateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                version: '',
                appliedDate: '',
                updateType: ''
        }
        this.changeversionHandler = this.changeversionHandler.bind(this);
        this.changeappliedDateHandler = this.changeappliedDateHandler.bind(this);
        this.changeUpdateTypeHandler = this.changeUpdateTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SoftwareUpdateService.getSoftwareUpdateById(this.state.id).then( (res) =>{
                let softwareUpdate = res.data;
                this.setState({
                    version: softwareUpdate.version,
                    appliedDate: softwareUpdate.appliedDate,
                    updateType: softwareUpdate.updateType
                });
            });
        }        
    }
    saveOrUpdateSoftwareUpdate = (e) => {
        e.preventDefault();
        let softwareUpdate = {
                softwareUpdateId: this.state.id,
                version: this.state.version,
                appliedDate: this.state.appliedDate,
                updateType: this.state.updateType
            };
        console.log('softwareUpdate => ' + JSON.stringify(softwareUpdate));

        // step 5
        if(this.state.id === '_add'){
            softwareUpdate.softwareUpdateId=''
            SoftwareUpdateService.createSoftwareUpdate(softwareUpdate).then(res =>{
                this.props.history.push('/softwareUpdates');
            });
        }else{
            SoftwareUpdateService.updateSoftwareUpdate(softwareUpdate).then( res => {
                this.props.history.push('/softwareUpdates');
            });
        }
    }
    
    changeversionHandler= (event) => {
        this.setState({version: event.target.value});
    }
    changeappliedDateHandler= (event) => {
        this.setState({appliedDate: event.target.value});
    }
    changeUpdateTypeHandler= (event) => {
        this.setState({updateType: event.target.value});
    }

    cancel(){
        this.props.history.push('/softwareUpdates');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SoftwareUpdate</h3>
        }else{
            return <h3 className="text-center">Update SoftwareUpdate</h3>
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

                                            <label> appliedDate:&emsp; </label>
                                                <input type="time" placeholder="appliedDate" name="appliedDate" className="form-control" value={this.state.appliedDate} onChange={this.changeappliedDateHandler}/>

                                            <label> UpdateType:&emsp; </label>
                                                <select value={this.state.updateType} onChange={this.changeUpdateTypeHandler}>
                      <option name="UpdateType" className="form-control" >
                          SecurityPatch
                      </option>
                      <option name="UpdateType" className="form-control" >
                          FeatureUpdate
                      </option>
                      <option name="UpdateType" className="form-control" >
                          BugFix
                      </option>
                      <option name="UpdateType" className="form-control" >
                          FirmwareUpgrade
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSoftwareUpdate}>Save</button>
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

export default CreateSoftwareUpdateComponent
