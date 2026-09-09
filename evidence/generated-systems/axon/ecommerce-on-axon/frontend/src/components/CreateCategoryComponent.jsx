import React, { Component } from 'react'
import CategoryService from '../services/CategoryService';

class CreateCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                slug: '',
                position: '',
                asActive: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeslugHandler = this.changeslugHandler.bind(this);
        this.changepositionHandler = this.changepositionHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CategoryService.getCategoryById(this.state.id).then( (res) =>{
                let category = res.data;
                this.setState({
                    name: category.name,
                    slug: category.slug,
                    position: category.position,
                    asActive: category.asActive
                });
            });
        }        
    }
    saveOrUpdateCategory = (e) => {
        e.preventDefault();
        let category = {
                categoryId: this.state.id,
                name: this.state.name,
                slug: this.state.slug,
                position: this.state.position,
                asActive: this.state.asActive
            };
        console.log('category => ' + JSON.stringify(category));

        // step 5
        if(this.state.id === '_add'){
            category.categoryId=''
            CategoryService.createCategory(category).then(res =>{
                this.props.history.push('/categorys');
            });
        }else{
            CategoryService.updateCategory(category).then( res => {
                this.props.history.push('/categorys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeslugHandler= (event) => {
        this.setState({slug: event.target.value});
    }
    changepositionHandler= (event) => {
        this.setState({position: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }

    cancel(){
        this.props.history.push('/categorys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Category</h3>
        }else{
            return <h3 className="text-center">Update Category</h3>
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

                                            <label> slug:&emsp; </label>
                                                <input placeholder="slug" name="slug" className="form-control" value={this.state.slug} onChange={this.changeslugHandler}/>

                                            <label> position:&emsp; </label>
                                                <input type="number" placeholder="position" name="position" className="form-control" value={this.state.position} onChange={this.changepositionHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCategory}>Save</button>
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

export default CreateCategoryComponent
