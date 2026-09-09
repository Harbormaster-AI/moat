import React, { Component } from 'react'
import PayrollCalendarService from '../services/PayrollCalendarService';

class UpdatePayrollCalendarComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                country: '',
                payFrequency: ''
        }
        this.updatePayrollCalendar = this.updatePayrollCalendar.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changePayFrequencyHandler = this.changePayFrequencyHandler.bind(this);
    }

    componentDidMount(){
        PayrollCalendarService.getPayrollCalendarById(this.state.id).then( (res) =>{
            let payrollCalendar = res.data;
            this.setState({
                name: payrollCalendar.name,
                country: payrollCalendar.country,
                payFrequency: payrollCalendar.payFrequency
            });
        });
    }

    updatePayrollCalendar = (e) => {
        e.preventDefault();
        let payrollCalendar = {
            payrollCalendarId: this.state.id,
            name: this.state.name,
            country: this.state.country,
            payFrequency: this.state.payFrequency
        };
        console.log('payrollCalendar => ' + JSON.stringify(payrollCalendar));
        console.log('id => ' + JSON.stringify(this.state.id));
        PayrollCalendarService.updatePayrollCalendar(payrollCalendar).then( res => {
            this.props.history.push('/payrollCalendars');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecountryHandler= (event) => {
        this.setState({country: event.target.value});
    }
    changePayFrequencyHandler= (event) => {
        this.setState({payFrequency: event.target.value});
    }

    cancel(){
        this.props.history.push('/payrollCalendars');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PayrollCalendar</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> country: </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> PayFrequency: </label>
                                                <select value={this.state.payFrequency} onChange={this.changePayFrequencyHandler}>
                      <option name="PayFrequency" className="form-control" >
                          Weekly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Biweekly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Semimonthly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Monthly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Quarterly
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePayrollCalendar}>Save</button>
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

export default UpdatePayrollCalendarComponent
