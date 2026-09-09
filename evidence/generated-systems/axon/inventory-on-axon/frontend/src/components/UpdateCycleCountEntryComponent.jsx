import React, { Component } from 'react'
import CycleCountEntryService from '../services/CycleCountEntryService';

class UpdateCycleCountEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                lineNumber: '',
                systemQuantity: '',
                countedQuantity: '',
                varianceQuantity: '',
                recountRequired: '',
                stockStatus: ''
        }
        this.updateCycleCountEntry = this.updateCycleCountEntry.bind(this);

        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changesystemQuantityHandler = this.changesystemQuantityHandler.bind(this);
        this.changecountedQuantityHandler = this.changecountedQuantityHandler.bind(this);
        this.changevarianceQuantityHandler = this.changevarianceQuantityHandler.bind(this);
        this.changerecountRequiredHandler = this.changerecountRequiredHandler.bind(this);
        this.changeStockStatusHandler = this.changeStockStatusHandler.bind(this);
    }

    componentDidMount(){
        CycleCountEntryService.getCycleCountEntryById(this.state.id).then( (res) =>{
            let cycleCountEntry = res.data;
            this.setState({
                lineNumber: cycleCountEntry.lineNumber,
                systemQuantity: cycleCountEntry.systemQuantity,
                countedQuantity: cycleCountEntry.countedQuantity,
                varianceQuantity: cycleCountEntry.varianceQuantity,
                recountRequired: cycleCountEntry.recountRequired,
                stockStatus: cycleCountEntry.stockStatus
            });
        });
    }

    updateCycleCountEntry = (e) => {
        e.preventDefault();
        let cycleCountEntry = {
            cycleCountEntryId: this.state.id,
            lineNumber: this.state.lineNumber,
            systemQuantity: this.state.systemQuantity,
            countedQuantity: this.state.countedQuantity,
            varianceQuantity: this.state.varianceQuantity,
            recountRequired: this.state.recountRequired,
            stockStatus: this.state.stockStatus
        };
        console.log('cycleCountEntry => ' + JSON.stringify(cycleCountEntry));
        console.log('id => ' + JSON.stringify(this.state.id));
        CycleCountEntryService.updateCycleCountEntry(cycleCountEntry).then( res => {
            this.props.history.push('/cycleCountEntrys');
        });
    }

    changelineNumberHandler= (event) => {
        this.setState({lineNumber: event.target.value});
    }
    changesystemQuantityHandler= (event) => {
        this.setState({systemQuantity: event.target.value});
    }
    changecountedQuantityHandler= (event) => {
        this.setState({countedQuantity: event.target.value});
    }
    changevarianceQuantityHandler= (event) => {
        this.setState({varianceQuantity: event.target.value});
    }
    changerecountRequiredHandler= (event) => {
        this.setState({recountRequired: event.target.value});
    }
    changeStockStatusHandler= (event) => {
        this.setState({stockStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/cycleCountEntrys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CycleCountEntry</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber: </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> systemQuantity: </label>
                                                <input placeholder="systemQuantity" name="systemQuantity" className="form-control" value={this.state.systemQuantity} onChange={this.changesystemQuantityHandler}/>

                                            <label> countedQuantity: </label>
                                                <input placeholder="countedQuantity" name="countedQuantity" className="form-control" value={this.state.countedQuantity} onChange={this.changecountedQuantityHandler}/>

                                            <label> varianceQuantity: </label>
                                                <input placeholder="varianceQuantity" name="varianceQuantity" className="form-control" value={this.state.varianceQuantity} onChange={this.changevarianceQuantityHandler}/>

                                            <label> recountRequired: </label>
                                                <input type="checkbox" placeholder="recountRequired" name="recountRequired" className="form-control" value={this.state.recountRequired} onChange={this.changerecountRequiredHandler}/>


                                            <label> StockStatus: </label>
                                                <select value={this.state.stockStatus} onChange={this.changeStockStatusHandler}>
                      <option name="StockStatus" className="form-control" >
                          Available
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Reserved
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Damaged
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Hold
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Quarantined
                      </option>
                      <option name="StockStatus" className="form-control" >
                          InTransit
                      </option>
                      <option name="StockStatus" className="form-control" >
                          PendingInspection
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCycleCountEntry}>Save</button>
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

export default UpdateCycleCountEntryComponent
