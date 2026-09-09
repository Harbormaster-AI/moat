import React, { Component } from 'react'
import LotService from '../services/LotService';

class UpdateLotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                batchNumber: '',
                manufactureDate: '',
                expirationDate: '',
                lotStatus: ''
        }
        this.updateLot = this.updateLot.bind(this);

        this.changebatchNumberHandler = this.changebatchNumberHandler.bind(this);
        this.changemanufactureDateHandler = this.changemanufactureDateHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeLotStatusHandler = this.changeLotStatusHandler.bind(this);
    }

    componentDidMount(){
        LotService.getLotById(this.state.id).then( (res) =>{
            let lot = res.data;
            this.setState({
                batchNumber: lot.batchNumber,
                manufactureDate: lot.manufactureDate,
                expirationDate: lot.expirationDate,
                lotStatus: lot.lotStatus
            });
        });
    }

    updateLot = (e) => {
        e.preventDefault();
        let lot = {
            lotId: this.state.id,
            batchNumber: this.state.batchNumber,
            manufactureDate: this.state.manufactureDate,
            expirationDate: this.state.expirationDate,
            lotStatus: this.state.lotStatus
        };
        console.log('lot => ' + JSON.stringify(lot));
        console.log('id => ' + JSON.stringify(this.state.id));
        LotService.updateLot(lot).then( res => {
            this.props.history.push('/lots');
        });
    }

    changebatchNumberHandler= (event) => {
        this.setState({batchNumber: event.target.value});
    }
    changemanufactureDateHandler= (event) => {
        this.setState({manufactureDate: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeLotStatusHandler= (event) => {
        this.setState({lotStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/lots');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Lot</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> batchNumber: </label>
                                                <input placeholder="batchNumber" name="batchNumber" className="form-control" value={this.state.batchNumber} onChange={this.changebatchNumberHandler}/>

                                            <label> manufactureDate: </label>
                                                <input type="date" placeholder="manufactureDate" name="manufactureDate" className="form-control" value={this.state.manufactureDate} onChange={this.changemanufactureDateHandler}/>

                                            <label> expirationDate: </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> LotStatus: </label>
                                                <select value={this.state.lotStatus} onChange={this.changeLotStatusHandler}>
                      <option name="LotStatus" className="form-control" >
                          Released
                      </option>
                      <option name="LotStatus" className="form-control" >
                          Quarantined
                      </option>
                      <option name="LotStatus" className="form-control" >
                          Expired
                      </option>
                      <option name="LotStatus" className="form-control" >
                          Blocked
                      </option>
                      <option name="LotStatus" className="form-control" >
                          PendingTest
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLot}>Save</button>
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

export default UpdateLotComponent
