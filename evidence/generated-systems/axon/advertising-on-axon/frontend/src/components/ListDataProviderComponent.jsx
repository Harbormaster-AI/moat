import React, { Component } from 'react'
import DataProviderService from '../services/DataProviderService'

class ListDataProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataProviders: []
        }
        this.addDataProvider = this.addDataProvider.bind(this);
        this.editDataProvider = this.editDataProvider.bind(this);
        this.deleteDataProvider = this.deleteDataProvider.bind(this);
    }

    deleteDataProvider(id){
        DataProviderService.deleteDataProvider(id).then( res => {
            this.setState({dataProviders: this.state.dataProviders.filter(dataProvider => dataProvider.dataProviderId !== id)});
        });
    }
    viewDataProvider(id){
        this.props.history.push(`/view-dataProvider/${id}`);
    }
    editDataProvider(id){
        this.props.history.push(`/add-dataProvider/${id}`);
    }

    componentDidMount(){
        DataProviderService.getDataProviders().then((res) => {
            this.setState({ dataProviders: res.data});
        });
    }

    addDataProvider(){
        this.props.history.push('/add-dataProvider/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataProvider List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataProvider}> Add DataProvider</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Website </th>
                                    <th> ProviderType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataProviders.map(
                                        dataProvider => 
                                        <tr key = {dataProvider.dataProviderId}>
                                             <td> { dataProvider.name } </td>
                                             <td> { dataProvider.website } </td>
                                             <td> { dataProvider.providerType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataProvider(dataProvider.dataProviderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataProvider(dataProvider.dataProviderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataProvider(dataProvider.dataProviderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataProviderComponent
