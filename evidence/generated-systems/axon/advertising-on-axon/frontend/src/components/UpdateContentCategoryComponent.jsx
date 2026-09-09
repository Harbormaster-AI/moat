import React, { Component } from 'react'
import ContentCategoryService from '../services/ContentCategoryService';

class UpdateContentCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                name: ''
        }
        this.updateContentCategory = this.updateContentCategory.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        ContentCategoryService.getContentCategoryById(this.state.id).then( (res) =>{
            let contentCategory = res.data;
            this.setState({
                code: contentCategory.code,
                name: contentCategory.name
            });
        });
    }

    updateContentCategory = (e) => {
        e.preventDefault();
        let contentCategory = {
            contentCategoryId: this.state.id,
            code: this.state.code,
            name: this.state.name
        };
        console.log('contentCategory => ' + JSON.stringify(contentCategory));
        console.log('id => ' + JSON.stringify(this.state.id));
        ContentCategoryService.updateContentCategory(contentCategory).then( res => {
            this.props.history.push('/contentCategorys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ContentCategory</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateContentCategory}>Save</button>
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

export default UpdateContentCategoryComponent
