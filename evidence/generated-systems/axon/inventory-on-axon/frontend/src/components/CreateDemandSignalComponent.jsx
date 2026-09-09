import React, { Component } from 'react'
import DemandSignalService from '../services/DemandSignalService';

class CreateDemandSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                externalReference: '',
                requestedDate: '',
                quantity: '',
                demandType: ''
        }
        this.changeexternalReferenceHandler = this.changeexternalReferenceHandler.bind(this);
        this.changerequestedDateHandler = this.changerequestedDateHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeDemandTypeHandler = this.changeDemandTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateDemandSignal = (e) => {
        e.preventDefault();
        let demandSignal = {
                demandSignalId: this.state.id,
                externalReference: this.state.externalReference,
                requestedDate: this.state.requestedDate,
                quantity: this.state.quantity,
                demandType: this.state.demandType
            };
        console.log('demandSignal => ' + JSON.stringify(demandSignal));

        // step 5
        if(this.state.id === '_add'){
            demandSignal.demandSignalId=''
            DemandSignalService.createDemandSignal(demandSignal).then(res =>{
                this.props.history.push('/demandSignals');
            });
        }else{
            DemandSignalService.updateDemandSignal(demandSignal).then( res => {
                this.props.history.push('/demandSignals');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DemandSignal</h3>
        }else{
            return <h3 className="text-center">Update DemandSignal</h3>
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
                                            <label> externalReference:&emsp; </label>
                                                <input placeholder="externalReference" name="externalReference" className="form-control" value={this.state.externalReference} onChange={this.changeexternalReferenceHandler}/>

                                            <label> requestedDate:&emsp; </label>
                                                <input type="date" placeholder="requestedDate" name="requestedDate" className="form-control" value={this.state.requestedDate} onChange={this.changerequestedDateHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> DemandType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDemandSignal}>Save</button>
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

export default CreateDemandSignalComponent
