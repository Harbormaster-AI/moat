import React, { Component } from 'react'
import MediaAssetService from '../services/MediaAssetService';

class UpdateMediaAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                url: '',
                altText: '',
                position: '',
                mediaType: ''
        }
        this.updateMediaAsset = this.updateMediaAsset.bind(this);

        this.changeurlHandler = this.changeurlHandler.bind(this);
        this.changealtTextHandler = this.changealtTextHandler.bind(this);
        this.changepositionHandler = this.changepositionHandler.bind(this);
        this.changeMediaTypeHandler = this.changeMediaTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateMediaAsset = (e) => {
        e.preventDefault();
        let mediaAsset = {
            mediaAssetId: this.state.id,
            url: this.state.url,
            altText: this.state.altText,
            position: this.state.position,
            mediaType: this.state.mediaType
        };
        console.log('mediaAsset => ' + JSON.stringify(mediaAsset));
        console.log('id => ' + JSON.stringify(this.state.id));
        MediaAssetService.updateMediaAsset(mediaAsset).then( res => {
            this.props.history.push('/mediaAssets');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MediaAsset</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> url: </label>
                                                <input placeholder="url" name="url" className="form-control" value={this.state.url} onChange={this.changeurlHandler}/>

                                            <label> altText: </label>
                                                <input placeholder="altText" name="altText" className="form-control" value={this.state.altText} onChange={this.changealtTextHandler}/>

                                            <label> position: </label>
                                                <input type="number" placeholder="position" name="position" className="form-control" value={this.state.position} onChange={this.changepositionHandler}/>

                                            <label> MediaType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateMediaAsset}>Save</button>
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

export default UpdateMediaAssetComponent
