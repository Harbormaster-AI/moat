import React, { Component } from 'react'
import CreativeAssetService from '../services/CreativeAssetService'

class ListCreativeAssetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                creativeAssets: []
        }
        this.addCreativeAsset = this.addCreativeAsset.bind(this);
        this.editCreativeAsset = this.editCreativeAsset.bind(this);
        this.deleteCreativeAsset = this.deleteCreativeAsset.bind(this);
    }

    deleteCreativeAsset(id){
        CreativeAssetService.deleteCreativeAsset(id).then( res => {
            this.setState({creativeAssets: this.state.creativeAssets.filter(creativeAsset => creativeAsset.creativeAssetId !== id)});
        });
    }
    viewCreativeAsset(id){
        this.props.history.push(`/view-creativeAsset/${id}`);
    }
    editCreativeAsset(id){
        this.props.history.push(`/add-creativeAsset/${id}`);
    }

    componentDidMount(){
        CreativeAssetService.getCreativeAssets().then((res) => {
            this.setState({ creativeAssets: res.data});
        });
    }

    addCreativeAsset(){
        this.props.history.push('/add-creativeAsset/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CreativeAsset List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCreativeAsset}> Add CreativeAsset</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ClickUrl </th>
                                    <th> LandingPage </th>
                                    <th> Width </th>
                                    <th> Height </th>
                                    <th> DurationSeconds </th>
                                    <th> CreativeType </th>
                                    <th> AdFormat </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.creativeAssets.map(
                                        creativeAsset => 
                                        <tr key = {creativeAsset.creativeAssetId}>
                                             <td> { creativeAsset.name } </td>
                                             <td> { creativeAsset.clickUrl } </td>
                                             <td> { creativeAsset.landingPage } </td>
                                             <td> { creativeAsset.width } </td>
                                             <td> { creativeAsset.height } </td>
                                             <td> { creativeAsset.durationSeconds } </td>
                                             <td> { creativeAsset.creativeType } </td>
                                             <td> { creativeAsset.adFormat } </td>
                                             <td>
                                                 <button onClick={ () => this.editCreativeAsset(creativeAsset.creativeAssetId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCreativeAsset(creativeAsset.creativeAssetId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCreativeAsset(creativeAsset.creativeAssetId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCreativeAssetComponent
