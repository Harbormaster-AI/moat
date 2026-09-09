import React, { Component } from 'react'
import BIQueryService from '../services/BIQueryService';

class UpdateBIQueryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                text: '',
                dialect: ''
        }
        this.updateBIQuery = this.updateBIQuery.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetextHandler = this.changetextHandler.bind(this);
        this.changeDialectHandler = this.changeDialectHandler.bind(this);
    }

    componentDidMount(){
        BIQueryService.getBIQueryById(this.state.id).then( (res) =>{
            let bIQuery = res.data;
            this.setState({
                name: bIQuery.name,
                text: bIQuery.text,
                dialect: bIQuery.dialect
            });
        });
    }

    updateBIQuery = (e) => {
        e.preventDefault();
        let bIQuery = {
            bIQueryId: this.state.id,
            name: this.state.name,
            text: this.state.text,
            dialect: this.state.dialect
        };
        console.log('bIQuery => ' + JSON.stringify(bIQuery));
        console.log('id => ' + JSON.stringify(this.state.id));
        BIQueryService.updateBIQuery(bIQuery).then( res => {
            this.props.history.push('/bIQuerys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BIQuery</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> text: </label>
                                                <input placeholder="text" name="text" className="form-control" value={this.state.text} onChange={this.changetextHandler}/>

                                            <label> Dialect: </label>
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
                                        <button className="btn btn-success" onClick={this.updateBIQuery}>Save</button>
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

export default UpdateBIQueryComponent
