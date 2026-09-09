import React, { Component } from 'react'
import MatterService from '../services/MatterService';

class CreateMatterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                matterName: '',
                leadCounsel: '',
                matterType: '',
                status: ''
        }
        this.changematterNameHandler = this.changematterNameHandler.bind(this);
        this.changeleadCounselHandler = this.changeleadCounselHandler.bind(this);
        this.changeMatterTypeHandler = this.changeMatterTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateMatter = (e) => {
        e.preventDefault();
        let matter = {
                matterId: this.state.id,
                matterName: this.state.matterName,
                leadCounsel: this.state.leadCounsel,
                matterType: this.state.matterType,
                status: this.state.status
            };
        console.log('matter => ' + JSON.stringify(matter));

        // step 5
        if(this.state.id === '_add'){
            matter.matterId=''
            MatterService.createMatter(matter).then(res =>{
                this.props.history.push('/matters');
            });
        }else{
            MatterService.updateMatter(matter).then( res => {
                this.props.history.push('/matters');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Matter</h3>
        }else{
            return <h3 className="text-center">Update Matter</h3>
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
                                            <label> matterName:&emsp; </label>
                                                <input placeholder="matterName" name="matterName" className="form-control" value={this.state.matterName} onChange={this.changematterNameHandler}/>

                                            <label> leadCounsel:&emsp; </label>
                                                <input placeholder="leadCounsel" name="leadCounsel" className="form-control" value={this.state.leadCounsel} onChange={this.changeleadCounselHandler}/>

                                            <label> MatterType:&emsp; </label>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMatter}>Save</button>
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

export default CreateMatterComponent
