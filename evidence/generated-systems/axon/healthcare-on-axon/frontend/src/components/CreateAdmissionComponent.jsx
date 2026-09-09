import React, { Component } from 'react'
import AdmissionService from '../services/AdmissionService';

class CreateAdmissionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                admitDateTime: '',
                bed: '',
                admissionType: ''
        }
        this.changeadmitDateTimeHandler = this.changeadmitDateTimeHandler.bind(this);
        this.changebedHandler = this.changebedHandler.bind(this);
        this.changeAdmissionTypeHandler = this.changeAdmissionTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AdmissionService.getAdmissionById(this.state.id).then( (res) =>{
                let admission = res.data;
                this.setState({
                    admitDateTime: admission.admitDateTime,
                    bed: admission.bed,
                    admissionType: admission.admissionType
                });
            });
        }        
    }
    saveOrUpdateAdmission = (e) => {
        e.preventDefault();
        let admission = {
                admissionId: this.state.id,
                admitDateTime: this.state.admitDateTime,
                bed: this.state.bed,
                admissionType: this.state.admissionType
            };
        console.log('admission => ' + JSON.stringify(admission));

        // step 5
        if(this.state.id === '_add'){
            admission.admissionId=''
            AdmissionService.createAdmission(admission).then(res =>{
                this.props.history.push('/admissions');
            });
        }else{
            AdmissionService.updateAdmission(admission).then( res => {
                this.props.history.push('/admissions');
            });
        }
    }
    
    changeadmitDateTimeHandler= (event) => {
        this.setState({admitDateTime: event.target.value});
    }
    changebedHandler= (event) => {
        this.setState({bed: event.target.value});
    }
    changeAdmissionTypeHandler= (event) => {
        this.setState({admissionType: event.target.value});
    }

    cancel(){
        this.props.history.push('/admissions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Admission</h3>
        }else{
            return <h3 className="text-center">Update Admission</h3>
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
                                            <label> admitDateTime:&emsp; </label>
                                                <input type="time" placeholder="admitDateTime" name="admitDateTime" className="form-control" value={this.state.admitDateTime} onChange={this.changeadmitDateTimeHandler}/>

                                            <label> bed:&emsp; </label>
                                                <input placeholder="bed" name="bed" className="form-control" value={this.state.bed} onChange={this.changebedHandler}/>

                                            <label> AdmissionType:&emsp; </label>
                                                <select value={this.state.admissionType} onChange={this.changeAdmissionTypeHandler}>
                      <option name="AdmissionType" className="form-control" >
                          Elective
                      </option>
                      <option name="AdmissionType" className="form-control" >
                          Emergency
                      </option>
                      <option name="AdmissionType" className="form-control" >
                          Urgent
                      </option>
                      <option name="AdmissionType" className="form-control" >
                          Newborn
                      </option>
                      <option name="AdmissionType" className="form-control" >
                          Trauma
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAdmission}>Save</button>
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

export default CreateAdmissionComponent
