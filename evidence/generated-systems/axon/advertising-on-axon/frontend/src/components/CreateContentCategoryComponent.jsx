import React, { Component } from 'react'
import ContentCategoryService from '../services/ContentCategoryService';

class CreateContentCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                name: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ContentCategoryService.getContentCategoryById(this.state.id).then( (res) =>{
                let contentCategory = res.data;
                this.setState({
                    code: contentCategory.code,
                    name: contentCategory.name
                });
            });
        }        
    }
    saveOrUpdateContentCategory = (e) => {
        e.preventDefault();
        let contentCategory = {
                contentCategoryId: this.state.id,
                code: this.state.code,
                name: this.state.name
            };
        console.log('contentCategory => ' + JSON.stringify(contentCategory));

        // step 5
        if(this.state.id === '_add'){
            contentCategory.contentCategoryId=''
            ContentCategoryService.createContentCategory(contentCategory).then(res =>{
                this.props.history.push('/contentCategorys');
            });
        }else{
            ContentCategoryService.updateContentCategory(contentCategory).then( res => {
                this.props.history.push('/contentCategorys');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/contentCategorys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ContentCategory</h3>
        }else{
            return <h3 className="text-center">Update ContentCategory</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateContentCategory}>Save</button>
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

export default CreateContentCategoryComponent
