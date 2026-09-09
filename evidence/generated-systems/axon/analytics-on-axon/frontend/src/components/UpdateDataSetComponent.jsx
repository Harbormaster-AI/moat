import React, { Component } from 'react'
import DataSetService from '../services/DataSetService';

class UpdateDataSetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                schemaVersion: '',
                refreshSchedule: '',
                sensitive: '',
                dataFormat: ''
        }
        this.updateDataSet = this.updateDataSet.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeschemaVersionHandler = this.changeschemaVersionHandler.bind(this);
        this.changerefreshScheduleHandler = this.changerefreshScheduleHandler.bind(this);
        this.changeSensitiveHandler = this.changeSensitiveHandler.bind(this);
        this.changeDataFormatHandler = this.changeDataFormatHandler.bind(this);
    }

    componentDidMount(){
        DataSetService.getDataSetById(this.state.id).then( (res) =>{
            let dataSet = res.data;
            this.setState({
                name: dataSet.name,
                schemaVersion: dataSet.schemaVersion,
                refreshSchedule: dataSet.refreshSchedule,
                sensitive: dataSet.sensitive,
                dataFormat: dataSet.dataFormat
            });
        });
    }

    updateDataSet = (e) => {
        e.preventDefault();
        let dataSet = {
            dataSetId: this.state.id,
            name: this.state.name,
            schemaVersion: this.state.schemaVersion,
            refreshSchedule: this.state.refreshSchedule,
            sensitive: this.state.sensitive,
            dataFormat: this.state.dataFormat
        };
        console.log('dataSet => ' + JSON.stringify(dataSet));
        console.log('id => ' + JSON.stringify(this.state.id));
        DataSetService.updateDataSet(dataSet).then( res => {
            this.props.history.push('/dataSets');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeschemaVersionHandler= (event) => {
        this.setState({schemaVersion: event.target.value});
    }
    changerefreshScheduleHandler= (event) => {
        this.setState({refreshSchedule: event.target.value});
    }
    changeSensitiveHandler= (event) => {
        this.setState({sensitive: event.target.value});
    }
    changeDataFormatHandler= (event) => {
        this.setState({dataFormat: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataSets');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DataSet</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> schemaVersion: </label>
                                                <input placeholder="schemaVersion" name="schemaVersion" className="form-control" value={this.state.schemaVersion} onChange={this.changeschemaVersionHandler}/>

                                            <label> refreshSchedule: </label>
                                                <input placeholder="refreshSchedule" name="refreshSchedule" className="form-control" value={this.state.refreshSchedule} onChange={this.changerefreshScheduleHandler}/>

                                            <label> Sensitive: </label>
                                                <input type="checkbox" placeholder="Sensitive" name="sensitive" className="form-control" value={this.state.sensitive} onChange={this.changeSensitiveHandler}/>


                                            <label> DataFormat: </label>
                                                <select value={this.state.dataFormat} onChange={this.changeDataFormatHandler}>
                      <option name="DataFormat" className="form-control" >
                          CSV
                      </option>
                      <option name="DataFormat" className="form-control" >
                          JSON
                      </option>
                      <option name="DataFormat" className="form-control" >
                          Parquet
                      </option>
                      <option name="DataFormat" className="form-control" >
                          Avro
                      </option>
                      <option name="DataFormat" className="form-control" >
                          ORC
                      </option>
                      <option name="DataFormat" className="form-control" >
                          XML
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDataSet}>Save</button>
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

export default UpdateDataSetComponent
