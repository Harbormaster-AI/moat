import React, { Component } from 'react'
import LaboratoryOrderService from '../services/LaboratoryOrderService';

class UpdateLaboratoryOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                testCode: '',
                fastingRequired: '',
                specimenType: ''
        }
        this.updateLaboratoryOrder = this.updateLaboratoryOrder.bind(this);

        this.changetestCodeHandler = this.changetestCodeHandler.bind(this);
        this.changefastingRequiredHandler = this.changefastingRequiredHandler.bind(this);
        this.changeSpecimenTypeHandler = this.changeSpecimenTypeHandler.bind(this);
    }

    componentDidMount(){
        LaboratoryOrderService.getLaboratoryOrderById(this.state.id).then( (res) =>{
            let laboratoryOrder = res.data;
            this.setState({
                testCode: laboratoryOrder.testCode,
                fastingRequired: laboratoryOrder.fastingRequired,
                specimenType: laboratoryOrder.specimenType
            });
        });
    }

    updateLaboratoryOrder = (e) => {
        e.preventDefault();
        let laboratoryOrder = {
            laboratoryOrderId: this.state.id,
            testCode: this.state.testCode,
            fastingRequired: this.state.fastingRequired,
            specimenType: this.state.specimenType
        };
        console.log('laboratoryOrder => ' + JSON.stringify(laboratoryOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        LaboratoryOrderService.updateLaboratoryOrder(laboratoryOrder).then( res => {
            this.props.history.push('/laboratoryOrders');
        });
    }

    changetestCodeHandler= (event) => {
        this.setState({testCode: event.target.value});
    }
    changefastingRequiredHandler= (event) => {
        this.setState({fastingRequired: event.target.value});
    }
    changeSpecimenTypeHandler= (event) => {
        this.setState({specimenType: event.target.value});
    }

    cancel(){
        this.props.history.push('/laboratoryOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LaboratoryOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> testCode: </label>
                                                <input placeholder="testCode" name="testCode" className="form-control" value={this.state.testCode} onChange={this.changetestCodeHandler}/>

                                            <label> fastingRequired: </label>
                                                <input type="checkbox" placeholder="fastingRequired" name="fastingRequired" className="form-control" value={this.state.fastingRequired} onChange={this.changefastingRequiredHandler}/>


                                            <label> SpecimenType: </label>
                                                <select value={this.state.specimenType} onChange={this.changeSpecimenTypeHandler}>
                      <option name="SpecimenType" className="form-control" >
                          Blood
                      </option>
                      <option name="SpecimenType" className="form-control" >
                          Urine
                      </option>
                      <option name="SpecimenType" className="form-control" >
                          Saliva
                      </option>
                      <option name="SpecimenType" className="form-control" >
                          Sputum
                      </option>
                      <option name="SpecimenType" className="form-control" >
                          Tissue
                      </option>
                      <option name="SpecimenType" className="form-control" >
                          CSF
                      </option>
                      <option name="SpecimenType" className="form-control" >
                          Stool
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLaboratoryOrder}>Save</button>
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

export default UpdateLaboratoryOrderComponent
