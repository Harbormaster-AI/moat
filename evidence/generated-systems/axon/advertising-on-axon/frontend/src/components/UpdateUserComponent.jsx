import React, { Component } from 'react'
import UserService from '../services/UserService';

class UpdateUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                email: '',
                role: ''
        }
        this.updateUser = this.updateUser.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changeRoleHandler = this.changeRoleHandler.bind(this);
    }

    componentDidMount(){
        UserService.getUserById(this.state.id).then( (res) =>{
            let user = res.data;
            this.setState({
                firstName: user.firstName,
                lastName: user.lastName,
                email: user.email,
                role: user.role
            });
        });
    }

    updateUser = (e) => {
        e.preventDefault();
        let user = {
            userId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            email: this.state.email,
            role: this.state.role
        };
        console.log('user => ' + JSON.stringify(user));
        console.log('id => ' + JSON.stringify(this.state.id));
        UserService.updateUser(user).then( res => {
            this.props.history.push('/users');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changeRoleHandler= (event) => {
        this.setState({role: event.target.value});
    }

    cancel(){
        this.props.history.push('/users');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update User</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> Role: </label>
                                                <select value={this.state.role} onChange={this.changeRoleHandler}>
                      <option name="Role" className="form-control" >
                          Admin
                      </option>
                      <option name="Role" className="form-control" >
                          Trader
                      </option>
                      <option name="Role" className="form-control" >
                          Analyst
                      </option>
                      <option name="Role" className="form-control" >
                          Viewer
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateUser}>Save</button>
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

export default UpdateUserComponent
