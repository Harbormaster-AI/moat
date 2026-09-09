import React, { Component } from 'react'
import AssetService from '../services/AssetService'

class ListAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                assets: []
        }
        this.addAsset = this.addAsset.bind(this);
        this.editAsset = this.editAsset.bind(this);
        this.deleteAsset = this.deleteAsset.bind(this);
    }

    deleteAsset(id){
        AssetService.deleteAsset(id).then( res => {
            this.setState({assets: this.state.assets.filter(asset => asset.assetId !== id)});
        });
    }
    viewAsset(id){
        this.props.history.push(`/view-asset/${id}`);
    }
    editAsset(id){
        this.props.history.push(`/add-asset/${id}`);
    }

    componentDidMount(){
        AssetService.getAssets().then((res) => {
            this.setState({ assets: res.data});
        });
    }

    addAsset(){
        this.props.history.push('/add-asset/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Asset List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAsset}> Add Asset</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AssetTag </th>
                                    <th> AssetName </th>
                                    <th> CommissioningDate </th>
                                    <th> AssetStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.assets.map(
                                        asset => 
                                        <tr key = {asset.assetId}>
                                             <td> { asset.assetTag } </td>
                                             <td> { asset.assetName } </td>
                                             <td> { asset.commissioningDate } </td>
                                             <td> { asset.assetStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editAsset(asset.assetId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAsset(asset.assetId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAsset(asset.assetId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAssetComponent
