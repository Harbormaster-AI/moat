import React, { Component } from 'react'
import TerminationService from '../services/TerminationService';

class CreateTerminationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                terminationNumber: '',
                terminationDate: '',
                notes: '',
                eligibleForRehire: '',
                reason: '',
                type: ''
        }
        this.changeterminationNumberHandler = this.changeterminationNumberHandler.bind(this);
        this.changeterminationDateHandler = this.changeterminationDateHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeeligibleForRehireHandler = this.changeeligibleForRehireHandler.bind(this);
        this.changeReasonHandler = this.changeReasonHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TerminationService.getTerminationById(this.state.id).then( (res) =>{
                let termination = res.data;
                this.setState({
                    terminationNumber: termination.terminationNumber,
                    terminationDate: termination.terminationDate,
                    notes: termination.notes,
                    eligibleForRehire: termination.eligibleForRehire,
                    reason: termination.reason,
                    type: termination.type
                });
            });
        }        
    }
    saveOrUpdateTermination = (e) => {
        e.preventDefault();
        let termination = {
                terminationId: this.state.id,
                terminationNumber: this.state.terminationNumber,
                terminationDate: this.state.terminationDate,
                notes: this.state.notes,
                eligibleForRehire: this.state.eligibleForRehire,
                reason: this.state.reason,
                type: this.state.type
            };
        console.log('termination => ' + JSON.stringify(termination));

        // step 5
        if(this.state.id === '_add'){
            termination.terminationId=''
            TerminationService.createTermination(termination).then(res =>{
                this.props.history.push('/terminations');
            });
        }else{
            TerminationService.updateTermination(termination).then( res => {
                this.props.history.push('/terminations');
            });
        }
    }
    
    changeterminationNumberHandler= (event) => {
        this.setState({terminationNumber: event.target.value});
    }
    changeterminationDateHandler= (event) => {
        this.setState({terminationDate: event.target.value});
    }
    changenotesHandler= (event) => {
        this.setState({notes: event.target.value});
    }
    changeeligibleForRehireHandler= (event) => {
        this.setState({eligibleForRehire: event.target.value});
    }
    changeReasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changeTypeHandler= (event) => {
        this.setState({type: event.target.value});
    }

    cancel(){
        this.props.history.push('/terminations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Termination</h3>
        }else{
            return <h3 className="text-center">Update Termination</h3>
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
                                            <label> terminationNumber:&emsp; </label>
                                                <input placeholder="terminationNumber" name="terminationNumber" className="form-control" value={this.state.terminationNumber} onChange={this.changeterminationNumberHandler}/>

                                            <label> terminationDate:&emsp; </label>
                                                <input type="date" placeholder="terminationDate" name="terminationDate" className="form-control" value={this.state.terminationDate} onChange={this.changeterminationDateHandler}/>

                                            <label> notes:&emsp; </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> eligibleForRehire:&emsp; </label>
                                                <input type="checkbox" placeholder="eligibleForRehire" name="eligibleForRehire" className="form-control" value={this.state.eligibleForRehire} onChange={this.changeeligibleForRehireHandler}/>


                                            <label> Reason:&emsp; </label>
                                                <select value={this.state.reason} onChange={this.changeReasonHandler}>
                      <option name="Reason" className="form-control" >
                          Voluntary
                      </option>
                      <option name="Reason" className="form-control" >
                          Involuntary
                      </option>
                      <option name="Reason" className="form-control" >
                          Retirement
                      </option>
                      <option name="Reason" className="form-control" >
                          Redundancy
                      </option>
                      <option name="Reason" className="form-control" >
                          EndOfContract
                      </option>
                    </select>

                                            <label> Type:&emsp; </label>
                                                <select value={this.state.type} onChange={this.changeTypeHandler}>
                      <option name="Type" className="form-control" >
                          Resignation
                      </option>
                      <option name="Type" className="form-control" >
                          Dismissal
                      </option>
                      <option name="Type" className="form-control" >
                          Layoff
                      </option>
                      <option name="Type" className="form-control" >
                          Retirement
                      </option>
                      <option name="Type" className="form-control" >
                          EndOfAssignment
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTermination}>Save</button>
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

export default CreateTerminationComponent
