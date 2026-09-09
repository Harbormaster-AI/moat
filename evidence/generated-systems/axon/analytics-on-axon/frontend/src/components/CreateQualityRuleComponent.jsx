import React, { Component } from 'react'
import QualityRuleService from '../services/QualityRuleService';

class CreateQualityRuleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                threshold: '',
                targetField: '',
                dimension: '',
                operator: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changethresholdHandler = this.changethresholdHandler.bind(this);
        this.changetargetFieldHandler = this.changetargetFieldHandler.bind(this);
        this.changeDimensionHandler = this.changeDimensionHandler.bind(this);
        this.changeOperatorHandler = this.changeOperatorHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            QualityRuleService.getQualityRuleById(this.state.id).then( (res) =>{
                let qualityRule = res.data;
                this.setState({
                    name: qualityRule.name,
                    threshold: qualityRule.threshold,
                    targetField: qualityRule.targetField,
                    dimension: qualityRule.dimension,
                    operator: qualityRule.operator
                });
            });
        }        
    }
    saveOrUpdateQualityRule = (e) => {
        e.preventDefault();
        let qualityRule = {
                qualityRuleId: this.state.id,
                name: this.state.name,
                threshold: this.state.threshold,
                targetField: this.state.targetField,
                dimension: this.state.dimension,
                operator: this.state.operator
            };
        console.log('qualityRule => ' + JSON.stringify(qualityRule));

        // step 5
        if(this.state.id === '_add'){
            qualityRule.qualityRuleId=''
            QualityRuleService.createQualityRule(qualityRule).then(res =>{
                this.props.history.push('/qualityRules');
            });
        }else{
            QualityRuleService.updateQualityRule(qualityRule).then( res => {
                this.props.history.push('/qualityRules');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changethresholdHandler= (event) => {
        this.setState({threshold: event.target.value});
    }
    changetargetFieldHandler= (event) => {
        this.setState({targetField: event.target.value});
    }
    changeDimensionHandler= (event) => {
        this.setState({dimension: event.target.value});
    }
    changeOperatorHandler= (event) => {
        this.setState({operator: event.target.value});
    }

    cancel(){
        this.props.history.push('/qualityRules');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add QualityRule</h3>
        }else{
            return <h3 className="text-center">Update QualityRule</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> threshold:&emsp; </label>
                                                <input placeholder="threshold" name="threshold" className="form-control" value={this.state.threshold} onChange={this.changethresholdHandler}/>

                                            <label> targetField:&emsp; </label>
                                                <input placeholder="targetField" name="targetField" className="form-control" value={this.state.targetField} onChange={this.changetargetFieldHandler}/>

                                            <label> Dimension:&emsp; </label>
                                                <select value={this.state.dimension} onChange={this.changeDimensionHandler}>
                      <option name="Dimension" className="form-control" >
                          Completeness
                      </option>
                      <option name="Dimension" className="form-control" >
                          Accuracy
                      </option>
                      <option name="Dimension" className="form-control" >
                          Consistency
                      </option>
                      <option name="Dimension" className="form-control" >
                          Timeliness
                      </option>
                      <option name="Dimension" className="form-control" >
                          Uniqueness
                      </option>
                      <option name="Dimension" className="form-control" >
                          Validity
                      </option>
                    </select>

                                            <label> Operator:&emsp; </label>
                                                <select value={this.state.operator} onChange={this.changeOperatorHandler}>
                      <option name="Operator" className="form-control" >
                          GreaterThan
                      </option>
                      <option name="Operator" className="form-control" >
                          GreaterThanOrEqual
                      </option>
                      <option name="Operator" className="form-control" >
                          LessThan
                      </option>
                      <option name="Operator" className="form-control" >
                          LessThanOrEqual
                      </option>
                      <option name="Operator" className="form-control" >
                          Equal
                      </option>
                      <option name="Operator" className="form-control" >
                          NotEqual
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateQualityRule}>Save</button>
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

export default CreateQualityRuleComponent
