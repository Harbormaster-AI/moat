import React, { Component } from 'react'
import SubrogationRecoveryService from '../services/SubrogationRecoveryService';

class CreateSubrogationRecoveryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                recoveryReference: '',
                amount: '',
                recoveryDate: '',
                status: ''
        }
        this.changerecoveryReferenceHandler = this.changerecoveryReferenceHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changerecoveryDateHandler = this.changerecoveryDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SubrogationRecoveryService.getSubrogationRecoveryById(this.state.id).then( (res) =>{
                let subrogationRecovery = res.data;
                this.setState({
                    recoveryReference: subrogationRecovery.recoveryReference,
                    amount: subrogationRecovery.amount,
                    recoveryDate: subrogationRecovery.recoveryDate,
                    status: subrogationRecovery.status
                });
            });
        }        
    }
    saveOrUpdateSubrogationRecovery = (e) => {
        e.preventDefault();
        let subrogationRecovery = {
                subrogationRecoveryId: this.state.id,
                recoveryReference: this.state.recoveryReference,
                amount: this.state.amount,
                recoveryDate: this.state.recoveryDate,
                status: this.state.status
            };
        console.log('subrogationRecovery => ' + JSON.stringify(subrogationRecovery));

        // step 5
        if(this.state.id === '_add'){
            subrogationRecovery.subrogationRecoveryId=''
            SubrogationRecoveryService.createSubrogationRecovery(subrogationRecovery).then(res =>{
                this.props.history.push('/subrogationRecoverys');
            });
        }else{
            SubrogationRecoveryService.updateSubrogationRecovery(subrogationRecovery).then( res => {
                this.props.history.push('/subrogationRecoverys');
            });
        }
    }
    
    changerecoveryReferenceHandler= (event) => {
        this.setState({recoveryReference: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changerecoveryDateHandler= (event) => {
        this.setState({recoveryDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/subrogationRecoverys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SubrogationRecovery</h3>
        }else{
            return <h3 className="text-center">Update SubrogationRecovery</h3>
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
                                            <label> recoveryReference:&emsp; </label>
                                                <input placeholder="recoveryReference" name="recoveryReference" className="form-control" value={this.state.recoveryReference} onChange={this.changerecoveryReferenceHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> recoveryDate:&emsp; </label>
                                                <input type="date" placeholder="recoveryDate" name="recoveryDate" className="form-control" value={this.state.recoveryDate} onChange={this.changerecoveryDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Negotiating
                      </option>
                      <option name="Status" className="form-control" >
                          Settled
                      </option>
                      <option name="Status" className="form-control" >
                          Uncollectible
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSubrogationRecovery}>Save</button>
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

export default CreateSubrogationRecoveryComponent
