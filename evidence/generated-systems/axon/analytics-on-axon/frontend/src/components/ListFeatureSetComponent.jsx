import React, { Component } from 'react'
import FeatureSetService from '../services/FeatureSetService'

class ListFeatureSetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                featureSets: []
        }
        this.addFeatureSet = this.addFeatureSet.bind(this);
        this.editFeatureSet = this.editFeatureSet.bind(this);
        this.deleteFeatureSet = this.deleteFeatureSet.bind(this);
    }

    deleteFeatureSet(id){
        FeatureSetService.deleteFeatureSet(id).then( res => {
            this.setState({featureSets: this.state.featureSets.filter(featureSet => featureSet.featureSetId !== id)});
        });
    }
    viewFeatureSet(id){
        this.props.history.push(`/view-featureSet/${id}`);
    }
    editFeatureSet(id){
        this.props.history.push(`/add-featureSet/${id}`);
    }

    componentDidMount(){
        FeatureSetService.getFeatureSets().then((res) => {
            this.setState({ featureSets: res.data});
        });
    }

    addFeatureSet(){
        this.props.history.push('/add-featureSet/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FeatureSet List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFeatureSet}> Add FeatureSet</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> RefreshSchedule </th>
                                    <th> StoreType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.featureSets.map(
                                        featureSet => 
                                        <tr key = {featureSet.featureSetId}>
                                             <td> { featureSet.name } </td>
                                             <td> { featureSet.refreshSchedule } </td>
                                             <td> { featureSet.storeType } </td>
                                             <td>
                                                 <button onClick={ () => this.editFeatureSet(featureSet.featureSetId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFeatureSet(featureSet.featureSetId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFeatureSet(featureSet.featureSetId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFeatureSetComponent
