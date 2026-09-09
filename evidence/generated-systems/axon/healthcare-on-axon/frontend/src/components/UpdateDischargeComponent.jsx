import React, { Component } from 'react'
import DischargeService from '../services/DischargeService';

class UpdateDischargeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                dischargeDateTime: '',
                disposition: ''
        }
        this.updateDischarge = this.updateDischarge.bind(this);

        this.changedischargeDateTimeHandler = this.changedischargeDateTimeHandler.bind(this);
        this.changeDispositionHandler = this.changeDispositionHandler.bind(this);
    }

    componentDidMount(){
        DischargeService.getDischargeById(this.state.id).then( (res) =>{
            let discharge = res.data;
            this.setState({
                dischargeDateTime: discharge.dischargeDateTime,
                disposition: discharge.disposition
            });
        });
    }

    updateDischarge = (e) => {
        e.preventDefault();
        let discharge = {
            dischargeId: this.state.id,
            dischargeDateTime: this.state.dischargeDateTime,
            disposition: this.state.disposition
        };
        console.log('discharge => ' + JSON.stringify(discharge));
        console.log('id => ' + JSON.stringify(this.state.id));
        DischargeService.updateDischarge(discharge).then( res => {
            this.props.history.push('/discharges');
        });
    }

    changedischargeDateTimeHandler= (event) => {
        this.setState({dischargeDateTime: event.target.value});
    }
    changeDispositionHandler= (event) => {
        this.setState({disposition: event.target.value});
    }

    cancel(){
        this.props.history.push('/discharges');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Discharge</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> dischargeDateTime: </label>
                                                <input type="time" placeholder="dischargeDateTime" name="dischargeDateTime" className="form-control" value={this.state.dischargeDateTime} onChange={this.changedischargeDateTimeHandler}/>

                                            <label> Disposition: </label>
                                                <select value={this.state.disposition} onChange={this.changeDispositionHandler}>
                      <option name="Disposition" className="form-control" >
                          Home
                      </option>
                      <option name="Disposition" className="form-control" >
                          HomeWithHomeCare
                      </option>
                      <option name="Disposition" className="form-control" >
                          SkilledNursingFacility
                      </option>
                      <option name="Disposition" className="form-control" >
                          AcuteCareFacility
                      </option>
                      <option name="Disposition" className="form-control" >
                          Expired
                      </option>
                      <option name="Disposition" className="form-control" >
                          AgainstMedicalAdvice
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDischarge}>Save</button>
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

export default UpdateDischargeComponent
