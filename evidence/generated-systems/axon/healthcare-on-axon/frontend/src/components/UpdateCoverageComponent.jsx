import React, { Component } from 'react'
import CoverageService from '../services/CoverageService';

class UpdateCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                memberId: '',
                groupNumber: '',
                effectiveDate: '',
                endDate: '',
                coverageType: ''
        }
        this.updateCoverage = this.updateCoverage.bind(this);

        this.changememberIdHandler = this.changememberIdHandler.bind(this);
        this.changegroupNumberHandler = this.changegroupNumberHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeCoverageTypeHandler = this.changeCoverageTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateCoverage = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        CoverageService.updateCoverage(coverage).then( res => {
            this.props.history.push('/coverages');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Coverage</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> memberId: </label>
                                                <input placeholder="memberId" name="memberId" className="form-control" value={this.state.memberId} onChange={this.changememberIdHandler}/>

                                            <label> groupNumber: </label>
                                                <input placeholder="groupNumber" name="groupNumber" className="form-control" value={this.state.groupNumber} onChange={this.changegroupNumberHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> CoverageType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateCoverage}>Save</button>
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

export default UpdateCoverageComponent
