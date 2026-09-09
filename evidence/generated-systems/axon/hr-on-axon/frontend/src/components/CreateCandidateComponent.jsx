import React, { Component } from 'react'
import CandidateService from '../services/CandidateService';

class CreateCandidateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                email: '',
                phone: '',
                source: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changeSourceHandler = this.changeSourceHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCandidate = (e) => {
        e.preventDefault();
        let candidate = {
                candidateId: this.state.id,
                name: this.state.name,
                email: this.state.email,
                phone: this.state.phone,
                source: this.state.source
            };
        console.log('candidate => ' + JSON.stringify(candidate));

        // step 5
        if(this.state.id === '_add'){
            candidate.candidateId=''
            CandidateService.createCandidate(candidate).then(res =>{
                this.props.history.push('/candidates');
            });
        }else{
            CandidateService.updateCandidate(candidate).then( res => {
                this.props.history.push('/candidates');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Candidate</h3>
        }else{
            return <h3 className="text-center">Update Candidate</h3>
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

                                            <label> email:&emsp; </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone:&emsp; </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> Source:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCandidate}>Save</button>
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

export default CreateCandidateComponent
