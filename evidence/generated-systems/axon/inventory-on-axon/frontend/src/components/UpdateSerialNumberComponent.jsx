import React, { Component } from 'react'
import SerialNumberService from '../services/SerialNumberService';

class UpdateSerialNumberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                serial: '',
                activationDate: '',
                status: ''
        }
        this.updateSerialNumber = this.updateSerialNumber.bind(this);

        this.changeserialHandler = this.changeserialHandler.bind(this);
        this.changeactivationDateHandler = this.changeactivationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        SerialNumberService.getSerialNumberById(this.state.id).then( (res) =>{
            let serialNumber = res.data;
            this.setState({
                serial: serialNumber.serial,
                activationDate: serialNumber.activationDate,
                status: serialNumber.status
            });
        });
    }

    updateSerialNumber = (e) => {
        e.preventDefault();
        let serialNumber = {
            serialNumberId: this.state.id,
            serial: this.state.serial,
            activationDate: this.state.activationDate,
            status: this.state.status
        };
        console.log('serialNumber => ' + JSON.stringify(serialNumber));
        console.log('id => ' + JSON.stringify(this.state.id));
        SerialNumberService.updateSerialNumber(serialNumber).then( res => {
            this.props.history.push('/serialNumbers');
        });
    }

    changeserialHandler= (event) => {
        this.setState({serial: event.target.value});
    }
    changeactivationDateHandler= (event) => {
        this.setState({activationDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/serialNumbers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SerialNumber</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> serial: </label>
                                                <input placeholder="serial" name="serial" className="form-control" value={this.state.serial} onChange={this.changeserialHandler}/>

                                            <label> activationDate: </label>
                                                <input type="date" placeholder="activationDate" name="activationDate" className="form-control" value={this.state.activationDate} onChange={this.changeactivationDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Assigned
                      </option>
                      <option name="Status" className="form-control" >
                          InTransit
                      </option>
                      <option name="Status" className="form-control" >
                          Consumed
                      </option>
                      <option name="Status" className="form-control" >
                          Returned
                      </option>
                      <option name="Status" className="form-control" >
                          Scrapped
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSerialNumber}>Save</button>
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

export default UpdateSerialNumberComponent
