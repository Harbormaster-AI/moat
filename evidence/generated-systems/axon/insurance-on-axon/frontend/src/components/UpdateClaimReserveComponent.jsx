import React, { Component } from 'react'
import ClaimReserveService from '../services/ClaimReserveService';

class UpdateClaimReserveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                amount: '',
                setDate: '',
                reserveType: '',
                status: ''
        }
        this.updateClaimReserve = this.updateClaimReserve.bind(this);

        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changesetDateHandler = this.changesetDateHandler.bind(this);
        this.changeReserveTypeHandler = this.changeReserveTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ClaimReserveService.getClaimReserveById(this.state.id).then( (res) =>{
            let claimReserve = res.data;
            this.setState({
                amount: claimReserve.amount,
                setDate: claimReserve.setDate,
                reserveType: claimReserve.reserveType,
                status: claimReserve.status
            });
        });
    }

    updateClaimReserve = (e) => {
        e.preventDefault();
        let claimReserve = {
            claimReserveId: this.state.id,
            amount: this.state.amount,
            setDate: this.state.setDate,
            reserveType: this.state.reserveType,
            status: this.state.status
        };
        console.log('claimReserve => ' + JSON.stringify(claimReserve));
        console.log('id => ' + JSON.stringify(this.state.id));
        ClaimReserveService.updateClaimReserve(claimReserve).then( res => {
            this.props.history.push('/claimReserves');
        });
    }

    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changesetDateHandler= (event) => {
        this.setState({setDate: event.target.value});
    }
    changeReserveTypeHandler= (event) => {
        this.setState({reserveType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/claimReserves');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ClaimReserve</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> setDate: </label>
                                                <input type="date" placeholder="setDate" name="setDate" className="form-control" value={this.state.setDate} onChange={this.changesetDateHandler}/>

                                            <label> ReserveType: </label>
                                                <select value={this.state.reserveType} onChange={this.changeReserveTypeHandler}>
                      <option name="ReserveType" className="form-control" >
                          Indemnity
                      </option>
                      <option name="ReserveType" className="form-control" >
                          Expense
                      </option>
                      <option name="ReserveType" className="form-control" >
                          Legal
                      </option>
                      <option name="ReserveType" className="form-control" >
                          Medical
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          Increased
                      </option>
                      <option name="Status" className="form-control" >
                          Decreased
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateClaimReserve}>Save</button>
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

export default UpdateClaimReserveComponent
