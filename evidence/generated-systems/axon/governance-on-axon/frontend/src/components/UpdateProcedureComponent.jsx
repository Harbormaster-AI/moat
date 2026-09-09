import React, { Component } from 'react'
import ProcedureService from '../services/ProcedureService';

class UpdateProcedureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                versionLabel: '',
                status: ''
        }
        this.updateProcedure = this.updateProcedure.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeversionLabelHandler = this.changeversionLabelHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ProcedureService.getProcedureById(this.state.id).then( (res) =>{
            let procedure = res.data;
            this.setState({
                title: procedure.title,
                versionLabel: procedure.versionLabel,
                status: procedure.status
            });
        });
    }

    updateProcedure = (e) => {
        e.preventDefault();
        let procedure = {
            procedureId: this.state.id,
            title: this.state.title,
            versionLabel: this.state.versionLabel,
            status: this.state.status
        };
        console.log('procedure => ' + JSON.stringify(procedure));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProcedureService.updateProcedure(procedure).then( res => {
            this.props.history.push('/procedures');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Procedure</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> versionLabel: </label>
                                                <input placeholder="versionLabel" name="versionLabel" className="form-control" value={this.state.versionLabel} onChange={this.changeversionLabelHandler}/>

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateProcedure}>Save</button>
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

export default UpdateProcedureComponent
