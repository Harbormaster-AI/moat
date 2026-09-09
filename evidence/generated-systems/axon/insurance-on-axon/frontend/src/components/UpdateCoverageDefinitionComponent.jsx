import React, { Component } from 'react'
import CoverageDefinitionService from '../services/CoverageDefinitionService';

class UpdateCoverageDefinitionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                defaultLimit: '',
                defaultDeductible: '',
                asMandatory: '',
                coverageType: ''
        }
        this.updateCoverageDefinition = this.updateCoverageDefinition.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedefaultLimitHandler = this.changedefaultLimitHandler.bind(this);
        this.changedefaultDeductibleHandler = this.changedefaultDeductibleHandler.bind(this);
        this.changeasMandatoryHandler = this.changeasMandatoryHandler.bind(this);
        this.changeCoverageTypeHandler = this.changeCoverageTypeHandler.bind(this);
    }

    componentDidMount(){
        CoverageDefinitionService.getCoverageDefinitionById(this.state.id).then( (res) =>{
            let coverageDefinition = res.data;
            this.setState({
                name: coverageDefinition.name,
                defaultLimit: coverageDefinition.defaultLimit,
                defaultDeductible: coverageDefinition.defaultDeductible,
                asMandatory: coverageDefinition.asMandatory,
                coverageType: coverageDefinition.coverageType
            });
        });
    }

    updateCoverageDefinition = (e) => {
        e.preventDefault();
        let coverageDefinition = {
            coverageDefinitionId: this.state.id,
            name: this.state.name,
            defaultLimit: this.state.defaultLimit,
            defaultDeductible: this.state.defaultDeductible,
            asMandatory: this.state.asMandatory,
            coverageType: this.state.coverageType
        };
        console.log('coverageDefinition => ' + JSON.stringify(coverageDefinition));
        console.log('id => ' + JSON.stringify(this.state.id));
        CoverageDefinitionService.updateCoverageDefinition(coverageDefinition).then( res => {
            this.props.history.push('/coverageDefinitions');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedefaultLimitHandler= (event) => {
        this.setState({defaultLimit: event.target.value});
    }
    changedefaultDeductibleHandler= (event) => {
        this.setState({defaultDeductible: event.target.value});
    }
    changeasMandatoryHandler= (event) => {
        this.setState({asMandatory: event.target.value});
    }
    changeCoverageTypeHandler= (event) => {
        this.setState({coverageType: event.target.value});
    }

    cancel(){
        this.props.history.push('/coverageDefinitions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CoverageDefinition</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> defaultLimit: </label>
                                                <input placeholder="defaultLimit" name="defaultLimit" className="form-control" value={this.state.defaultLimit} onChange={this.changedefaultLimitHandler}/>

                                            <label> defaultDeductible: </label>
                                                <input placeholder="defaultDeductible" name="defaultDeductible" className="form-control" value={this.state.defaultDeductible} onChange={this.changedefaultDeductibleHandler}/>

                                            <label> asMandatory: </label>
                                                <input type="checkbox" placeholder="asMandatory" name="asMandatory" className="form-control" value={this.state.asMandatory} onChange={this.changeasMandatoryHandler}/>


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
                                        <button className="btn btn-success" onClick={this.updateCoverageDefinition}>Save</button>
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

export default UpdateCoverageDefinitionComponent
