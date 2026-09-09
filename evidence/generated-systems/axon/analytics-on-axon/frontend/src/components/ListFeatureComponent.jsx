import React, { Component } from 'react'
import FeatureService from '../services/FeatureService'

class ListFeatureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                features: []
        }
        this.addFeature = this.addFeature.bind(this);
        this.editFeature = this.editFeature.bind(this);
        this.deleteFeature = this.deleteFeature.bind(this);
    }

    deleteFeature(id){
        FeatureService.deleteFeature(id).then( res => {
            this.setState({features: this.state.features.filter(feature => feature.featureId !== id)});
        });
    }
    viewFeature(id){
        this.props.history.push(`/view-feature/${id}`);
    }
    editFeature(id){
        this.props.history.push(`/add-feature/${id}`);
    }

    componentDidMount(){
        FeatureService.getFeatures().then((res) => {
            this.setState({ features: res.data});
        });
    }

    addFeature(){
        this.props.history.push('/add-feature/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Feature List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFeature}> Add Feature</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Description </th>
                                    <th> DataType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.features.map(
                                        feature => 
                                        <tr key = {feature.featureId}>
                                             <td> { feature.name } </td>
                                             <td> { feature.description } </td>
                                             <td> { feature.dataType } </td>
                                             <td>
                                                 <button onClick={ () => this.editFeature(feature.featureId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFeature(feature.featureId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFeature(feature.featureId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFeatureComponent
