import React, { Component } from 'react'
import SettlementBatchService from '../services/SettlementBatchService';

class CreateSettlementBatchComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                batchId: '',
                periodStart: '',
                periodEnd: '',
                totalVolume: '',
                totalCount: '',
                status: ''
        }
        this.changebatchIdHandler = this.changebatchIdHandler.bind(this);
        this.changeperiodStartHandler = this.changeperiodStartHandler.bind(this);
        this.changeperiodEndHandler = this.changeperiodEndHandler.bind(this);
        this.changetotalVolumeHandler = this.changetotalVolumeHandler.bind(this);
        this.changetotalCountHandler = this.changetotalCountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SettlementBatchService.getSettlementBatchById(this.state.id).then( (res) =>{
                let settlementBatch = res.data;
                this.setState({
                    batchId: settlementBatch.batchId,
                    periodStart: settlementBatch.periodStart,
                    periodEnd: settlementBatch.periodEnd,
                    totalVolume: settlementBatch.totalVolume,
                    totalCount: settlementBatch.totalCount,
                    status: settlementBatch.status
                });
            });
        }        
    }
    saveOrUpdateSettlementBatch = (e) => {
        e.preventDefault();
        let settlementBatch = {
                settlementBatchId: this.state.id,
                batchId: this.state.batchId,
                periodStart: this.state.periodStart,
                periodEnd: this.state.periodEnd,
                totalVolume: this.state.totalVolume,
                totalCount: this.state.totalCount,
                status: this.state.status
            };
        console.log('settlementBatch => ' + JSON.stringify(settlementBatch));

        // step 5
        if(this.state.id === '_add'){
            settlementBatch.settlementBatchId=''
            SettlementBatchService.createSettlementBatch(settlementBatch).then(res =>{
                this.props.history.push('/settlementBatchs');
            });
        }else{
            SettlementBatchService.updateSettlementBatch(settlementBatch).then( res => {
                this.props.history.push('/settlementBatchs');
            });
        }
    }
    
    changebatchIdHandler= (event) => {
        this.setState({batchId: event.target.value});
    }
    changeperiodStartHandler= (event) => {
        this.setState({periodStart: event.target.value});
    }
    changeperiodEndHandler= (event) => {
        this.setState({periodEnd: event.target.value});
    }
    changetotalVolumeHandler= (event) => {
        this.setState({totalVolume: event.target.value});
    }
    changetotalCountHandler= (event) => {
        this.setState({totalCount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/settlementBatchs');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SettlementBatch</h3>
        }else{
            return <h3 className="text-center">Update SettlementBatch</h3>
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
                                            <label> batchId:&emsp; </label>
                                                <input placeholder="batchId" name="batchId" className="form-control" value={this.state.batchId} onChange={this.changebatchIdHandler}/>

                                            <label> periodStart:&emsp; </label>
                                                <input type="time" placeholder="periodStart" name="periodStart" className="form-control" value={this.state.periodStart} onChange={this.changeperiodStartHandler}/>

                                            <label> periodEnd:&emsp; </label>
                                                <input type="time" placeholder="periodEnd" name="periodEnd" className="form-control" value={this.state.periodEnd} onChange={this.changeperiodEndHandler}/>

                                            <label> totalVolume:&emsp; </label>
                                                <input placeholder="totalVolume" name="totalVolume" className="form-control" value={this.state.totalVolume} onChange={this.changetotalVolumeHandler}/>

                                            <label> totalCount:&emsp; </label>
                                                <input type="number" placeholder="totalCount" name="totalCount" className="form-control" value={this.state.totalCount} onChange={this.changetotalCountHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Processing
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Reconciled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSettlementBatch}>Save</button>
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

export default CreateSettlementBatchComponent
