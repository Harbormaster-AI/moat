import React, { Component } from 'react'
import PolicyCoverageService from '../services/PolicyCoverageService';

class CreatePolicyCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                limit: '',
                deductible: '',
                premium: '',
                coverageType: ''
        }
        this.changelimitHandler = this.changelimitHandler.bind(this);
        this.changedeductibleHandler = this.changedeductibleHandler.bind(this);
        this.changepremiumHandler = this.changepremiumHandler.bind(this);
        this.changeCoverageTypeHandler = this.changeCoverageTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePolicyCoverage = (e) => {
        e.preventDefault();
        let policyCoverage = {
                policyCoverageId: this.state.id,
                limit: this.state.limit,
                deductible: this.state.deductible,
                premium: this.state.premium,
                coverageType: this.state.coverageType
            };
        console.log('policyCoverage => ' + JSON.stringify(policyCoverage));

        // step 5
        if(this.state.id === '_add'){
            policyCoverage.policyCoverageId=''
            PolicyCoverageService.createPolicyCoverage(policyCoverage).then(res =>{
                this.props.history.push('/policyCoverages');
            });
        }else{
            PolicyCoverageService.updatePolicyCoverage(policyCoverage).then( res => {
                this.props.history.push('/policyCoverages');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PolicyCoverage</h3>
        }else{
            return <h3 className="text-center">Update PolicyCoverage</h3>
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
                                            <label> limit:&emsp; </label>
                                                <input placeholder="limit" name="limit" className="form-control" value={this.state.limit} onChange={this.changelimitHandler}/>

                                            <label> deductible:&emsp; </label>
                                                <input placeholder="deductible" name="deductible" className="form-control" value={this.state.deductible} onChange={this.changedeductibleHandler}/>

                                            <label> premium:&emsp; </label>
                                                <input placeholder="premium" name="premium" className="form-control" value={this.state.premium} onChange={this.changepremiumHandler}/>

                                            <label> CoverageType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePolicyCoverage}>Save</button>
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

export default CreatePolicyCoverageComponent
