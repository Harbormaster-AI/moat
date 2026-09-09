import React, { Component } from 'react'
import SerialNumberService from '../services/SerialNumberService';

class CreateSerialNumberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                serial: '',
                activationDate: '',
                status: ''
        }
        this.changeserialHandler = this.changeserialHandler.bind(this);
        this.changeactivationDateHandler = this.changeactivationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SerialNumberService.getSerialNumberById(this.state.id).then( (res) =>{
                let serialNumber = res.data;
                this.setState({
                    serial: serialNumber.serial,
                    activationDate: serialNumber.activationDate,
                    status: serialNumber.status
                });
            });
        }        
    }
    saveOrUpdateSerialNumber = (e) => {
        e.preventDefault();
        let serialNumber = {
                serialNumberId: this.state.id,
                serial: this.state.serial,
                activationDate: this.state.activationDate,
                status: this.state.status
            };
        console.log('serialNumber => ' + JSON.stringify(serialNumber));

        // step 5
        if(this.state.id === '_add'){
            serialNumber.serialNumberId=''
            SerialNumberService.createSerialNumber(serialNumber).then(res =>{
                this.props.history.push('/serialNumbers');
            });
        }else{
            SerialNumberService.updateSerialNumber(serialNumber).then( res => {
                this.props.history.push('/serialNumbers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SerialNumber</h3>
        }else{
            return <h3 className="text-center">Update SerialNumber</h3>
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
                                            <label> serial:&emsp; </label>
                                                <input placeholder="serial" name="serial" className="form-control" value={this.state.serial} onChange={this.changeserialHandler}/>

                                            <label> activationDate:&emsp; </label>
                                                <input type="date" placeholder="activationDate" name="activationDate" className="form-control" value={this.state.activationDate} onChange={this.changeactivationDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSerialNumber}>Save</button>
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

export default CreateSerialNumberComponent
