import React, { Component } from 'react'
import UnderwriterService from '../services/UnderwriterService';

class UpdateUnderwriterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                employeeId: '',
                authorityLimit: ''
        }
        this.updateUnderwriter = this.updateUnderwriter.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeemployeeIdHandler = this.changeemployeeIdHandler.bind(this);
        this.changeauthorityLimitHandler = this.changeauthorityLimitHandler.bind(this);
    }

    componentDidMount(){
        UnderwriterService.getUnderwriterById(this.state.id).then( (res) =>{
            let underwriter = res.data;
            this.setState({
                firstName: underwriter.firstName,
                lastName: underwriter.lastName,
                employeeId: underwriter.employeeId,
                authorityLimit: underwriter.authorityLimit
            });
        });
    }

    updateUnderwriter = (e) => {
        e.preventDefault();
        let underwriter = {
            underwriterId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            employeeId: this.state.employeeId,
            authorityLimit: this.state.authorityLimit
        };
        console.log('underwriter => ' + JSON.stringify(underwriter));
        console.log('id => ' + JSON.stringify(this.state.id));
        UnderwriterService.updateUnderwriter(underwriter).then( res => {
            this.props.history.push('/underwriters');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changeemployeeIdHandler= (event) => {
        this.setState({employeeId: event.target.value});
    }
    changeauthorityLimitHandler= (event) => {
        this.setState({authorityLimit: event.target.value});
    }

    cancel(){
        this.props.history.push('/underwriters');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Underwriter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> employeeId: </label>
                                                <input placeholder="employeeId" name="employeeId" className="form-control" value={this.state.employeeId} onChange={this.changeemployeeIdHandler}/>

                                            <label> authorityLimit: </label>
                                                <input placeholder="authorityLimit" name="authorityLimit" className="form-control" value={this.state.authorityLimit} onChange={this.changeauthorityLimitHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateUnderwriter}>Save</button>
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

export default UpdateUnderwriterComponent
