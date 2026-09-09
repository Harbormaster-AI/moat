import React, { Component } from 'react'
import CreativeAssetService from '../services/CreativeAssetService';

class CreateCreativeAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
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
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeclickUrlHandler = this.changeclickUrlHandler.bind(this);
        this.changelandingPageHandler = this.changelandingPageHandler.bind(this);
        this.changewidthHandler = this.changewidthHandler.bind(this);
        this.changeheightHandler = this.changeheightHandler.bind(this);
        this.changedurationSecondsHandler = this.changedurationSecondsHandler.bind(this);
        this.changeCreativeTypeHandler = this.changeCreativeTypeHandler.bind(this);
        this.changeAdFormatHandler = this.changeAdFormatHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCreativeAsset = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            creativeAsset.creativeAssetId=''
            CreativeAssetService.createCreativeAsset(creativeAsset).then(res =>{
                this.props.history.push('/creativeAssets');
            });
        }else{
            CreativeAssetService.updateCreativeAsset(creativeAsset).then( res => {
                this.props.history.push('/creativeAssets');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CreativeAsset</h3>
        }else{
            return <h3 className="text-center">Update CreativeAsset</h3>
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

                                            <label> clickUrl:&emsp; </label>
                                                <input placeholder="clickUrl" name="clickUrl" className="form-control" value={this.state.clickUrl} onChange={this.changeclickUrlHandler}/>

                                            <label> landingPage:&emsp; </label>
                                                <input placeholder="landingPage" name="landingPage" className="form-control" value={this.state.landingPage} onChange={this.changelandingPageHandler}/>

                                            <label> width:&emsp; </label>
                                                <input type="number" placeholder="width" name="width" className="form-control" value={this.state.width} onChange={this.changewidthHandler}/>

                                            <label> height:&emsp; </label>
                                                <input type="number" placeholder="height" name="height" className="form-control" value={this.state.height} onChange={this.changeheightHandler}/>

                                            <label> durationSeconds:&emsp; </label>
                                                <input type="number" placeholder="durationSeconds" name="durationSeconds" className="form-control" value={this.state.durationSeconds} onChange={this.changedurationSecondsHandler}/>

                                            <label> CreativeType:&emsp; </label>
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

                                            <label> AdFormat:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCreativeAsset}>Save</button>
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

export default CreateCreativeAssetComponent
