import React, { Component } from 'react'
import Record_Service from '../services/Record_Service';

class UpdateRecord_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                creationDate: '',
                recordType: '',
                classification: '',
                status: ''
        }
        this.updateRecord_ = this.updateRecord_.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changecreationDateHandler = this.changecreationDateHandler.bind(this);
        this.changeRecordTypeHandler = this.changeRecordTypeHandler.bind(this);
        this.changeClassificationHandler = this.changeClassificationHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        Record_Service.getRecord_ById(this.state.id).then( (res) =>{
            let record_ = res.data;
            this.setState({
                title: record_.title,
                creationDate: record_.creationDate,
                recordType: record_.recordType,
                classification: record_.classification,
                status: record_.status
            });
        });
    }

    updateRecord_ = (e) => {
        e.preventDefault();
        let record_ = {
            record_Id: this.state.id,
            title: this.state.title,
            creationDate: this.state.creationDate,
            recordType: this.state.recordType,
            classification: this.state.classification,
            status: this.state.status
        };
        console.log('record_ => ' + JSON.stringify(record_));
        console.log('id => ' + JSON.stringify(this.state.id));
        Record_Service.updateRecord_(record_).then( res => {
            this.props.history.push('/record_s');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changecreationDateHandler= (event) => {
        this.setState({creationDate: event.target.value});
    }
    changeRecordTypeHandler= (event) => {
        this.setState({recordType: event.target.value});
    }
    changeClassificationHandler= (event) => {
        this.setState({classification: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/record_s');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Record_</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> creationDate: </label>
                                                <input type="date" placeholder="creationDate" name="creationDate" className="form-control" value={this.state.creationDate} onChange={this.changecreationDateHandler}/>

                                            <label> RecordType: </label>
                                                <select value={this.state.recordType} onChange={this.changeRecordTypeHandler}>
                      <option name="RecordType" className="form-control" >
                          PolicyRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          ContractRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          FinancialRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          HRRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          CustomerRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          TechnicalRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          AuditRecord
                      </option>
                      <option name="RecordType" className="form-control" >
                          LegalRecord
                      </option>
                    </select>

                                            <label> Classification: </label>
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

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Archived
                      </option>
                      <option name="Status" className="form-control" >
                          PendingDisposition
                      </option>
                      <option name="Status" className="form-control" >
                          Disposed
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRecord_}>Save</button>
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

export default UpdateRecord_Component
