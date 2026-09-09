import React, { Component } from 'react'
import EvidenceService from '../services/EvidenceService';

class CreateEvidenceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                locationUrl: '',
                receivedDate: '',
                evidenceType: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changelocationUrlHandler = this.changelocationUrlHandler.bind(this);
        this.changereceivedDateHandler = this.changereceivedDateHandler.bind(this);
        this.changeEvidenceTypeHandler = this.changeEvidenceTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateEvidence = (e) => {
        e.preventDefault();
        let evidence = {
                evidenceId: this.state.id,
                title: this.state.title,
                locationUrl: this.state.locationUrl,
                receivedDate: this.state.receivedDate,
                evidenceType: this.state.evidenceType
            };
        console.log('evidence => ' + JSON.stringify(evidence));

        // step 5
        if(this.state.id === '_add'){
            evidence.evidenceId=''
            EvidenceService.createEvidence(evidence).then(res =>{
                this.props.history.push('/evidences');
            });
        }else{
            EvidenceService.updateEvidence(evidence).then( res => {
                this.props.history.push('/evidences');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Evidence</h3>
        }else{
            return <h3 className="text-center">Update Evidence</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> locationUrl:&emsp; </label>
                                                <input placeholder="locationUrl" name="locationUrl" className="form-control" value={this.state.locationUrl} onChange={this.changelocationUrlHandler}/>

                                            <label> receivedDate:&emsp; </label>
                                                <input type="date" placeholder="receivedDate" name="receivedDate" className="form-control" value={this.state.receivedDate} onChange={this.changereceivedDateHandler}/>

                                            <label> EvidenceType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEvidence}>Save</button>
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

export default CreateEvidenceComponent
