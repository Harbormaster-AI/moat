import React, { Component } from 'react'
import NonconformanceService from '../services/NonconformanceService';

class CreateNonconformanceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                ncNumber: '',
                description: '',
                containmentAction: '',
                ncType: '',
                severity: '',
                status: ''
        }
        this.changencNumberHandler = this.changencNumberHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changecontainmentActionHandler = this.changecontainmentActionHandler.bind(this);
        this.changeNcTypeHandler = this.changeNcTypeHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            NonconformanceService.getNonconformanceById(this.state.id).then( (res) =>{
                let nonconformance = res.data;
                this.setState({
                    ncNumber: nonconformance.ncNumber,
                    description: nonconformance.description,
                    containmentAction: nonconformance.containmentAction,
                    ncType: nonconformance.ncType,
                    severity: nonconformance.severity,
                    status: nonconformance.status
                });
            });
        }        
    }
    saveOrUpdateNonconformance = (e) => {
        e.preventDefault();
        let nonconformance = {
                nonconformanceId: this.state.id,
                ncNumber: this.state.ncNumber,
                description: this.state.description,
                containmentAction: this.state.containmentAction,
                ncType: this.state.ncType,
                severity: this.state.severity,
                status: this.state.status
            };
        console.log('nonconformance => ' + JSON.stringify(nonconformance));

        // step 5
        if(this.state.id === '_add'){
            nonconformance.nonconformanceId=''
            NonconformanceService.createNonconformance(nonconformance).then(res =>{
                this.props.history.push('/nonconformances');
            });
        }else{
            NonconformanceService.updateNonconformance(nonconformance).then( res => {
                this.props.history.push('/nonconformances');
            });
        }
    }
    
    changencNumberHandler= (event) => {
        this.setState({ncNumber: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changecontainmentActionHandler= (event) => {
        this.setState({containmentAction: event.target.value});
    }
    changeNcTypeHandler= (event) => {
        this.setState({ncType: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/nonconformances');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Nonconformance</h3>
        }else{
            return <h3 className="text-center">Update Nonconformance</h3>
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
                                            <label> ncNumber:&emsp; </label>
                                                <input placeholder="ncNumber" name="ncNumber" className="form-control" value={this.state.ncNumber} onChange={this.changencNumberHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> containmentAction:&emsp; </label>
                                                <input placeholder="containmentAction" name="containmentAction" className="form-control" value={this.state.containmentAction} onChange={this.changecontainmentActionHandler}/>

                                            <label> NcType:&emsp; </label>
                                                <select value={this.state.ncType} onChange={this.changeNcTypeHandler}>
                      <option name="NcType" className="form-control" >
                          Dimension
                      </option>
                      <option name="NcType" className="form-control" >
                          Functional
                      </option>
                      <option name="NcType" className="form-control" >
                          Cosmetic
                      </option>
                      <option name="NcType" className="form-control" >
                          Documentation
                      </option>
                      <option name="NcType" className="form-control" >
                          Supplier
                      </option>
                      <option name="NcType" className="form-control" >
                          Process
                      </option>
                    </select>

                                            <label> Severity:&emsp; </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Minor
                      </option>
                      <option name="Severity" className="form-control" >
                          Major
                      </option>
                      <option name="Severity" className="form-control" >
                          Critical
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Contained
                      </option>
                      <option name="Status" className="form-control" >
                          UnderInvestigation
                      </option>
                      <option name="Status" className="form-control" >
                          Dispositioned
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateNonconformance}>Save</button>
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

export default CreateNonconformanceComponent
