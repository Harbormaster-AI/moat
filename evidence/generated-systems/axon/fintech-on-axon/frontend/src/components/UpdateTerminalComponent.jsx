import React, { Component } from 'react'
import TerminalService from '../services/TerminalService';

class UpdateTerminalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                location: '',
                type: '',
                status: ''
        }
        this.updateTerminal = this.updateTerminal.bind(this);

        this.changelocationHandler = this.changelocationHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        TerminalService.getTerminalById(this.state.id).then( (res) =>{
            let terminal = res.data;
            this.setState({
                location: terminal.location,
                type: terminal.type,
                status: terminal.status
            });
        });
    }

    updateTerminal = (e) => {
        e.preventDefault();
        let terminal = {
            terminalId: this.state.id,
            location: this.state.location,
            type: this.state.type,
            status: this.state.status
        };
        console.log('terminal => ' + JSON.stringify(terminal));
        console.log('id => ' + JSON.stringify(this.state.id));
        TerminalService.updateTerminal(terminal).then( res => {
            this.props.history.push('/terminals');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Terminal</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> location: </label>
                                                <input placeholder="location" name="location" className="form-control" value={this.state.location} onChange={this.changelocationHandler}/>

                                            <label> Type: </label>
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

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateTerminal}>Save</button>
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

export default UpdateTerminalComponent
