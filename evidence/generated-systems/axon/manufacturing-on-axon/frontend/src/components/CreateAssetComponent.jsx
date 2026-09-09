import React, { Component } from 'react'
import AssetService from '../services/AssetService';

class CreateAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                assetTag: '',
                assetName: '',
                commissioningDate: '',
                assetStatus: ''
        }
        this.changeassetTagHandler = this.changeassetTagHandler.bind(this);
        this.changeassetNameHandler = this.changeassetNameHandler.bind(this);
        this.changecommissioningDateHandler = this.changecommissioningDateHandler.bind(this);
        this.changeAssetStatusHandler = this.changeAssetStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AssetService.getAssetById(this.state.id).then( (res) =>{
                let asset = res.data;
                this.setState({
                    assetTag: asset.assetTag,
                    assetName: asset.assetName,
                    commissioningDate: asset.commissioningDate,
                    assetStatus: asset.assetStatus
                });
            });
        }        
    }
    saveOrUpdateAsset = (e) => {
        e.preventDefault();
        let asset = {
                assetId: this.state.id,
                assetTag: this.state.assetTag,
                assetName: this.state.assetName,
                commissioningDate: this.state.commissioningDate,
                assetStatus: this.state.assetStatus
            };
        console.log('asset => ' + JSON.stringify(asset));

        // step 5
        if(this.state.id === '_add'){
            asset.assetId=''
            AssetService.createAsset(asset).then(res =>{
                this.props.history.push('/assets');
            });
        }else{
            AssetService.updateAsset(asset).then( res => {
                this.props.history.push('/assets');
            });
        }
    }
    
    changeassetTagHandler= (event) => {
        this.setState({assetTag: event.target.value});
    }
    changeassetNameHandler= (event) => {
        this.setState({assetName: event.target.value});
    }
    changecommissioningDateHandler= (event) => {
        this.setState({commissioningDate: event.target.value});
    }
    changeAssetStatusHandler= (event) => {
        this.setState({assetStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/assets');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Asset</h3>
        }else{
            return <h3 className="text-center">Update Asset</h3>
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
                                            <label> assetTag:&emsp; </label>
                                                <input placeholder="assetTag" name="assetTag" className="form-control" value={this.state.assetTag} onChange={this.changeassetTagHandler}/>

                                            <label> assetName:&emsp; </label>
                                                <input placeholder="assetName" name="assetName" className="form-control" value={this.state.assetName} onChange={this.changeassetNameHandler}/>

                                            <label> commissioningDate:&emsp; </label>
                                                <input type="date" placeholder="commissioningDate" name="commissioningDate" className="form-control" value={this.state.commissioningDate} onChange={this.changecommissioningDateHandler}/>

                                            <label> AssetStatus:&emsp; </label>
                                                <select value={this.state.assetStatus} onChange={this.changeAssetStatusHandler}>
                      <option name="AssetStatus" className="form-control" >
                          Commissioned
                      </option>
                      <option name="AssetStatus" className="form-control" >
                          Available
                      </option>
                      <option name="AssetStatus" className="form-control" >
                          InMaintenance
                      </option>
                      <option name="AssetStatus" className="form-control" >
                          Down
                      </option>
                      <option name="AssetStatus" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAsset}>Save</button>
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

export default CreateAssetComponent
