import React, { Component } from 'react'
import LeadService from '../services/LeadService';

class UpdateLeadComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                company: '',
                email: '',
                phone: '',
                converted: '',
                status: '',
                source: '',
                rating: ''
        }
        this.updateLead = this.updateLead.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changecompanyHandler = this.changecompanyHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changeconvertedHandler = this.changeconvertedHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeSourceHandler = this.changeSourceHandler.bind(this);
        this.changeRatingHandler = this.changeRatingHandler.bind(this);
    }

    componentDidMount(){
        LeadService.getLeadById(this.state.id).then( (res) =>{
            let lead = res.data;
            this.setState({
                firstName: lead.firstName,
                lastName: lead.lastName,
                company: lead.company,
                email: lead.email,
                phone: lead.phone,
                converted: lead.converted,
                status: lead.status,
                source: lead.source,
                rating: lead.rating
            });
        });
    }

    updateLead = (e) => {
        e.preventDefault();
        let lead = {
            leadId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            company: this.state.company,
            email: this.state.email,
            phone: this.state.phone,
            converted: this.state.converted,
            status: this.state.status,
            source: this.state.source,
            rating: this.state.rating
        };
        console.log('lead => ' + JSON.stringify(lead));
        console.log('id => ' + JSON.stringify(this.state.id));
        LeadService.updateLead(lead).then( res => {
            this.props.history.push('/leads');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changecompanyHandler= (event) => {
        this.setState({company: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changephoneHandler= (event) => {
        this.setState({phone: event.target.value});
    }
    changeconvertedHandler= (event) => {
        this.setState({converted: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeSourceHandler= (event) => {
        this.setState({source: event.target.value});
    }
    changeRatingHandler= (event) => {
        this.setState({rating: event.target.value});
    }

    cancel(){
        this.props.history.push('/leads');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Lead</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> company: </label>
                                                <input placeholder="company" name="company" className="form-control" value={this.state.company} onChange={this.changecompanyHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone: </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> converted: </label>
                                                <input type="checkbox" placeholder="converted" name="converted" className="form-control" value={this.state.converted} onChange={this.changeconvertedHandler}/>


                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          New
                      </option>
                      <option name="Status" className="form-control" >
                          Working
                      </option>
                      <option name="Status" className="form-control" >
                          Nurturing
                      </option>
                      <option name="Status" className="form-control" >
                          Qualified
                      </option>
                      <option name="Status" className="form-control" >
                          Unqualified
                      </option>
                      <option name="Status" className="form-control" >
                          Converted
                      </option>
                    </select>

                                            <label> Source: </label>
                                                <select value={this.state.source} onChange={this.changeSourceHandler}>
                      <option name="Source" className="form-control" >
                          Web
                      </option>
                      <option name="Source" className="form-control" >
                          Referral
                      </option>
                      <option name="Source" className="form-control" >
                          Event
                      </option>
                      <option name="Source" className="form-control" >
                          Partner
                      </option>
                      <option name="Source" className="form-control" >
                          Advertisement
                      </option>
                      <option name="Source" className="form-control" >
                          Outbound
                      </option>
                      <option name="Source" className="form-control" >
                          Inbound
                      </option>
                      <option name="Source" className="form-control" >
                          Social
                      </option>
                      <option name="Source" className="form-control" >
                          Other
                      </option>
                    </select>

                                            <label> Rating: </label>
                                                <select value={this.state.rating} onChange={this.changeRatingHandler}>
                      <option name="Rating" className="form-control" >
                          Hot
                      </option>
                      <option name="Rating" className="form-control" >
                          Warm
                      </option>
                      <option name="Rating" className="form-control" >
                          Cold
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLead}>Save</button>
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

export default UpdateLeadComponent
