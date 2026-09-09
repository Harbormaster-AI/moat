import React, { Component } from 'react'
import CreativeAssetService from '../services/CreativeAssetService';

class UpdateCreativeAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                clickUrl: '',
                landingPage: '',
                width: '',
                height: '',
                durationSeconds: '',
                creativeType: '',
                adFormat: ''
        }
        this.updateCreativeAsset = this.updateCreativeAsset.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeclickUrlHandler = this.changeclickUrlHandler.bind(this);
        this.changelandingPageHandler = this.changelandingPageHandler.bind(this);
        this.changewidthHandler = this.changewidthHandler.bind(this);
        this.changeheightHandler = this.changeheightHandler.bind(this);
        this.changedurationSecondsHandler = this.changedurationSecondsHandler.bind(this);
        this.changeCreativeTypeHandler = this.changeCreativeTypeHandler.bind(this);
        this.changeAdFormatHandler = this.changeAdFormatHandler.bind(this);
    }

    componentDidMount(){
        CreativeAssetService.getCreativeAssetById(this.state.id).then( (res) =>{
            let creativeAsset = res.data;
            this.setState({
                name: creativeAsset.name,
                clickUrl: creativeAsset.clickUrl,
                landingPage: creativeAsset.landingPage,
                width: creativeAsset.width,
                height: creativeAsset.height,
                durationSeconds: creativeAsset.durationSeconds,
                creativeType: creativeAsset.creativeType,
                adFormat: creativeAsset.adFormat
            });
        });
    }

    updateCreativeAsset = (e) => {
        e.preventDefault();
        let creativeAsset = {
            creativeAssetId: this.state.id,
            name: this.state.name,
            clickUrl: this.state.clickUrl,
            landingPage: this.state.landingPage,
            width: this.state.width,
            height: this.state.height,
            durationSeconds: this.state.durationSeconds,
            creativeType: this.state.creativeType,
            adFormat: this.state.adFormat
        };
        console.log('creativeAsset => ' + JSON.stringify(creativeAsset));
        console.log('id => ' + JSON.stringify(this.state.id));
        CreativeAssetService.updateCreativeAsset(creativeAsset).then( res => {
            this.props.history.push('/creativeAssets');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeclickUrlHandler= (event) => {
        this.setState({clickUrl: event.target.value});
    }
    changelandingPageHandler= (event) => {
        this.setState({landingPage: event.target.value});
    }
    changewidthHandler= (event) => {
        this.setState({width: event.target.value});
    }
    changeheightHandler= (event) => {
        this.setState({height: event.target.value});
    }
    changedurationSecondsHandler= (event) => {
        this.setState({durationSeconds: event.target.value});
    }
    changeCreativeTypeHandler= (event) => {
        this.setState({creativeType: event.target.value});
    }
    changeAdFormatHandler= (event) => {
        this.setState({adFormat: event.target.value});
    }

    cancel(){
        this.props.history.push('/creativeAssets');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CreativeAsset</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> clickUrl: </label>
                                                <input placeholder="clickUrl" name="clickUrl" className="form-control" value={this.state.clickUrl} onChange={this.changeclickUrlHandler}/>

                                            <label> landingPage: </label>
                                                <input placeholder="landingPage" name="landingPage" className="form-control" value={this.state.landingPage} onChange={this.changelandingPageHandler}/>

                                            <label> width: </label>
                                                <input type="number" placeholder="width" name="width" className="form-control" value={this.state.width} onChange={this.changewidthHandler}/>

                                            <label> height: </label>
                                                <input type="number" placeholder="height" name="height" className="form-control" value={this.state.height} onChange={this.changeheightHandler}/>

                                            <label> durationSeconds: </label>
                                                <input type="number" placeholder="durationSeconds" name="durationSeconds" className="form-control" value={this.state.durationSeconds} onChange={this.changedurationSecondsHandler}/>

                                            <label> CreativeType: </label>
                                                <select value={this.state.creativeType} onChange={this.changeCreativeTypeHandler}>
                      <option name="CreativeType" className="form-control" >
                          Image
                      </option>
                      <option name="CreativeType" className="form-control" >
                          Video
                      </option>
                      <option name="CreativeType" className="form-control" >
                          HTML5
                      </option>
                      <option name="CreativeType" className="form-control" >
                          Audio
                      </option>
                    </select>

                                            <label> AdFormat: </label>
                                                <select value={this.state.adFormat} onChange={this.changeAdFormatHandler}>
                      <option name="AdFormat" className="form-control" >
                          Banner
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Video
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Native
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Audio
                      </option>
                      <option name="AdFormat" className="form-control" >
                          Interstitial
                      </option>
                      <option name="AdFormat" className="form-control" >
                          RichMedia
                      </option>
                      <option name="AdFormat" className="form-control" >
                          SearchText
                      </option>
                      <option name="AdFormat" className="form-control" >
                          SocialPost
                      </option>
                      <option name="AdFormat" className="form-control" >
                          CTVVideo
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCreativeAsset}>Save</button>
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

export default UpdateCreativeAssetComponent
