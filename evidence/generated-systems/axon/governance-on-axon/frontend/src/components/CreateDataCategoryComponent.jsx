import React, { Component } from 'react'
import DataCategoryService from '../services/DataCategoryService';

class CreateDataCategoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                description: '',
                classification: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeClassificationHandler = this.changeClassificationHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DataCategoryService.getDataCategoryById(this.state.id).then( (res) =>{
                let dataCategory = res.data;
                this.setState({
                    name: dataCategory.name,
                    description: dataCategory.description,
                    classification: dataCategory.classification
                });
            });
        }        
    }
    saveOrUpdateDataCategory = (e) => {
        e.preventDefault();
        let dataCategory = {
                dataCategoryId: this.state.id,
                name: this.state.name,
                description: this.state.description,
                classification: this.state.classification
            };
        console.log('dataCategory => ' + JSON.stringify(dataCategory));

        // step 5
        if(this.state.id === '_add'){
            dataCategory.dataCategoryId=''
            DataCategoryService.createDataCategory(dataCategory).then(res =>{
                this.props.history.push('/dataCategorys');
            });
        }else{
            DataCategoryService.updateDataCategory(dataCategory).then( res => {
                this.props.history.push('/dataCategorys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeClassificationHandler= (event) => {
        this.setState({classification: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataCategorys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DataCategory</h3>
        }else{
            return <h3 className="text-center">Update DataCategory</h3>
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

                                            <label> Classification:&emsp; </label>
                                                <select value={this.state.classification} onChange={this.changeClassificationHandler}>
                      <option name="Classification" className="form-control" >
                          Public
                      </option>
                      <option name="Classification" className="form-control" >
                          Internal
                      </option>
                      <option name="Classification" className="form-control" >
                          Confidential
                      </option>
                      <option name="Classification" className="form-control" >
                          Restricted
                      </option>
                      <option name="Classification" className="form-control" >
                          HighlyRestricted
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDataCategory}>Save</button>
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

export default CreateDataCategoryComponent
