import React, { Component } from 'react'
import QuarantineService from '../services/QuarantineService';

class UpdateQuarantineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                reason: '',
                startedAt: '',
                releasedAt: '',
                disposition: ''
        }
        this.updateQuarantine = this.updateQuarantine.bind(this);

        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changestartedAtHandler = this.changestartedAtHandler.bind(this);
        this.changereleasedAtHandler = this.changereleasedAtHandler.bind(this);
        this.changeDispositionHandler = this.changeDispositionHandler.bind(this);
    }

    componentDidMount(){
        QuarantineService.getQuarantineById(this.state.id).then( (res) =>{
            let quarantine = res.data;
            this.setState({
                reason: quarantine.reason,
                startedAt: quarantine.startedAt,
                releasedAt: quarantine.releasedAt,
                disposition: quarantine.disposition
            });
        });
    }

    updateQuarantine = (e) => {
        e.preventDefault();
        let quarantine = {
            quarantineId: this.state.id,
            reason: this.state.reason,
            startedAt: this.state.startedAt,
            releasedAt: this.state.releasedAt,
            disposition: this.state.disposition
        };
        console.log('quarantine => ' + JSON.stringify(quarantine));
        console.log('id => ' + JSON.stringify(this.state.id));
        QuarantineService.updateQuarantine(quarantine).then( res => {
            this.props.history.push('/quarantines');
        });
    }

    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changestartedAtHandler= (event) => {
        this.setState({startedAt: event.target.value});
    }
    changereleasedAtHandler= (event) => {
        this.setState({releasedAt: event.target.value});
    }
    changeDispositionHandler= (event) => {
        this.setState({disposition: event.target.value});
    }

    cancel(){
        this.props.history.push('/quarantines');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Quarantine</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> startedAt: </label>
                                                <input type="date" placeholder="startedAt" name="startedAt" className="form-control" value={this.state.startedAt} onChange={this.changestartedAtHandler}/>

                                            <label> releasedAt: </label>
                                                <input type="date" placeholder="releasedAt" name="releasedAt" className="form-control" value={this.state.releasedAt} onChange={this.changereleasedAtHandler}/>

                                            <label> Disposition: </label>
                                                <select value={this.state.disposition} onChange={this.changeDispositionHandler}>
                      <option name="Disposition" className="form-control" >
                          Release
                      </option>
                      <option name="Disposition" className="form-control" >
                          Scrap
                      </option>
                      <option name="Disposition" className="form-control" >
                          ReturnToVendor
                      </option>
                      <option name="Disposition" className="form-control" >
                          Rework
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateQuarantine}>Save</button>
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

export default UpdateQuarantineComponent
