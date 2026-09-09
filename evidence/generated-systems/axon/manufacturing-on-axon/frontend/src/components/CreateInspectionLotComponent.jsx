import React, { Component } from 'react'
import InspectionLotService from '../services/InspectionLotService';

class CreateInspectionLotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                lotNumber: '',
                quantity: '',
                sampleSize: '',
                createdOn: '',
                inspectionType: '',
                status: ''
        }
        this.changelotNumberHandler = this.changelotNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changesampleSizeHandler = this.changesampleSizeHandler.bind(this);
        this.changecreatedOnHandler = this.changecreatedOnHandler.bind(this);
        this.changeInspectionTypeHandler = this.changeInspectionTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InspectionLotService.getInspectionLotById(this.state.id).then( (res) =>{
                let inspectionLot = res.data;
                this.setState({
                    lotNumber: inspectionLot.lotNumber,
                    quantity: inspectionLot.quantity,
                    sampleSize: inspectionLot.sampleSize,
                    createdOn: inspectionLot.createdOn,
                    inspectionType: inspectionLot.inspectionType,
                    status: inspectionLot.status
                });
            });
        }        
    }
    saveOrUpdateInspectionLot = (e) => {
        e.preventDefault();
        let inspectionLot = {
                inspectionLotId: this.state.id,
                lotNumber: this.state.lotNumber,
                quantity: this.state.quantity,
                sampleSize: this.state.sampleSize,
                createdOn: this.state.createdOn,
                inspectionType: this.state.inspectionType,
                status: this.state.status
            };
        console.log('inspectionLot => ' + JSON.stringify(inspectionLot));

        // step 5
        if(this.state.id === '_add'){
            inspectionLot.inspectionLotId=''
            InspectionLotService.createInspectionLot(inspectionLot).then(res =>{
                this.props.history.push('/inspectionLots');
            });
        }else{
            InspectionLotService.updateInspectionLot(inspectionLot).then( res => {
                this.props.history.push('/inspectionLots');
            });
        }
    }
    
    changelotNumberHandler= (event) => {
        this.setState({lotNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changesampleSizeHandler= (event) => {
        this.setState({sampleSize: event.target.value});
    }
    changecreatedOnHandler= (event) => {
        this.setState({createdOn: event.target.value});
    }
    changeInspectionTypeHandler= (event) => {
        this.setState({inspectionType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/inspectionLots');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InspectionLot</h3>
        }else{
            return <h3 className="text-center">Update InspectionLot</h3>
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
                                            <label> lotNumber:&emsp; </label>
                                                <input placeholder="lotNumber" name="lotNumber" className="form-control" value={this.state.lotNumber} onChange={this.changelotNumberHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> sampleSize:&emsp; </label>
                                                <input type="number" placeholder="sampleSize" name="sampleSize" className="form-control" value={this.state.sampleSize} onChange={this.changesampleSizeHandler}/>

                                            <label> createdOn:&emsp; </label>
                                                <input type="time" placeholder="createdOn" name="createdOn" className="form-control" value={this.state.createdOn} onChange={this.changecreatedOnHandler}/>

                                            <label> InspectionType:&emsp; </label>
                                                <select value={this.state.inspectionType} onChange={this.changeInspectionTypeHandler}>
                      <option name="InspectionType" className="form-control" >
                          Incoming
                      </option>
                      <option name="InspectionType" className="form-control" >
                          InProcess
                      </option>
                      <option name="InspectionType" className="form-control" >
                          Final
                      </option>
                      <option name="InspectionType" className="form-control" >
                          Audit
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Accepted
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInspectionLot}>Save</button>
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

export default CreateInspectionLotComponent
