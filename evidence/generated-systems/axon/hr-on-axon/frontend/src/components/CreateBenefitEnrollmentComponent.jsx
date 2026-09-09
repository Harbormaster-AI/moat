import React, { Component } from 'react'
import BenefitEnrollmentService from '../services/BenefitEnrollmentService';

class CreateBenefitEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                enrollmentId: '',
                effectiveFrom: '',
                effectiveTo: '',
                status: '',
                coverageLevel: ''
        }
        this.changeenrollmentIdHandler = this.changeenrollmentIdHandler.bind(this);
        this.changeeffectiveFromHandler = this.changeeffectiveFromHandler.bind(this);
        this.changeeffectiveToHandler = this.changeeffectiveToHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeCoverageLevelHandler = this.changeCoverageLevelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BenefitEnrollmentService.getBenefitEnrollmentById(this.state.id).then( (res) =>{
                let benefitEnrollment = res.data;
                this.setState({
                    enrollmentId: benefitEnrollment.enrollmentId,
                    effectiveFrom: benefitEnrollment.effectiveFrom,
                    effectiveTo: benefitEnrollment.effectiveTo,
                    status: benefitEnrollment.status,
                    coverageLevel: benefitEnrollment.coverageLevel
                });
            });
        }        
    }
    saveOrUpdateBenefitEnrollment = (e) => {
        e.preventDefault();
        let benefitEnrollment = {
                benefitEnrollmentId: this.state.id,
                enrollmentId: this.state.enrollmentId,
                effectiveFrom: this.state.effectiveFrom,
                effectiveTo: this.state.effectiveTo,
                status: this.state.status,
                coverageLevel: this.state.coverageLevel
            };
        console.log('benefitEnrollment => ' + JSON.stringify(benefitEnrollment));

        // step 5
        if(this.state.id === '_add'){
            benefitEnrollment.benefitEnrollmentId=''
            BenefitEnrollmentService.createBenefitEnrollment(benefitEnrollment).then(res =>{
                this.props.history.push('/benefitEnrollments');
            });
        }else{
            BenefitEnrollmentService.updateBenefitEnrollment(benefitEnrollment).then( res => {
                this.props.history.push('/benefitEnrollments');
            });
        }
    }
    
    changeenrollmentIdHandler= (event) => {
        this.setState({enrollmentId: event.target.value});
    }
    changeeffectiveFromHandler= (event) => {
        this.setState({effectiveFrom: event.target.value});
    }
    changeeffectiveToHandler= (event) => {
        this.setState({effectiveTo: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeCoverageLevelHandler= (event) => {
        this.setState({coverageLevel: event.target.value});
    }

    cancel(){
        this.props.history.push('/benefitEnrollments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BenefitEnrollment</h3>
        }else{
            return <h3 className="text-center">Update BenefitEnrollment</h3>
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
                                            <label> enrollmentId:&emsp; </label>
                                                <input placeholder="enrollmentId" name="enrollmentId" className="form-control" value={this.state.enrollmentId} onChange={this.changeenrollmentIdHandler}/>

                                            <label> effectiveFrom:&emsp; </label>
                                                <input type="date" placeholder="effectiveFrom" name="effectiveFrom" className="form-control" value={this.state.effectiveFrom} onChange={this.changeeffectiveFromHandler}/>

                                            <label> effectiveTo:&emsp; </label>
                                                <input type="date" placeholder="effectiveTo" name="effectiveTo" className="form-control" value={this.state.effectiveTo} onChange={this.changeeffectiveToHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Waived
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                    </select>

                                            <label> CoverageLevel:&emsp; </label>
                                                <select value={this.state.coverageLevel} onChange={this.changeCoverageLevelHandler}>
                      <option name="CoverageLevel" className="form-control" >
                          EmployeeOnly
                      </option>
                      <option name="CoverageLevel" className="form-control" >
                          EmployeeSpouse
                      </option>
                      <option name="CoverageLevel" className="form-control" >
                          EmployeeChildren
                      </option>
                      <option name="CoverageLevel" className="form-control" >
                          Family
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBenefitEnrollment}>Save</button>
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

export default CreateBenefitEnrollmentComponent
