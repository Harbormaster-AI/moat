import React, { Component } from 'react'
import CatalogService from '../services/CatalogService'

class ListCatalogComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                catalogs: []
        }
        this.addCatalog = this.addCatalog.bind(this);
        this.editCatalog = this.editCatalog.bind(this);
        this.deleteCatalog = this.deleteCatalog.bind(this);
    }

    deleteCatalog(id){
        CatalogService.deleteCatalog(id).then( res => {
            this.setState({catalogs: this.state.catalogs.filter(catalog => catalog.catalogId !== id)});
        });
    }
    viewCatalog(id){
        this.props.history.push(`/view-catalog/${id}`);
    }
    editCatalog(id){
        this.props.history.push(`/add-catalog/${id}`);
    }

    componentDidMount(){
        CatalogService.getCatalogs().then((res) => {
            this.setState({ catalogs: res.data});
        });
    }

    addCatalog(){
        this.props.history.push('/add-catalog/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Catalog List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCatalog}> Add Catalog</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> CatalogCode </th>
                                    <th> AsActive </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.catalogs.map(
                                        catalog => 
                                        <tr key = {catalog.catalogId}>
                                             <td> { catalog.name } </td>
                                             <td> { catalog.catalogCode } </td>
                                             <td> { catalog.asActive } </td>
                                             <td>
                                                 <button onClick={ () => this.editCatalog(catalog.catalogId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCatalog(catalog.catalogId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCatalog(catalog.catalogId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCatalogComponent
