import React, { Component } from 'react'
import TagService from '../services/TagService';

class CreateTagComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                category: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TagService.getTagById(this.state.id).then( (res) =>{
                let tag = res.data;
                this.setState({
                    name: tag.name,
                    category: tag.category
                });
            });
        }        
    }
    saveOrUpdateTag = (e) => {
        e.preventDefault();
        let tag = {
                tagId: this.state.id,
                name: this.state.name,
                category: this.state.category
            };
        console.log('tag => ' + JSON.stringify(tag));

        // step 5
        if(this.state.id === '_add'){
            tag.tagId=''
            TagService.createTag(tag).then(res =>{
                this.props.history.push('/tags');
            });
        }else{
            TagService.updateTag(tag).then( res => {
                this.props.history.push('/tags');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/tags');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Tag</h3>
        }else{
            return <h3 className="text-center">Update Tag</h3>
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

                                            <label> Category:&emsp; </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          Domain
                      </option>
                      <option name="Category" className="form-control" >
                          Sensitivity
                      </option>
                      <option name="Category" className="form-control" >
                          Priority
                      </option>
                      <option name="Category" className="form-control" >
                          Lifecycle
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTag}>Save</button>
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

export default CreateTagComponent
