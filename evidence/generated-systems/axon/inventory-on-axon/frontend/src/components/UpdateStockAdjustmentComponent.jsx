import React, { Component } from 'react'
import StockAdjustmentService from '../services/StockAdjustmentService';

class UpdateStockAdjustmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                adjustmentNumber: '',
                reason: '',
                adjustmentDate: '',
                adjustmentType: '',
                status: ''
        }
        this.updateStockAdjustment = this.updateStockAdjustment.bind(this);

        this.changeadjustmentNumberHandler = this.changeadjustmentNumberHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changeadjustmentDateHandler = this.changeadjustmentDateHandler.bind(this);
        this.changeAdjustmentTypeHandler = this.changeAdjustmentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        StockAdjustmentService.getStockAdjustmentById(this.state.id).then( (res) =>{
            let stockAdjustment = res.data;
            this.setState({
                adjustmentNumber: stockAdjustment.adjustmentNumber,
                reason: stockAdjustment.reason,
                adjustmentDate: stockAdjustment.adjustmentDate,
                adjustmentType: stockAdjustment.adjustmentType,
                status: stockAdjustment.status
            });
        });
    }

    updateStockAdjustment = (e) => {
        e.preventDefault();
        let stockAdjustment = {
            stockAdjustmentId: this.state.id,
            adjustmentNumber: this.state.adjustmentNumber,
            reason: this.state.reason,
            adjustmentDate: this.state.adjustmentDate,
            adjustmentType: this.state.adjustmentType,
            status: this.state.status
        };
        console.log('stockAdjustment => ' + JSON.stringify(stockAdjustment));
        console.log('id => ' + JSON.stringify(this.state.id));
        StockAdjustmentService.updateStockAdjustment(stockAdjustment).then( res => {
            this.props.history.push('/stockAdjustments');
        });
    }

    changeadjustmentNumberHandler= (event) => {
        this.setState({adjustmentNumber: event.target.value});
    }
    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changeadjustmentDateHandler= (event) => {
        this.setState({adjustmentDate: event.target.value});
    }
    changeAdjustmentTypeHandler= (event) => {
        this.setState({adjustmentType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/stockAdjustments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update StockAdjustment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> adjustmentNumber: </label>
                                                <input placeholder="adjustmentNumber" name="adjustmentNumber" className="form-control" value={this.state.adjustmentNumber} onChange={this.changeadjustmentNumberHandler}/>

                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> adjustmentDate: </label>
                                                <input type="date" placeholder="adjustmentDate" name="adjustmentDate" className="form-control" value={this.state.adjustmentDate} onChange={this.changeadjustmentDateHandler}/>

                                            <label> AdjustmentType: </label>
                                                <select value={this.state.adjustmentType} onChange={this.changeAdjustmentTypeHandler}>
                      <option name="AdjustmentType" className="form-control" >
                          Increase
                      </option>
                      <option name="AdjustmentType" className="form-control" >
                          Decrease
                      </option>
                      <option name="AdjustmentType" className="form-control" >
                          Reclassification
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Posted
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateStockAdjustment}>Save</button>
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

export default UpdateStockAdjustmentComponent
