import React, { Component } from 'react'
import AssetService from '../services/AssetService';

class UpdateAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                assetTag: '',
                assetName: '',
                commissioningDate: '',
                assetStatus: ''
        }
        this.updateAsset = this.updateAsset.bind(this);

        this.changeassetTagHandler = this.changeassetTagHandler.bind(this);
        this.changeassetNameHandler = this.changeassetNameHandler.bind(this);
        this.changecommissioningDateHandler = this.changecommissioningDateHandler.bind(this);
        this.changeAssetStatusHandler = this.changeAssetStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateAsset = (e) => {
        e.preventDefault();
        let asset = {
            assetId: this.state.id,
            assetTag: this.state.assetTag,
            assetName: this.state.assetName,
            commissioningDate: this.state.commissioningDate,
            assetStatus: this.state.assetStatus
        };
        console.log('asset => ' + JSON.stringify(asset));
        console.log('id => ' + JSON.stringify(this.state.id));
        AssetService.updateAsset(asset).then( res => {
            this.props.history.push('/assets');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Asset</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> assetTag: </label>
                                                <input placeholder="assetTag" name="assetTag" className="form-control" value={this.state.assetTag} onChange={this.changeassetTagHandler}/>

                                            <label> assetName: </label>
                                                <input placeholder="assetName" name="assetName" className="form-control" value={this.state.assetName} onChange={this.changeassetNameHandler}/>

                                            <label> commissioningDate: </label>
                                                <input type="date" placeholder="commissioningDate" name="commissioningDate" className="form-control" value={this.state.commissioningDate} onChange={this.changecommissioningDateHandler}/>

                                            <label> AssetStatus: </label>
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
                                        <button className="btn btn-success" onClick={this.updateAsset}>Save</button>
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

export default UpdateAssetComponent
