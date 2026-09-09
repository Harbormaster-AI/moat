import React, { Component } from 'react'
import AdmissionService from '../services/AdmissionService';

class UpdateAdmissionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                admitDateTime: '',
                bed: '',
                admissionType: ''
        }
        this.updateAdmission = this.updateAdmission.bind(this);

        this.changeadmitDateTimeHandler = this.changeadmitDateTimeHandler.bind(this);
        this.changebedHandler = this.changebedHandler.bind(this);
        this.changeAdmissionTypeHandler = this.changeAdmissionTypeHandler.bind(this);
    }

    componentDidMount(){
        AdmissionService.getAdmissionById(this.state.id).then( (res) =>{
            let admission = res.data;
            this.setState({
                admitDateTime: admission.admitDateTime,
                bed: admission.bed,
                admissionType: admission.admissionType
            });
        });
    }

    updateAdmission = (e) => {
        e.preventDefault();
        let admission = {
            admissionId: this.state.id,
            admitDateTime: this.state.admitDateTime,
            bed: this.state.bed,
            admissionType: this.state.admissionType
        };
        console.log('admission => ' + JSON.stringify(admission));
        console.log('id => ' + JSON.stringify(this.state.id));
        AdmissionService.updateAdmission(admission).then( res => {
            this.props.history.push('/admissions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Admission</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> admitDateTime: </label>
                                                <input type="time" placeholder="admitDateTime" name="admitDateTime" className="form-control" value={this.state.admitDateTime} onChange={this.changeadmitDateTimeHandler}/>

                                            <label> bed: </label>
                                                <input placeholder="bed" name="bed" className="form-control" value={this.state.bed} onChange={this.changebedHandler}/>

                                            <label> AdmissionType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateAdmission}>Save</button>
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

export default UpdateAdmissionComponent
