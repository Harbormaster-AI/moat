import React, { Component } from 'react'
import BrandService from '../services/BrandService';

class UpdateBrandComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                description: '',
                website: ''
        }
        this.updateBrand = this.updateBrand.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    componentDidMount(){
        BrandService.getBrandById(this.state.id).then( (res) =>{
            let brand = res.data;
            this.setState({
                name: brand.name,
                description: brand.description,
                website: brand.website
            });
        });
    }

    updateBrand = (e) => {
        e.preventDefault();
        let brand = {
            brandId: this.state.id,
            name: this.state.name,
            description: this.state.description,
            website: this.state.website
        };
        console.log('brand => ' + JSON.stringify(brand));
        console.log('id => ' + JSON.stringify(this.state.id));
        BrandService.updateBrand(brand).then( res => {
            this.props.history.push('/brands');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/brands');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Brand</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBrand}>Save</button>
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

export default UpdateBrandComponent
