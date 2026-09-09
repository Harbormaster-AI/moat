import React, { Component } from 'react'
import CatalogService from '../services/CatalogService';

class UpdateCatalogComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                catalogCode: '',
                asActive: ''
        }
        this.updateCatalog = this.updateCatalog.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecatalogCodeHandler = this.changecatalogCodeHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
    }

    componentDidMount(){
        CatalogService.getCatalogById(this.state.id).then( (res) =>{
            let catalog = res.data;
            this.setState({
                name: catalog.name,
                catalogCode: catalog.catalogCode,
                asActive: catalog.asActive
            });
        });
    }

    updateCatalog = (e) => {
        e.preventDefault();
        let catalog = {
            catalogId: this.state.id,
            name: this.state.name,
            catalogCode: this.state.catalogCode,
            asActive: this.state.asActive
        };
        console.log('catalog => ' + JSON.stringify(catalog));
        console.log('id => ' + JSON.stringify(this.state.id));
        CatalogService.updateCatalog(catalog).then( res => {
            this.props.history.push('/catalogs');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecatalogCodeHandler= (event) => {
        this.setState({catalogCode: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }

    cancel(){
        this.props.history.push('/catalogs');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Catalog</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> catalogCode: </label>
                                                <input placeholder="catalogCode" name="catalogCode" className="form-control" value={this.state.catalogCode} onChange={this.changecatalogCodeHandler}/>

                                            <label> asActive: </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCatalog}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateCatalogComponent
