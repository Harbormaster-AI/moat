import React, { Component } from 'react'
import DataSourceService from '../services/DataSourceService'

class ListDataSourceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataSources: []
        }
        this.addDataSource = this.addDataSource.bind(this);
        this.editDataSource = this.editDataSource.bind(this);
        this.deleteDataSource = this.deleteDataSource.bind(this);
    }

    deleteDataSource(id){
        DataSourceService.deleteDataSource(id).then( res => {
            this.setState({dataSources: this.state.dataSources.filter(dataSource => dataSource.dataSourceId !== id)});
        });
    }
    viewDataSource(id){
        this.props.history.push(`/view-dataSource/${id}`);
    }
    editDataSource(id){
        this.props.history.push(`/add-dataSource/${id}`);
    }

    componentDidMount(){
        DataSourceService.getDataSources().then((res) => {
            this.setState({ dataSources: res.data});
        });
    }

    addDataSource(){
        this.props.history.push('/add-dataSource/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataSource List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataSource}> Add DataSource</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Connection </th>
                                    <th> Streaming </th>
                                    <th> SourceType </th>
                                    <th> Format </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataSources.map(
                                        dataSource => 
                                        <tr key = {dataSource.dataSourceId}>
                                             <td> { dataSource.name } </td>
                                             <td> { dataSource.connection } </td>
                                             <td> { dataSource.streaming } </td>
                                             <td> { dataSource.sourceType } </td>
                                             <td> { dataSource.format } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataSource(dataSource.dataSourceId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataSource(dataSource.dataSourceId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataSource(dataSource.dataSourceId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataSourceComponent
