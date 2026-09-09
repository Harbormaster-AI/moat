import React, { Component } from 'react'
import DataSourceService from '../services/DataSourceService';

class CreateDataSourceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                connection: '',
                streaming: '',
                sourceType: '',
                format: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeconnectionHandler = this.changeconnectionHandler.bind(this);
        this.changeStreamingHandler = this.changeStreamingHandler.bind(this);
        this.changeSourceTypeHandler = this.changeSourceTypeHandler.bind(this);
        this.changeFormatHandler = this.changeFormatHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DataSourceService.getDataSourceById(this.state.id).then( (res) =>{
                let dataSource = res.data;
                this.setState({
                    name: dataSource.name,
                    connection: dataSource.connection,
                    streaming: dataSource.streaming,
                    sourceType: dataSource.sourceType,
                    format: dataSource.format
                });
            });
        }        
    }
    saveOrUpdateDataSource = (e) => {
        e.preventDefault();
        let dataSource = {
                dataSourceId: this.state.id,
                name: this.state.name,
                connection: this.state.connection,
                streaming: this.state.streaming,
                sourceType: this.state.sourceType,
                format: this.state.format
            };
        console.log('dataSource => ' + JSON.stringify(dataSource));

        // step 5
        if(this.state.id === '_add'){
            dataSource.dataSourceId=''
            DataSourceService.createDataSource(dataSource).then(res =>{
                this.props.history.push('/dataSources');
            });
        }else{
            DataSourceService.updateDataSource(dataSource).then( res => {
                this.props.history.push('/dataSources');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeconnectionHandler= (event) => {
        this.setState({connection: event.target.value});
    }
    changeStreamingHandler= (event) => {
        this.setState({streaming: event.target.value});
    }
    changeSourceTypeHandler= (event) => {
        this.setState({sourceType: event.target.value});
    }
    changeFormatHandler= (event) => {
        this.setState({format: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataSources');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DataSource</h3>
        }else{
            return <h3 className="text-center">Update DataSource</h3>
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

                                            <label> connection:&emsp; </label>
                                                <input placeholder="connection" name="connection" className="form-control" value={this.state.connection} onChange={this.changeconnectionHandler}/>

                                            <label> Streaming:&emsp; </label>
                                                <input type="checkbox" placeholder="Streaming" name="streaming" className="form-control" value={this.state.streaming} onChange={this.changeStreamingHandler}/>


                                            <label> SourceType:&emsp; </label>
                                                <select value={this.state.sourceType} onChange={this.changeSourceTypeHandler}>
                      <option name="SourceType" className="form-control" >
                          Database
                      </option>
                      <option name="SourceType" className="form-control" >
                          File
                      </option>
                      <option name="SourceType" className="form-control" >
                          Stream
                      </option>
                      <option name="SourceType" className="form-control" >
                          API
                      </option>
                      <option name="SourceType" className="form-control" >
                          DataWarehouse
                      </option>
                      <option name="SourceType" className="form-control" >
                          DataLake
                      </option>
                    </select>

                                            <label> Format:&emsp; </label>
                                                <select value={this.state.format} onChange={this.changeFormatHandler}>
                      <option name="Format" className="form-control" >
                          CSV
                      </option>
                      <option name="Format" className="form-control" >
                          JSON
                      </option>
                      <option name="Format" className="form-control" >
                          Parquet
                      </option>
                      <option name="Format" className="form-control" >
                          Avro
                      </option>
                      <option name="Format" className="form-control" >
                          ORC
                      </option>
                      <option name="Format" className="form-control" >
                          XML
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDataSource}>Save</button>
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

export default CreateDataSourceComponent
