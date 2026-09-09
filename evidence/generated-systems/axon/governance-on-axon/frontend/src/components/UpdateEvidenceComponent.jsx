import React, { Component } from 'react'
import EvidenceService from '../services/EvidenceService';

class UpdateEvidenceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                locationUrl: '',
                receivedDate: '',
                evidenceType: ''
        }
        this.updateEvidence = this.updateEvidence.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changelocationUrlHandler = this.changelocationUrlHandler.bind(this);
        this.changereceivedDateHandler = this.changereceivedDateHandler.bind(this);
        this.changeEvidenceTypeHandler = this.changeEvidenceTypeHandler.bind(this);
    }

    componentDidMount(){
        EvidenceService.getEvidenceById(this.state.id).then( (res) =>{
            let evidence = res.data;
            this.setState({
                title: evidence.title,
                locationUrl: evidence.locationUrl,
                receivedDate: evidence.receivedDate,
                evidenceType: evidence.evidenceType
            });
        });
    }

    updateEvidence = (e) => {
        e.preventDefault();
        let evidence = {
            evidenceId: this.state.id,
            title: this.state.title,
            locationUrl: this.state.locationUrl,
            receivedDate: this.state.receivedDate,
            evidenceType: this.state.evidenceType
        };
        console.log('evidence => ' + JSON.stringify(evidence));
        console.log('id => ' + JSON.stringify(this.state.id));
        EvidenceService.updateEvidence(evidence).then( res => {
            this.props.history.push('/evidences');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changelocationUrlHandler= (event) => {
        this.setState({locationUrl: event.target.value});
    }
    changereceivedDateHandler= (event) => {
        this.setState({receivedDate: event.target.value});
    }
    changeEvidenceTypeHandler= (event) => {
        this.setState({evidenceType: event.target.value});
    }

    cancel(){
        this.props.history.push('/evidences');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Evidence</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> locationUrl: </label>
                                                <input placeholder="locationUrl" name="locationUrl" className="form-control" value={this.state.locationUrl} onChange={this.changelocationUrlHandler}/>

                                            <label> receivedDate: </label>
                                                <input type="date" placeholder="receivedDate" name="receivedDate" className="form-control" value={this.state.receivedDate} onChange={this.changereceivedDateHandler}/>

                                            <label> EvidenceType: </label>
                                                <select value={this.state.evidenceType} onChange={this.changeEvidenceTypeHandler}>
                      <option name="EvidenceType" className="form-control" >
                          Document
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          Screenshot
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          LogExport
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          SystemReport
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          Ticket
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          Attestation
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          Configuration
                      </option>
                      <option name="EvidenceType" className="form-control" >
                          Dataset
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEvidence}>Save</button>
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

export default UpdateEvidenceComponent
