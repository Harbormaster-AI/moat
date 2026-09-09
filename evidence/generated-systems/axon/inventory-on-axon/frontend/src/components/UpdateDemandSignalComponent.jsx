import React, { Component } from 'react'
import DemandSignalService from '../services/DemandSignalService';

class UpdateDemandSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                externalReference: '',
                requestedDate: '',
                quantity: '',
                demandType: ''
        }
        this.updateDemandSignal = this.updateDemandSignal.bind(this);

        this.changeexternalReferenceHandler = this.changeexternalReferenceHandler.bind(this);
        this.changerequestedDateHandler = this.changerequestedDateHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeDemandTypeHandler = this.changeDemandTypeHandler.bind(this);
    }

    componentDidMount(){
        DemandSignalService.getDemandSignalById(this.state.id).then( (res) =>{
            let demandSignal = res.data;
            this.setState({
                externalReference: demandSignal.externalReference,
                requestedDate: demandSignal.requestedDate,
                quantity: demandSignal.quantity,
                demandType: demandSignal.demandType
            });
        });
    }

    updateDemandSignal = (e) => {
        e.preventDefault();
        let demandSignal = {
            demandSignalId: this.state.id,
            externalReference: this.state.externalReference,
            requestedDate: this.state.requestedDate,
            quantity: this.state.quantity,
            demandType: this.state.demandType
        };
        console.log('demandSignal => ' + JSON.stringify(demandSignal));
        console.log('id => ' + JSON.stringify(this.state.id));
        DemandSignalService.updateDemandSignal(demandSignal).then( res => {
            this.props.history.push('/demandSignals');
        });
    }

    changeexternalReferenceHandler= (event) => {
        this.setState({externalReference: event.target.value});
    }
    changerequestedDateHandler= (event) => {
        this.setState({requestedDate: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeDemandTypeHandler= (event) => {
        this.setState({demandType: event.target.value});
    }

    cancel(){
        this.props.history.push('/demandSignals');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DemandSignal</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> externalReference: </label>
                                                <input placeholder="externalReference" name="externalReference" className="form-control" value={this.state.externalReference} onChange={this.changeexternalReferenceHandler}/>

                                            <label> requestedDate: </label>
                                                <input type="date" placeholder="requestedDate" name="requestedDate" className="form-control" value={this.state.requestedDate} onChange={this.changerequestedDateHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> DemandType: </label>
                                                <select value={this.state.demandType} onChange={this.changeDemandTypeHandler}>
                      <option name="DemandType" className="form-control" >
                          SalesOrder
                      </option>
                      <option name="DemandType" className="form-control" >
                          WorkOrder
                      </option>
                      <option name="DemandType" className="form-control" >
                          TransferOrder
                      </option>
                      <option name="DemandType" className="form-control" >
                          Forecast
                      </option>
                      <option name="DemandType" className="form-control" >
                          SampleRequest
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDemandSignal}>Save</button>
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

export default UpdateDemandSignalComponent
