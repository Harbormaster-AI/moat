import React, { Component } from 'react'
import ProcedureService from '../services/ProcedureService';

class CreateProcedureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                versionLabel: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeversionLabelHandler = this.changeversionLabelHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProcedureService.getProcedureById(this.state.id).then( (res) =>{
                let procedure = res.data;
                this.setState({
                    title: procedure.title,
                    versionLabel: procedure.versionLabel,
                    status: procedure.status
                });
            });
        }        
    }
    saveOrUpdateProcedure = (e) => {
        e.preventDefault();
        let procedure = {
                procedureId: this.state.id,
                title: this.state.title,
                versionLabel: this.state.versionLabel,
                status: this.state.status
            };
        console.log('procedure => ' + JSON.stringify(procedure));

        // step 5
        if(this.state.id === '_add'){
            procedure.procedureId=''
            ProcedureService.createProcedure(procedure).then(res =>{
                this.props.history.push('/procedures');
            });
        }else{
            ProcedureService.updateProcedure(procedure).then( res => {
                this.props.history.push('/procedures');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeversionLabelHandler= (event) => {
        this.setState({versionLabel: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/procedures');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Procedure</h3>
        }else{
            return <h3 className="text-center">Update Procedure</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> versionLabel:&emsp; </label>
                                                <input placeholder="versionLabel" name="versionLabel" className="form-control" value={this.state.versionLabel} onChange={this.changeversionLabelHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          InReview
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProcedure}>Save</button>
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

export default CreateProcedureComponent
