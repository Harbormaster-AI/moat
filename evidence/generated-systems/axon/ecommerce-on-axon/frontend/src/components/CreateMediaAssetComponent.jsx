import React, { Component } from 'react'
import MediaAssetService from '../services/MediaAssetService';

class CreateMediaAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                url: '',
                altText: '',
                position: '',
                mediaType: ''
        }
        this.changeurlHandler = this.changeurlHandler.bind(this);
        this.changealtTextHandler = this.changealtTextHandler.bind(this);
        this.changepositionHandler = this.changepositionHandler.bind(this);
        this.changeMediaTypeHandler = this.changeMediaTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MediaAssetService.getMediaAssetById(this.state.id).then( (res) =>{
                let mediaAsset = res.data;
                this.setState({
                    url: mediaAsset.url,
                    altText: mediaAsset.altText,
                    position: mediaAsset.position,
                    mediaType: mediaAsset.mediaType
                });
            });
        }        
    }
    saveOrUpdateMediaAsset = (e) => {
        e.preventDefault();
        let mediaAsset = {
                mediaAssetId: this.state.id,
                url: this.state.url,
                altText: this.state.altText,
                position: this.state.position,
                mediaType: this.state.mediaType
            };
        console.log('mediaAsset => ' + JSON.stringify(mediaAsset));

        // step 5
        if(this.state.id === '_add'){
            mediaAsset.mediaAssetId=''
            MediaAssetService.createMediaAsset(mediaAsset).then(res =>{
                this.props.history.push('/mediaAssets');
            });
        }else{
            MediaAssetService.updateMediaAsset(mediaAsset).then( res => {
                this.props.history.push('/mediaAssets');
            });
        }
    }
    
    changeurlHandler= (event) => {
        this.setState({url: event.target.value});
    }
    changealtTextHandler= (event) => {
        this.setState({altText: event.target.value});
    }
    changepositionHandler= (event) => {
        this.setState({position: event.target.value});
    }
    changeMediaTypeHandler= (event) => {
        this.setState({mediaType: event.target.value});
    }

    cancel(){
        this.props.history.push('/mediaAssets');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MediaAsset</h3>
        }else{
            return <h3 className="text-center">Update MediaAsset</h3>
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
                                            <label> url:&emsp; </label>
                                                <input placeholder="url" name="url" className="form-control" value={this.state.url} onChange={this.changeurlHandler}/>

                                            <label> altText:&emsp; </label>
                                                <input placeholder="altText" name="altText" className="form-control" value={this.state.altText} onChange={this.changealtTextHandler}/>

                                            <label> position:&emsp; </label>
                                                <input type="number" placeholder="position" name="position" className="form-control" value={this.state.position} onChange={this.changepositionHandler}/>

                                            <label> MediaType:&emsp; </label>
                                                <select value={this.state.mediaType} onChange={this.changeMediaTypeHandler}>
                      <option name="MediaType" className="form-control" >
                          Image
                      </option>
                      <option name="MediaType" className="form-control" >
                          Video
                      </option>
                      <option name="MediaType" className="form-control" >
                          Document
                      </option>
                      <option name="MediaType" className="form-control" >
                          Audio
                      </option>
                      <option name="MediaType" className="form-control" >
                          Other
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMediaAsset}>Save</button>
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

export default CreateMediaAssetComponent
