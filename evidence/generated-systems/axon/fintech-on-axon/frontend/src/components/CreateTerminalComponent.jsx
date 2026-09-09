import React, { Component } from 'react'
import TerminalService from '../services/TerminalService';

class CreateTerminalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                location: '',
                type: '',
                status: ''
        }
        this.changelocationHandler = this.changelocationHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TerminalService.getTerminalById(this.state.id).then( (res) =>{
                let terminal = res.data;
                this.setState({
                    location: terminal.location,
                    type: terminal.type,
                    status: terminal.status
                });
            });
        }        
    }
    saveOrUpdateTerminal = (e) => {
        e.preventDefault();
        let terminal = {
                terminalId: this.state.id,
                location: this.state.location,
                type: this.state.type,
                status: this.state.status
            };
        console.log('terminal => ' + JSON.stringify(terminal));

        // step 5
        if(this.state.id === '_add'){
            terminal.terminalId=''
            TerminalService.createTerminal(terminal).then(res =>{
                this.props.history.push('/terminals');
            });
        }else{
            TerminalService.updateTerminal(terminal).then( res => {
                this.props.history.push('/terminals');
            });
        }
    }
    
    changelocationHandler= (event) => {
        this.setState({location: event.target.value});
    }
    changeTypeHandler= (event) => {
        this.setState({type: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/terminals');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Terminal</h3>
        }else{
            return <h3 className="text-center">Update Terminal</h3>
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
                                            <label> location:&emsp; </label>
                                                <input placeholder="location" name="location" className="form-control" value={this.state.location} onChange={this.changelocationHandler}/>

                                            <label> Type:&emsp; </label>
                                                <select value={this.state.type} onChange={this.changeTypeHandler}>
                      <option name="Type" className="form-control" >
                          POS
                      </option>
                      <option name="Type" className="form-control" >
                          mPOS
                      </option>
                      <option name="Type" className="form-control" >
                          ECommerce
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Inactive
                      </option>
                      <option name="Status" className="form-control" >
                          Decommissioned
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTerminal}>Save</button>
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

export default CreateTerminalComponent
