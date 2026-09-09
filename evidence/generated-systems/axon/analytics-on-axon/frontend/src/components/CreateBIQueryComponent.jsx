import React, { Component } from 'react'
import BIQueryService from '../services/BIQueryService';

class CreateBIQueryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                text: '',
                dialect: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetextHandler = this.changetextHandler.bind(this);
        this.changeDialectHandler = this.changeDialectHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BIQueryService.getBIQueryById(this.state.id).then( (res) =>{
                let bIQuery = res.data;
                this.setState({
                    name: bIQuery.name,
                    text: bIQuery.text,
                    dialect: bIQuery.dialect
                });
            });
        }        
    }
    saveOrUpdateBIQuery = (e) => {
        e.preventDefault();
        let bIQuery = {
                bIQueryId: this.state.id,
                name: this.state.name,
                text: this.state.text,
                dialect: this.state.dialect
            };
        console.log('bIQuery => ' + JSON.stringify(bIQuery));

        // step 5
        if(this.state.id === '_add'){
            bIQuery.bIQueryId=''
            BIQueryService.createBIQuery(bIQuery).then(res =>{
                this.props.history.push('/bIQuerys');
            });
        }else{
            BIQueryService.updateBIQuery(bIQuery).then( res => {
                this.props.history.push('/bIQuerys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetextHandler= (event) => {
        this.setState({text: event.target.value});
    }
    changeDialectHandler= (event) => {
        this.setState({dialect: event.target.value});
    }

    cancel(){
        this.props.history.push('/bIQuerys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BIQuery</h3>
        }else{
            return <h3 className="text-center">Update BIQuery</h3>
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

                                            <label> text:&emsp; </label>
                                                <input placeholder="text" name="text" className="form-control" value={this.state.text} onChange={this.changetextHandler}/>

                                            <label> Dialect:&emsp; </label>
                                                <select value={this.state.dialect} onChange={this.changeDialectHandler}>
                      <option name="Dialect" className="form-control" >
                          ANSI
                      </option>
                      <option name="Dialect" className="form-control" >
                          Postgres
                      </option>
                      <option name="Dialect" className="form-control" >
                          MySQL
                      </option>
                      <option name="Dialect" className="form-control" >
                          SQLServer
                      </option>
                      <option name="Dialect" className="form-control" >
                          Oracle
                      </option>
                      <option name="Dialect" className="form-control" >
                          SparkSQL
                      </option>
                      <option name="Dialect" className="form-control" >
                          BigQuery
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBIQuery}>Save</button>
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

export default CreateBIQueryComponent
