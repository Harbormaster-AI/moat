import React, { Component } from 'react'
import TagService from '../services/TagService';

class UpdateTagComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                category: ''
        }
        this.updateTag = this.updateTag.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    componentDidMount(){
        TagService.getTagById(this.state.id).then( (res) =>{
            let tag = res.data;
            this.setState({
                name: tag.name,
                category: tag.category
            });
        });
    }

    updateTag = (e) => {
        e.preventDefault();
        let tag = {
            tagId: this.state.id,
            name: this.state.name,
            category: this.state.category
        };
        console.log('tag => ' + JSON.stringify(tag));
        console.log('id => ' + JSON.stringify(this.state.id));
        TagService.updateTag(tag).then( res => {
            this.props.history.push('/tags');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Tag</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> Category: </label>
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
                                        <button className="btn btn-success" onClick={this.updateTag}>Save</button>
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

export default UpdateTagComponent
