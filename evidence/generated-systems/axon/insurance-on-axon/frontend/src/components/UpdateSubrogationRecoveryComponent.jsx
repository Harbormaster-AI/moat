import React, { Component } from 'react'
import SubrogationRecoveryService from '../services/SubrogationRecoveryService';

class UpdateSubrogationRecoveryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                recoveryReference: '',
                amount: '',
                recoveryDate: '',
                status: ''
        }
        this.updateSubrogationRecovery = this.updateSubrogationRecovery.bind(this);

        this.changerecoveryReferenceHandler = this.changerecoveryReferenceHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changerecoveryDateHandler = this.changerecoveryDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateSubrogationRecovery = (e) => {
        e.preventDefault();
        let subrogationRecovery = {
            subrogationRecoveryId: this.state.id,
            recoveryReference: this.state.recoveryReference,
            amount: this.state.amount,
            recoveryDate: this.state.recoveryDate,
            status: this.state.status
        };
        console.log('subrogationRecovery => ' + JSON.stringify(subrogationRecovery));
        console.log('id => ' + JSON.stringify(this.state.id));
        SubrogationRecoveryService.updateSubrogationRecovery(subrogationRecovery).then( res => {
            this.props.history.push('/subrogationRecoverys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SubrogationRecovery</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> recoveryReference: </label>
                                                <input placeholder="recoveryReference" name="recoveryReference" className="form-control" value={this.state.recoveryReference} onChange={this.changerecoveryReferenceHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> recoveryDate: </label>
                                                <input type="date" placeholder="recoveryDate" name="recoveryDate" className="form-control" value={this.state.recoveryDate} onChange={this.changerecoveryDateHandler}/>

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateSubrogationRecovery}>Save</button>
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

export default UpdateSubrogationRecoveryComponent
