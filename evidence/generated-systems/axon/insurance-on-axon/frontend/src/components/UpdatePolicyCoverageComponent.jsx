import React, { Component } from 'react'
import PolicyCoverageService from '../services/PolicyCoverageService';

class UpdatePolicyCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                limit: '',
                deductible: '',
                premium: '',
                coverageType: ''
        }
        this.updatePolicyCoverage = this.updatePolicyCoverage.bind(this);

        this.changelimitHandler = this.changelimitHandler.bind(this);
        this.changedeductibleHandler = this.changedeductibleHandler.bind(this);
        this.changepremiumHandler = this.changepremiumHandler.bind(this);
        this.changeCoverageTypeHandler = this.changeCoverageTypeHandler.bind(this);
    }

    componentDidMount(){
        PolicyCoverageService.getPolicyCoverageById(this.state.id).then( (res) =>{
            let policyCoverage = res.data;
            this.setState({
                limit: policyCoverage.limit,
                deductible: policyCoverage.deductible,
                premium: policyCoverage.premium,
                coverageType: policyCoverage.coverageType
            });
        });
    }

    updatePolicyCoverage = (e) => {
        e.preventDefault();
        let policyCoverage = {
            policyCoverageId: this.state.id,
            limit: this.state.limit,
            deductible: this.state.deductible,
            premium: this.state.premium,
            coverageType: this.state.coverageType
        };
        console.log('policyCoverage => ' + JSON.stringify(policyCoverage));
        console.log('id => ' + JSON.stringify(this.state.id));
        PolicyCoverageService.updatePolicyCoverage(policyCoverage).then( res => {
            this.props.history.push('/policyCoverages');
        });
    }

    changelimitHandler= (event) => {
        this.setState({limit: event.target.value});
    }
    changedeductibleHandler= (event) => {
        this.setState({deductible: event.target.value});
    }
    changepremiumHandler= (event) => {
        this.setState({premium: event.target.value});
    }
    changeCoverageTypeHandler= (event) => {
        this.setState({coverageType: event.target.value});
    }

    cancel(){
        this.props.history.push('/policyCoverages');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PolicyCoverage</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> limit: </label>
                                                <input placeholder="limit" name="limit" className="form-control" value={this.state.limit} onChange={this.changelimitHandler}/>

                                            <label> deductible: </label>
                                                <input placeholder="deductible" name="deductible" className="form-control" value={this.state.deductible} onChange={this.changedeductibleHandler}/>

                                            <label> premium: </label>
                                                <input placeholder="premium" name="premium" className="form-control" value={this.state.premium} onChange={this.changepremiumHandler}/>

                                            <label> CoverageType: </label>
                                                <select value={this.state.coverageType} onChange={this.changeCoverageTypeHandler}>
                      <option name="CoverageType" className="form-control" >
                          Liability
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Collision
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Comprehensive
                      </option>
                      <option name="CoverageType" className="form-control" >
                          PropertyDamage
                      </option>
                      <option name="CoverageType" className="form-control" >
                          BodilyInjury
                      </option>
                      <option name="CoverageType" className="form-control" >
                          UninsuredMotorist
                      </option>
                      <option name="CoverageType" className="form-control" >
                          MedicalPayments
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Dwelling
                      </option>
                      <option name="CoverageType" className="form-control" >
                          Contents
                      </option>
                      <option name="CoverageType" className="form-control" >
                          PersonalLiability
                      </option>
                      <option name="CoverageType" className="form-control" >
                          BusinessInterruption
                      </option>
                      <option name="CoverageType" className="form-control" >
                          ProfessionalLiability
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePolicyCoverage}>Save</button>
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

export default UpdatePolicyCoverageComponent
