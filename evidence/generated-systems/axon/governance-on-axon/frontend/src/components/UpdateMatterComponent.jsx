import React, { Component } from 'react'
import MatterService from '../services/MatterService';

class UpdateMatterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                matterName: '',
                leadCounsel: '',
                matterType: '',
                status: ''
        }
        this.updateMatter = this.updateMatter.bind(this);

        this.changematterNameHandler = this.changematterNameHandler.bind(this);
        this.changeleadCounselHandler = this.changeleadCounselHandler.bind(this);
        this.changeMatterTypeHandler = this.changeMatterTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        MatterService.getMatterById(this.state.id).then( (res) =>{
            let matter = res.data;
            this.setState({
                matterName: matter.matterName,
                leadCounsel: matter.leadCounsel,
                matterType: matter.matterType,
                status: matter.status
            });
        });
    }

    updateMatter = (e) => {
        e.preventDefault();
        let matter = {
            matterId: this.state.id,
            matterName: this.state.matterName,
            leadCounsel: this.state.leadCounsel,
            matterType: this.state.matterType,
            status: this.state.status
        };
        console.log('matter => ' + JSON.stringify(matter));
        console.log('id => ' + JSON.stringify(this.state.id));
        MatterService.updateMatter(matter).then( res => {
            this.props.history.push('/matters');
        });
    }

    changematterNameHandler= (event) => {
        this.setState({matterName: event.target.value});
    }
    changeleadCounselHandler= (event) => {
        this.setState({leadCounsel: event.target.value});
    }
    changeMatterTypeHandler= (event) => {
        this.setState({matterType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/matters');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Matter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> matterName: </label>
                                                <input placeholder="matterName" name="matterName" className="form-control" value={this.state.matterName} onChange={this.changematterNameHandler}/>

                                            <label> leadCounsel: </label>
                                                <input placeholder="leadCounsel" name="leadCounsel" className="form-control" value={this.state.leadCounsel} onChange={this.changeleadCounselHandler}/>

                                            <label> MatterType: </label>
                                                <select value={this.state.matterType} onChange={this.changeMatterTypeHandler}>
                      <option name="MatterType" className="form-control" >
                          Litigation
                      </option>
                      <option name="MatterType" className="form-control" >
                          Investigation
                      </option>
                      <option name="MatterType" className="form-control" >
                          RegulatoryInquiry
                      </option>
                      <option name="MatterType" className="form-control" >
                          Complaint
                      </option>
                      <option name="MatterType" className="form-control" >
                          Arbitration
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          ActiveDiscovery
                      </option>
                      <option name="Status" className="form-control" >
                          Negotiation
                      </option>
                      <option name="Status" className="form-control" >
                          Settled
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMatter}>Save</button>
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

export default UpdateMatterComponent
