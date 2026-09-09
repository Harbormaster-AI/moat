import React, { Component } from 'react'
import CatalogService from '../services/CatalogService';

class CreateCatalogComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                catalogCode: '',
                asActive: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecatalogCodeHandler = this.changecatalogCodeHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CatalogService.getCatalogById(this.state.id).then( (res) =>{
                let catalog = res.data;
                this.setState({
                    name: catalog.name,
                    catalogCode: catalog.catalogCode,
                    asActive: catalog.asActive
                });
            });
        }        
    }
    saveOrUpdateCatalog = (e) => {
        e.preventDefault();
        let catalog = {
                catalogId: this.state.id,
                name: this.state.name,
                catalogCode: this.state.catalogCode,
                asActive: this.state.asActive
            };
        console.log('catalog => ' + JSON.stringify(catalog));

        // step 5
        if(this.state.id === '_add'){
            catalog.catalogId=''
            CatalogService.createCatalog(catalog).then(res =>{
                this.props.history.push('/catalogs');
            });
        }else{
            CatalogService.updateCatalog(catalog).then( res => {
                this.props.history.push('/catalogs');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Catalog</h3>
        }else{
            return <h3 className="text-center">Update Catalog</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> catalogCode:&emsp; </label>
                                                <input placeholder="catalogCode" name="catalogCode" className="form-control" value={this.state.catalogCode} onChange={this.changecatalogCodeHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCatalog}>Save</button>
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

export default CreateCatalogComponent
