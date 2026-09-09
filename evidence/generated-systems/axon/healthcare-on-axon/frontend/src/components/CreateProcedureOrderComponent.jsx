import React, { Component } from 'react'
import ProcedureOrderService from '../services/ProcedureOrderService';

class CreateProcedureOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                procedureCode: '',
                consentObtained: '',
                anesthesiaType: ''
        }
        this.changeprocedureCodeHandler = this.changeprocedureCodeHandler.bind(this);
        this.changeconsentObtainedHandler = this.changeconsentObtainedHandler.bind(this);
        this.changeAnesthesiaTypeHandler = this.changeAnesthesiaTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProcedureOrderService.getProcedureOrderById(this.state.id).then( (res) =>{
                let procedureOrder = res.data;
                this.setState({
                    procedureCode: procedureOrder.procedureCode,
                    consentObtained: procedureOrder.consentObtained,
                    anesthesiaType: procedureOrder.anesthesiaType
                });
            });
        }        
    }
    saveOrUpdateProcedureOrder = (e) => {
        e.preventDefault();
        let procedureOrder = {
                procedureOrderId: this.state.id,
                procedureCode: this.state.procedureCode,
                consentObtained: this.state.consentObtained,
                anesthesiaType: this.state.anesthesiaType
            };
        console.log('procedureOrder => ' + JSON.stringify(procedureOrder));

        // step 5
        if(this.state.id === '_add'){
            procedureOrder.procedureOrderId=''
            ProcedureOrderService.createProcedureOrder(procedureOrder).then(res =>{
                this.props.history.push('/procedureOrders');
            });
        }else{
            ProcedureOrderService.updateProcedureOrder(procedureOrder).then( res => {
                this.props.history.push('/procedureOrders');
            });
        }
    }
    
    changeprocedureCodeHandler= (event) => {
        this.setState({procedureCode: event.target.value});
    }
    changeconsentObtainedHandler= (event) => {
        this.setState({consentObtained: event.target.value});
    }
    changeAnesthesiaTypeHandler= (event) => {
        this.setState({anesthesiaType: event.target.value});
    }

    cancel(){
        this.props.history.push('/procedureOrders');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProcedureOrder</h3>
        }else{
            return <h3 className="text-center">Update ProcedureOrder</h3>
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
                                            <label> procedureCode:&emsp; </label>
                                                <input placeholder="procedureCode" name="procedureCode" className="form-control" value={this.state.procedureCode} onChange={this.changeprocedureCodeHandler}/>

                                            <label> consentObtained:&emsp; </label>
                                                <input type="checkbox" placeholder="consentObtained" name="consentObtained" className="form-control" value={this.state.consentObtained} onChange={this.changeconsentObtainedHandler}/>


                                            <label> AnesthesiaType:&emsp; </label>
                                                <select value={this.state.anesthesiaType} onChange={this.changeAnesthesiaTypeHandler}>
                      <option name="AnesthesiaType" className="form-control" >
                          None
                      </option>
                      <option name="AnesthesiaType" className="form-control" >
                          Local
                      </option>
                      <option name="AnesthesiaType" className="form-control" >
                          Regional
                      </option>
                      <option name="AnesthesiaType" className="form-control" >
                          General
                      </option>
                      <option name="AnesthesiaType" className="form-control" >
                          Sedation
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProcedureOrder}>Save</button>
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

export default CreateProcedureOrderComponent
