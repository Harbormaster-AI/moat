import React, { Component } from 'react'
import DataSetService from '../services/DataSetService'

class ListDataSetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataSets: []
        }
        this.addDataSet = this.addDataSet.bind(this);
        this.editDataSet = this.editDataSet.bind(this);
        this.deleteDataSet = this.deleteDataSet.bind(this);
    }

    deleteDataSet(id){
        DataSetService.deleteDataSet(id).then( res => {
            this.setState({dataSets: this.state.dataSets.filter(dataSet => dataSet.dataSetId !== id)});
        });
    }
    viewDataSet(id){
        this.props.history.push(`/view-dataSet/${id}`);
    }
    editDataSet(id){
        this.props.history.push(`/add-dataSet/${id}`);
    }

    componentDidMount(){
        DataSetService.getDataSets().then((res) => {
            this.setState({ dataSets: res.data});
        });
    }

    addDataSet(){
        this.props.history.push('/add-dataSet/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataSet List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataSet}> Add DataSet</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> SchemaVersion </th>
                                    <th> RefreshSchedule </th>
                                    <th> Sensitive </th>
                                    <th> DataFormat </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataSets.map(
                                        dataSet => 
                                        <tr key = {dataSet.dataSetId}>
                                             <td> { dataSet.name } </td>
                                             <td> { dataSet.schemaVersion } </td>
                                             <td> { dataSet.refreshSchedule } </td>
                                             <td> { dataSet.sensitive } </td>
                                             <td> { dataSet.dataFormat } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataSet(dataSet.dataSetId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataSet(dataSet.dataSetId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataSet(dataSet.dataSetId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataSetComponent
