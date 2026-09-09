import React, { Component } from 'react'
import CandidateService from '../services/CandidateService';

class UpdateCandidateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                email: '',
                phone: '',
                source: ''
        }
        this.updateCandidate = this.updateCandidate.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changeSourceHandler = this.changeSourceHandler.bind(this);
    }

    componentDidMount(){
        CandidateService.getCandidateById(this.state.id).then( (res) =>{
            let candidate = res.data;
            this.setState({
                name: candidate.name,
                email: candidate.email,
                phone: candidate.phone,
                source: candidate.source
            });
        });
    }

    updateCandidate = (e) => {
        e.preventDefault();
        let candidate = {
            candidateId: this.state.id,
            name: this.state.name,
            email: this.state.email,
            phone: this.state.phone,
            source: this.state.source
        };
        console.log('candidate => ' + JSON.stringify(candidate));
        console.log('id => ' + JSON.stringify(this.state.id));
        CandidateService.updateCandidate(candidate).then( res => {
            this.props.history.push('/candidates');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changephoneHandler= (event) => {
        this.setState({phone: event.target.value});
    }
    changeSourceHandler= (event) => {
        this.setState({source: event.target.value});
    }

    cancel(){
        this.props.history.push('/candidates');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Candidate</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone: </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> Source: </label>
                                                <select value={this.state.source} onChange={this.changeSourceHandler}>
                      <option name="Source" className="form-control" >
                          Referral
                      </option>
                      <option name="Source" className="form-control" >
                          Agency
                      </option>
                      <option name="Source" className="form-control" >
                          JobBoard
                      </option>
                      <option name="Source" className="form-control" >
                          CareerSite
                      </option>
                      <option name="Source" className="form-control" >
                          Campus
                      </option>
                      <option name="Source" className="form-control" >
                          Social
                      </option>
                      <option name="Source" className="form-control" >
                          Internal
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCandidate}>Save</button>
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

export default UpdateCandidateComponent
