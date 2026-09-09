import React, { Component } from 'react'
import CoverageService from '../services/CoverageService';

class CreateCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                memberId: '',
                groupNumber: '',
                effectiveDate: '',
                endDate: '',
                coverageType: ''
        }
        this.changememberIdHandler = this.changememberIdHandler.bind(this);
        this.changegroupNumberHandler = this.changegroupNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeCoverageTypeHandler = this.changeCoverageTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CoverageService.getCoverageById(this.state.id).then( (res) =>{
                let coverage = res.data;
                this.setState({
                    memberId: coverage.memberId,
                    groupNumber: coverage.groupNumber,
                    effectiveDate: coverage.effectiveDate,
                    endDate: coverage.endDate,
                    coverageType: coverage.coverageType
                });
            });
        }        
    }
    saveOrUpdateCoverage = (e) => {
        e.preventDefault();
        let coverage = {
                coverageId: this.state.id,
                memberId: this.state.memberId,
                groupNumber: this.state.groupNumber,
                effectiveDate: this.state.effectiveDate,
                endDate: this.state.endDate,
                coverageType: this.state.coverageType
            };
        console.log('coverage => ' + JSON.stringify(coverage));

        // step 5
        if(this.state.id === '_add'){
            coverage.coverageId=''
            CoverageService.createCoverage(coverage).then(res =>{
                this.props.history.push('/coverages');
            });
        }else{
            CoverageService.updateCoverage(coverage).then( res => {
                this.props.history.push('/coverages');
            });
        }
    }
    
    changememberIdHandler= (event) => {
        this.setState({memberId: event.target.value});
    }
    changegroupNumberHandler= (event) => {
        this.setState({groupNumber: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeCoverageTypeHandler= (event) => {
        this.setState({coverageType: event.target.value});
    }

    cancel(){
        this.props.history.push('/coverages');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Coverage</h3>
        }else{
            return <h3 className="text-center">Update Coverage</h3>
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
                                            <label> memberId:&emsp; </label>
                                                <input placeholder="memberId" name="memberId" className="form-control" value={this.state.memberId} onChange={this.changememberIdHandler}/>

                                            <label> groupNumber:&emsp; </label>
                                                <input placeholder="groupNumber" name="groupNumber" className="form-control" value={this.state.groupNumber} onChange={this.changegroupNumberHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> CoverageType:&emsp; </label>
                                                <select value={this.state.coverageType} onChange={this.changeCoverageTypeHandler}>
                      <option name="CoverageType" className="form-control" >
                          Medical
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Pharmacy
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Dental
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Vision
                      </option>
                      <option name="CoverageType" className="form-control" >
                          BehavioralHealth
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCoverage}>Save</button>
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

export default CreateCoverageComponent
