import React, { Component } from 'react'
import MediaAssetService from '../services/MediaAssetService'

class ListMediaAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                mediaAssets: []
        }
        this.addMediaAsset = this.addMediaAsset.bind(this);
        this.editMediaAsset = this.editMediaAsset.bind(this);
        this.deleteMediaAsset = this.deleteMediaAsset.bind(this);
    }

    deleteMediaAsset(id){
        MediaAssetService.deleteMediaAsset(id).then( res => {
            this.setState({mediaAssets: this.state.mediaAssets.filter(mediaAsset => mediaAsset.mediaAssetId !== id)});
        });
    }
    viewMediaAsset(id){
        this.props.history.push(`/view-mediaAsset/${id}`);
    }
    editMediaAsset(id){
        this.props.history.push(`/add-mediaAsset/${id}`);
    }

    componentDidMount(){
        MediaAssetService.getMediaAssets().then((res) => {
            this.setState({ mediaAssets: res.data});
        });
    }

    addMediaAsset(){
        this.props.history.push('/add-mediaAsset/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MediaAsset List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMediaAsset}> Add MediaAsset</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Url </th>
                                    <th> AltText </th>
                                    <th> Position </th>
                                    <th> MediaType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.mediaAssets.map(
                                        mediaAsset => 
                                        <tr key = {mediaAsset.mediaAssetId}>
                                             <td> { mediaAsset.url } </td>
                                             <td> { mediaAsset.altText } </td>
                                             <td> { mediaAsset.position } </td>
                                             <td> { mediaAsset.mediaType } </td>
                                             <td>
                                                 <button onClick={ () => this.editMediaAsset(mediaAsset.mediaAssetId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMediaAsset(mediaAsset.mediaAssetId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMediaAsset(mediaAsset.mediaAssetId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListMediaAssetComponent
