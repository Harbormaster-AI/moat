import React, { Component } from 'react'
import DataCategoryService from '../services/DataCategoryService'

class ListDataCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataCategorys: []
        }
        this.addDataCategory = this.addDataCategory.bind(this);
        this.editDataCategory = this.editDataCategory.bind(this);
        this.deleteDataCategory = this.deleteDataCategory.bind(this);
    }

    deleteDataCategory(id){
        DataCategoryService.deleteDataCategory(id).then( res => {
            this.setState({dataCategorys: this.state.dataCategorys.filter(dataCategory => dataCategory.dataCategoryId !== id)});
        });
    }
    viewDataCategory(id){
        this.props.history.push(`/view-dataCategory/${id}`);
    }
    editDataCategory(id){
        this.props.history.push(`/add-dataCategory/${id}`);
    }

    componentDidMount(){
        DataCategoryService.getDataCategorys().then((res) => {
            this.setState({ dataCategorys: res.data});
        });
    }

    addDataCategory(){
        this.props.history.push('/add-dataCategory/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataCategory List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataCategory}> Add DataCategory</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Description </th>
                                    <th> Classification </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataCategorys.map(
                                        dataCategory => 
                                        <tr key = {dataCategory.dataCategoryId}>
                                             <td> { dataCategory.name } </td>
                                             <td> { dataCategory.description } </td>
                                             <td> { dataCategory.classification } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataCategory(dataCategory.dataCategoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataCategory(dataCategory.dataCategoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataCategory(dataCategory.dataCategoryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataCategoryComponent
