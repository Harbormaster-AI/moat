import React, { Component } from 'react'
import BrandService from '../services/BrandService';

class CreateBrandComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                description: '',
                website: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BrandService.getBrandById(this.state.id).then( (res) =>{
                let brand = res.data;
                this.setState({
                    name: brand.name,
                    description: brand.description,
                    website: brand.website
                });
            });
        }        
    }
    saveOrUpdateBrand = (e) => {
        e.preventDefault();
        let brand = {
                brandId: this.state.id,
                name: this.state.name,
                description: this.state.description,
                website: this.state.website
            };
        console.log('brand => ' + JSON.stringify(brand));

        // step 5
        if(this.state.id === '_add'){
            brand.brandId=''
            BrandService.createBrand(brand).then(res =>{
                this.props.history.push('/brands');
            });
        }else{
            BrandService.updateBrand(brand).then( res => {
                this.props.history.push('/brands');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Brand</h3>
        }else{
            return <h3 className="text-center">Update Brand</h3>
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

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBrand}>Save</button>
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

export default CreateBrandComponent
